package crawl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// 云课堂相关

type spocCoursesResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		TermName string `json:"termName"`
		SiteList []struct {
			SiteID             string      `json:"siteId"`
			SiteName           string      `json:"siteName"`
			TeacherCode        string      `json:"teacherCode"`
			DepartmentCode     string      `json:"departmentCode"`
			CourseCode         string      `json:"courseCode"`
			IsLocalDomain      string      `json:"isLocalDomain"`
			TermCode           string      `json:"termCode"`
			Otherdomainid      string      `json:"otherdomainid"`
			CourseType         interface{} `json:"courseType"`
			ClassDesc          interface{} `json:"classDesc"`
			ClassIconURL       string      `json:"classIconUrl"`
			Ispublic           string      `json:"ispublic"`
			Isdelete           string      `json:"isdelete"`
			CertificateType    string      `json:"certificateType"`
			RegularGradeWeight interface{} `json:"regularGradeWeight"`
			CertificateStatus  interface{} `json:"certificateStatus"`
			DepName            interface{} `json:"depName"`
			TeacherName        string      `json:"teacherName"`
			UserID             string      `json:"userId"`
			DepartmentName     string      `json:"departmentName"`
			CourseName         string      `json:"courseName"`
			TermName           string      `json:"termName"`
			DomainName         string      `json:"domainName"`
			CsID               interface{} `json:"csId"`
			SiteAccessCount    int         `json:"siteAccessCount"`
		} `json:"siteList"`
		TermCode string `json:"termCode"`
	} `json:"data"`
}

func (m *spocManager) getSpocInfo() (CourseInfo, error) {
	var ret CourseInfo

	url := "http://spoc.ccnu.edu.cn/studentHomepage/getMySite"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"termCode": null,"userId": "%s"}`, m.spocUserid))

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := m.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	var courseResp spocCoursesResp
	err = json.Unmarshal(body, &courseResp)
	if err != nil {
		fmt.Println(string(body))
		return ret, errors.New(fmt.Sprintf("获取云课堂信息失败：%v", err))
	}
	//fmt.Println(string(body))
	if courseResp.Msg != "获取站点信息成功" {
		return ret, errors.New(fmt.Sprintf("获取云课堂信息失败：%s", courseResp.Msg))
	}

	for _, courses := range courseResp.Data {
		for _, course := range courses.SiteList {
			info, err := m.getCourseInfo(course.SiteID, m.spocUserid)
			if err != nil {
				// TODO
				log.Println(err)
			}
			sy, _ := strconv.Atoi(course.TermCode[:4])
			ss, _ := strconv.Atoi(course.TermCode[5:])
			now := Course{
				TotalHomeWorksNum: len(info),
				SemesterYear:      sy,
				SemesterSpan:      ss,
				Name:              course.CourseName,
				ID:                course.CourseCode,
				HomeWorks:         info,
			}
			ret.Courses = append(ret.Courses, now)
			fmt.Println(now)
		}
	}
	ret.TotalCoursesNum = len(ret.Courses)

	return ret, nil
}

type spocHomeWorkResp struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
		Size     int `json:"size"`
		StartRow int `json:"startRow"`
		EndRow   int `json:"endRow"`
		Total    int `json:"total"`
		Pages    int `json:"pages"`
		List     []struct {
			ID            string      `json:"id"`
			SiteID        string      `json:"siteId"`
			SiteName      interface{} `json:"siteName"`
			Title         string      `json:"title"`
			Content       string      `json:"content"`
			Creater       string      `json:"creater"`
			Creatername   interface{} `json:"creatername"`
			Begintime     int64       `json:"begintime"`
			Endtime       int64       `json:"endtime"`
			Isgroup       string      `json:"isgroup"`
			AttmentNum    interface{} `json:"attmentNum"`
			GroupNum      int         `json:"groupNum"`
			PointNum      int         `json:"pointNum"`
			StudentNum    int         `json:"studentNum"`
			CommitNum     int         `json:"commitNum"`
			UnCommitNum   interface{} `json:"unCommitNum"`
			GroupPoint    interface{} `json:"groupPoint"`
			PersonalPoint interface{} `json:"personalPoint"`
			Createtime    int64       `json:"createtime"`
			Status        string      `json:"status"`
			RoleCode      interface{} `json:"roleCode"`
			Isevaluation  string      `json:"isevaluation"`
		} `json:"list"`
		FirstPage        int   `json:"firstPage"`
		PrePage          int   `json:"prePage"`
		NextPage         int   `json:"nextPage"`
		LastPage         int   `json:"lastPage"`
		IsFirstPage      bool  `json:"isFirstPage"`
		IsLastPage       bool  `json:"isLastPage"`
		HasPreviousPage  bool  `json:"hasPreviousPage"`
		HasNextPage      bool  `json:"hasNextPage"`
		NavigatePages    int   `json:"navigatePages"`
		NavigatepageNums []int `json:"navigatepageNums"`
	} `json:"data"`
}

func (m *spocManager) getCourseInfo(siteId, userId string) ([]HomeWork, error) {
	url := "http://spoc.ccnu.edu.cn/assignment/getStudentAssignmentList"
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf(`{"siteId":"%s","userId":"%s","pageNum":1,"pageSize":100}`,
			siteId, userId,
		))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var homeWorkResp spocHomeWorkResp
	err = json.Unmarshal(body, &homeWorkResp)
	if err != nil {
		return nil, err
	}
	ret := make([]HomeWork, homeWorkResp.Data.Total)
	for k, v := range homeWorkResp.Data.List {
		ret[k].Status, err = strconv.Atoi(v.Status)
		if err != nil {
			return ret, err
		}
		ret[k].Name = v.Title
		if ret[k].Status == Finished {
			t := reflect.ValueOf(v.PersonalPoint).Kind()
			if t == reflect.Int {
				ret[k].Score = float64(v.PersonalPoint.(int))
			} else if t == reflect.Float64 {
				ret[k].Score = v.PersonalPoint.(float64)
			}
		}
		ret[k].Time = v.Begintime
	}
	return ret, nil
}
