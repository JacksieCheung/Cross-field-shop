package crawl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type spocGradeResp struct {
	CurrentPage   int  `json:"currentPage"`
	CurrentResult int  `json:"currentResult"`
	EntityOrField bool `json:"entityOrField"`
	Items         []struct {
		Bfzcj              string `json:"bfzcj"`
		Bh                 string `json:"bh"`
		BhID               string `json:"bh_id"`
		Bj                 string `json:"bj"`
		Cj                 string `json:"cj"`
		Cjsfzf             string `json:"cjsfzf"`
		Date               string `json:"date"`
		DateDigit          string `json:"dateDigit"`
		DateDigitSeparator string `json:"dateDigitSeparator"`
		Day                string `json:"day"`
		Jd                 string `json:"jd"`
		JgID               string `json:"jg_id"`
		Jgmc               string `json:"jgmc"`
		Jgpxzd             string `json:"jgpxzd"`
		Jsxm               string `json:"jsxm"`
		JxbID              string `json:"jxb_id"`
		Jxbmc              string `json:"jxbmc,omitempty"`
		Kcbj               string `json:"kcbj"`
		Kcgsmc             string `json:"kcgsmc,omitempty"`
		Kch                string `json:"kch"`
		KchID              string `json:"kch_id"`
		Kclbmc             string `json:"kclbmc"`
		Kcmc               string `json:"kcmc"`
		Kcxzdm             string `json:"kcxzdm"`
		Kcxzmc             string `json:"kcxzmc"`
		Key                string `json:"key"`
		Kkbmmc             string `json:"kkbmmc"`
		Ksxz               string `json:"ksxz"`
		Ksxzdm             string `json:"ksxzdm"`
		Listnav            string `json:"listnav"`
		LocaleKey          string `json:"localeKey"`
		Month              string `json:"month"`
		NjdmID             string `json:"njdm_id"`
		Njmc               string `json:"njmc"`
		Pageable           bool   `json:"pageable"`
		QueryModel         struct {
			CurrentPage   int           `json:"currentPage"`
			CurrentResult int           `json:"currentResult"`
			EntityOrField bool          `json:"entityOrField"`
			Limit         int           `json:"limit"`
			Offset        int           `json:"offset"`
			PageNo        int           `json:"pageNo"`
			PageSize      int           `json:"pageSize"`
			ShowCount     int           `json:"showCount"`
			Sorts         []interface{} `json:"sorts"`
			TotalCount    int           `json:"totalCount"`
			TotalPage     int           `json:"totalPage"`
			TotalResult   int           `json:"totalResult"`
		} `json:"queryModel"`
		Rangeable   bool   `json:"rangeable"`
		RowID       string `json:"row_id"`
		Rwzxs       string `json:"rwzxs,omitempty"`
		Sfdkbcx     string `json:"sfdkbcx,omitempty"`
		Sfxwkc      string `json:"sfxwkc"`
		Sfzh        string `json:"sfzh"`
		Tjrxm       string `json:"tjrxm,omitempty"`
		Tjsj        string `json:"tjsj,omitempty"`
		TotalResult string `json:"totalResult"`
		UserModel   struct {
			Monitor    bool   `json:"monitor"`
			RoleCount  int    `json:"roleCount"`
			RoleKeys   string `json:"roleKeys"`
			RoleValues string `json:"roleValues"`
			Status     int    `json:"status"`
			Usable     bool   `json:"usable"`
		} `json:"userModel"`
		Xb     string `json:"xb"`
		Xbm    string `json:"xbm"`
		Xf     string `json:"xf"`
		Xfjd   string `json:"xfjd"`
		Xh     string `json:"xh"`
		XhID   string `json:"xh_id"`
		Xm     string `json:"xm"`
		Xnm    string `json:"xnm"`
		Xnmmc  string `json:"xnmmc"`
		Xqm    string `json:"xqm"`
		Xqmmc  string `json:"xqmmc"`
		Xsbjmc string `json:"xsbjmc"`
		Xslb   string `json:"xslb"`
		Year   string `json:"year"`
		Zsxymc string `json:"zsxymc"`
		ZyhID  string `json:"zyh_id"`
		Zymc   string `json:"zymc"`
		Khfsmc string `json:"khfsmc,omitempty"`
	} `json:"items"`
	Limit       int           `json:"limit"`
	Offset      int           `json:"offset"`
	PageNo      int           `json:"pageNo"`
	PageSize    int           `json:"pageSize"`
	ShowCount   int           `json:"showCount"`
	SortName    string        `json:"sortName"`
	SortOrder   string        `json:"sortOrder"`
	Sorts       []interface{} `json:"sorts"`
	TotalCount  int           `json:"totalCount"`
	TotalPage   int           `json:"totalPage"`
	TotalResult int           `json:"totalResult"`
}

// 期末成绩爬取
func (m *spocManager) getGrade() (GradesInfo, error) {
	url := fmt.Sprintf(
		"http://xk.ccnu.edu.cn/jwglxt/cjcx/cjcx_cxXsgrcj.html?su=%s&doType=query&gnmkdm=N305005",
		m.uid,
	)
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf("xqm=&xnm=&_search=false&nd=%d&queryModel.showCount=100&queryModel.currentPage=1&queryModel.sortName=&queryModel.sortOrder=&time=1",
			time.Now().UnixMilli(),
		))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := m.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gradeResp spocGradeResp
	err = json.Unmarshal(body, &gradeResp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ret := make([]GradeInfo, len(gradeResp.Items))
	for idx, item := range gradeResp.Items {
		ret[idx].CourseName = item.Kcmc
		ret[idx].SemesterYear, _ = strconv.Atoi(item.Xnm)
		ret[idx].SemesterSpan, _ = strconv.Atoi(item.Xqm)
		ret[idx].GPA, _ = strconv.ParseFloat(item.Jd, 64)
		ret[idx].Grade, _ = strconv.ParseFloat(item.Cj, 64)
		ret[idx].ID = item.Kch
	}

	return ret, nil
}
