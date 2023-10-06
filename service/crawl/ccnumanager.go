package crawl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

var _ ManagerBuilder = &spocManagerBuilder{}
var _ Manager = &spocManager{}

type spocManagerBuilder struct {
	uid      string
	password string
}

func (m *spocManagerBuilder) SetUid(uid string) ManagerBuilder {
	m.uid = uid
	return m
}

func (m *spocManagerBuilder) SetPassword(password string) ManagerBuilder {
	m.password = password
	return m
}

func (m *spocManagerBuilder) Build() (Manager, error) {
	result := spocManager{
		uid:        m.uid,
		password:   m.password,
		httpClient: http.Client{},
	}

	return &result, nil
}

type spocManager struct {
	uid        string
	password   string
	spocUserid string
	httpClient http.Client
}

type SpocUserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID       string      `json:"id"`
		Archived bool        `json:"archived"`
		Username string      `json:"username"`
		Phone    interface{} `json:"phone"`
		Email    string      `json:"email"`
		UserInfo struct {
			ID                     string      `json:"id"`
			Realname               string      `json:"realname"`
			Sex                    interface{} `json:"sex"`
			Age                    interface{} `json:"age"`
			Phone                  interface{} `json:"phone"`
			Email                  string      `json:"email"`
			Qq                     interface{} `json:"qq"`
			Wechat                 interface{} `json:"wechat"`
			DegreeCode             interface{} `json:"degreeCode"`
			MajorCode              interface{} `json:"majorCode"`
			DepartmentCode         string      `json:"departmentCode"`
			University             interface{} `json:"university"`
			LoginName              string      `json:"loginName"`
			HeadImageURL           interface{} `json:"headImageUrl"`
			AcademicTitle          interface{} `json:"academicTitle"`
			Sign                   interface{} `json:"sign"`
			Status                 string      `json:"status"`
			Addtime                int64       `json:"addtime"`
			Updatetime             int64       `json:"updatetime"`
			UserID                 interface{} `json:"userId"`
			Password               interface{} `json:"password"`
			DomainName             interface{} `json:"domainName"`
			DepartmentName         interface{} `json:"departmentName"`
			RoleName               interface{} `json:"roleName"`
			UserRoleStr            interface{} `json:"userRoleStr"`
			GroupRoleCode          interface{} `json:"groupRoleCode"`
			GroupManaRoleName      interface{} `json:"groupManaRoleName"`
			BaseRoleDepartmentList interface{} `json:"baseRoleDepartmentList"`
		} `json:"userInfo"`
		Privileges []interface{} `json:"privileges"`
	} `json:"data"`
}

func (m *spocManager) Reload() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("Got error while creating cookie jar %s\n", err.Error())
		return errors.New("Got error while creating cookie jar %s" + err.Error())
	}
	m.httpClient = http.Client{
		Jar: jar,
	}
	return m.preload()
}

func (m *spocManager) wrapperRequest(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
}

func (m *spocManager) preload() error {
	spocHomeURL := "http://spoc.ccnu.edu.cn/"
	method := "GET"

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Println(fmt.Sprintf("Got error while creating cookie jar %s\n", err.Error()))
		return err
	}
	m.httpClient = http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(method, spocHomeURL, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	m.wrapperRequest(req)

	res, err := m.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	casURL := "https://account.ccnu.edu.cn/cas/login?service=http://spoc.ccnu.edu.cn/userLoginController/userCasLogin"
	req, err = http.NewRequest(method, casURL, nil)

	if err != nil {
		panic(err)
	}
	m.wrapperRequest(req)

	res, err = m.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	var jsessionid string
	urlObj, _ := url.Parse("https://account.ccnu.edu.cn/cas/login")
	for _, v := range m.httpClient.Jar.Cookies(urlObj) {
		if v.Name == "JSESSIONID" {
			jsessionid = v.Name
		}
	}
	if jsessionid == "" {
		return errors.New("not found jsessionid")
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// regex
	exp := regexp.MustCompile(`<input type="hidden" name="lt" value="(.*?)" />`)
	result := exp.FindAllStringSubmatch(string(content), -1)
	if len(result) == 0 {
		res.Body.Close()
		return errors.New("regexp lt result length eq 0")
	}
	lt := result[0][1]
	res.Body.Close()

	casLoginURL := "https://account.ccnu.edu.cn/cas/login;jsessionid=" + jsessionid + "?service=http://spoc.ccnu.edu.cn/userLoginController/userCasLogin"
	method = "POST"

	payload := strings.NewReader("username=" + m.uid + "&password=" + m.password + "&lt=" + lt + "&execution=e1s1&_eventId=submit&submit=%E7%99%BB%E5%BD%95")

	req, err = http.NewRequest(method, casLoginURL, payload)

	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Origin", "https://account.ccnu.edu.cn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Referer", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,en-US;q=0.6")

	res, err = m.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	res.Body.Close()

	userInfoURl := "http://spoc.ccnu.edu.cn/studentHomepage/getUserInfo"
	method = "POST"

	req, err = http.NewRequest(method, userInfoURl, nil)

	if err != nil {
		panic(err)
	}
	m.wrapperRequest(req)

	res, err = m.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	content, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	if !strings.Contains(string(content), "获取用户信息成功") {
		fmt.Println(string(content))
		return errors.New("login failed")
	}
	var userInfo SpocUserInfo
	err = json.Unmarshal(content, &userInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal spoc userinfo error: %v", err))
	}
	m.spocUserid = userInfo.Data.ID

	// 预登录 xk
	xkSSOURL := "http://xk.ccnu.edu.cn/sso/pziotlogin"
	method = "GET"

	req, err = http.NewRequest(method, xkSSOURL, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	m.wrapperRequest(req)

	res, err = m.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	content, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	if !strings.Contains(string(content), "退出") {
		fmt.Println(string(content))
		return errors.New("login failed ONE XK CCNU")
	}

	return nil
}

// 爬取云课堂数据
func (m *spocManager) CrawlCourseInfo() (CourseInfo, error) {
	return m.getSpocInfo()
}

// 爬取成绩
func (m *spocManager) CrawlGradeInfo() (GradesInfo, error) {
	return m.getGrade()
}

// 登出
func (m *spocManager) Logout() error {
	logoutURL := "http://one.ccnu.edu.cn/security_portal/logout"
	method := "GET"

	req, err := http.NewRequest(method, logoutURL, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := m.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	return nil
}
