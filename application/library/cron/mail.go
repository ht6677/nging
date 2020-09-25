package cron

import (
	"strings"

	"github.com/admpub/nging/application/dbschema"
	"github.com/admpub/nging/application/library/cron/send"
	"github.com/admpub/nging/application/library/cron/writer"
	alertRegistry "github.com/admpub/nging/application/registry/alert"
	"github.com/webx-top/echo/defaults"
	"github.com/webx-top/echo/param"
)

var (
	// SendMail 发送Email
	// @param toEmail 收信邮箱
	// @param toUsername 收信人名称
	// @param title 邮件标题
	// @param content 邮件内容
	// @param ccList 抄送地址
	SendMail = send.Mail

	// SendMailWithID 发送Email(带ID参数)
	SendMailWithID = send.MailWithID

	// SendMailWithNoticer 发送Email(带Noticer参数)
	SendMailWithNoticer = send.MailWithNoticer

	// SendMailWithIDAndNoticer 发送Email(带ID和Noticer参数)
	SendMailWithIDAndNoticer = send.MailWithIDAndNoticer

	NewCmdRec = writer.New
)

type OutputWriter = writer.OutputWriter

func OtherSender(params param.Store) error {
	if alertRegistry.SendTopic == nil {
		return nil
	}
	ctx := defaults.NewMockContext()
	return alertRegistry.SendTopic(ctx, `cron`, params)
}

func EmailSender(params param.Store) error {
	task, ok := params.Get(`task`).(dbschema.NgingTask)
	if !ok {
		return nil
	}
	var ccList []string
	if len(task.NotifyEmail) > 0 {
		ccList = strings.Split(task.NotifyEmail, "\n")
		for index, email := range ccList {
			email = strings.TrimSpace(email)
			if len(email) == 0 {
				continue
			}
			ccList[index] = email
		}
	}
	if len(ccList) == 0 {
		return nil
	}
	toEmail := ccList[0]
	toUsername := strings.SplitN(toEmail, "@", 2)[0]
	if len(ccList) > 1 {
		ccList = append([]string{}, ccList[1:]...)
	} else {
		ccList = []string{}
	}
	ct, ok := params.Get(`content`).(send.ContentType)
	if !ok {
		return nil
	}
	content := ct.EmailContent(params)
	return SendMail(toEmail, toUsername, params.String(`title`), content, ccList...)
}

func init() {
	AddSender(EmailSender)
	AddSender(OtherSender)
	alertRegistry.Topics.Add(`cron`, `定时任务`)
}
