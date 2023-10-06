package util

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
	"time"

	"Data-acquisition-subsystem/pkg/errno"
	"golang.org/x/net/publicsuffix"
)

var TIMEOUT = time.Duration(30 * time.Second)

type AccountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

type CCNUUserCenter struct {
	Errcode string
	Errmsg  string
	User    struct {
		DeptId       string
		DeptName     string
		Email        string
		Id           string
		Mobile       string
		Name         string
		SchoolEmail  string
		Status       int
		UserFace     string
		Username     string
		Usernumber   string
		Usertype     string
		UsertypeName string
		Xb           string
	}
}

func MakeAccountPreflightRequest() (*AccountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &AccountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
		return params, err
	}

	// 发起请求
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return params, err
	}

	// 读取 Body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return params, errors.New("Can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		log.Println("Can not get form paramater: lt")
		return params, errors.New("Can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get form paramater: execution")
		return params, errors.New("Can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get form paramater: _eventId")
		return params, errors.New("Can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}

// account.ccnu.edu.cn 模拟登录，用于验证账号密码是否可以正常登录
func MakeAccountRequest(sid, password string, params *AccountReqeustParams, client *http.Client) error {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)

	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	if err != nil {
		log.Print(err)
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	// fmt.Println(string(body))

	// check
	reg := regexp.MustCompile("class=\"success\"")
	matched := reg.MatchString(string(body))
	if !matched {
		log.Println("Wrong sid or pwd")
		return errno.ErrAuthFailed
	}

	log.Println("Login successfully")
	return nil
}

// xk.ccnu.edu.cn 模拟登录
func MakeXKLogin(client *http.Client) error {
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fxk.ccnu.edu.cn%2Fsso%2Fpziotlogin", nil)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// LoginRequest ... 检查 sid 并且 返回 username
func LoginRequest(Sid, Password string) (error, string) {
	params, err := MakeAccountPreflightRequest()
	if err != nil {
		return errno.InternalServerError, ""
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return errno.InternalServerError, ""
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}

	if err := MakeAccountRequest(Sid, Password, params, &client); err != nil {
		return errno.ErrAuthFailed, ""
	}

	// 必要的步骤，
	if err := MakeONERequest(&client); err != nil {
		return errno.ErrAuthFailed, ""
	}

	// 获取用户名并返回
	err, username := GetUserNameFromCCNU(&client)
	if err != nil {
		return errno.ErrAuthFailed, ""
	}

	// 登出
	if err := MakeAccountLogoutRequest(&client); err != nil {
		return errno.ErrAuthFailed, ""
	}

	return nil, username
}

// GetUserNameFromCCNU ... 获取用户信息
func GetUserNameFromCCNU(client *http.Client) (error, string) {
	request, err := http.NewRequest("POST", "http://one.ccnu.edu.cn/user_portal/userDetailCcnu", nil)
	if err != nil {
		log.Print(err)
		return err, ""
	}

	// token 单独拿出来一个字段
	var token string

	u, err := url.Parse("http://one.ccnu.edu.cn")
	if err != nil {
		log.Println(err)
		return err, ""
	}

	for _, cookie := range client.Jar.Cookies(u) {
		if cookie.Name == "PORTAL_TOKEN" {
			token = cookie.Value
			fmt.Println("token:" + token)
		}
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Origin", "http://one.ccnu.edu.cn")
	request.Header.Set("Host", "http://one.ccnu.edu.cnn")
	request.Header.Set("Referer", "http://one.ccnu.edu.cn/index")
	request.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
		return err, ""
	}

	var user CCNUUserCenter
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return err, ""
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return err, ""
	}

	return nil, user.User.Name
}

// MakeONERequest ... one.ccnu.edu.cn
func MakeONERequest(client *http.Client) error {
	request, err := http.NewRequest("GET", "http://one.ccnu.edu.cn", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}

	u, err := url.Parse("http://one.ccnu.edu.cn")
	if err != nil {
		log.Println(err)
		return err
	}

	for _, cookie := range client.Jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
	return nil
}

// MakeAccountLogoutRequest ... 退出信息门户登录
func MakeAccountLogoutRequest(client *http.Client) error {
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/logout", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
