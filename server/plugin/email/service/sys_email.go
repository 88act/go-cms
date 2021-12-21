package service

import (
	"go-cms/plugin/email/utils"
)

type EmailService struct {
}

//@author: [88act-3](https://github.com/88act)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

//@author: [88act-3](https://github.com/88act)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
