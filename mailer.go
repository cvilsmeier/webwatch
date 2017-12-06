package main

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

type Mailer interface {
	SendRestarted(now time.Time)
	SendReport(now time.Time, ok bool, text string)
}

func NewMailer(config MailConfig) Mailer {
	return mailerImpl{config}
}

type mailerImpl struct {
	config MailConfig
}

func (this mailerImpl) SendRestarted(now time.Time) {
	tstamp := now.UTC().Format(time.RFC3339)
	body := fmt.Sprintf("%s\r\n", tstamp)
	this.send("restarted", body)
}

func (this mailerImpl) SendReport(now time.Time, ok bool, text string) {
	suffix := ""
	if ok {
		suffix = "OK"
	} else {
		suffix = "ERR"
	}
	tstamp := now.UTC().Format(time.RFC3339)
	body := fmt.Sprintf("%s\r\n\r\n%s\r\n", tstamp, text)
	this.send(suffix, body)
}

func (this mailerImpl) send(suffix, body string) {
	subject := this.config.Subject + " " + suffix
	subject = strings.TrimSpace(subject)
	INFO.Printf("send mail %q to %q\n", subject, this.config.To)
	auth := smtp.PlainAuth("", this.config.Username, this.config.Password, this.config.Host)
	addr := this.config.Host + ":25"
	msg := "To: " + this.config.To + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"\r\n" +
		body + "\r\n" +
		"\r\n"
	err := smtp.SendMail(addr, auth, this.config.From, []string{this.config.To}, []byte(msg))
	if err != nil {
		INFO.Printf("cannot send mail: %s\n", err)
		return
	}
	INFO.Printf("mail sent\n")
}
