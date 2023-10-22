package util

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

var Email *email.Email

func init() {
	Email.From = "邮箱"
	//设置主题
	Email.Subject = "User's Email Confirm"
}

func SendEmail(email, code string) error {
	Email.To = []string{email}
	//设置文件发送的内容
	Email.HTML = []byte(`
    <h1>your code is ` + code + `</h1>    
    `)
	//设置服务器相关的配置 TODO: 更改邮箱和授权码
	err := Email.Send("smtp.qq.com:25", smtp.PlainAuth("",
		"邮箱", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		return err
	}

	return nil
}
