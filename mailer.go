package webwatch

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

// A Mailer sends emails.
type Mailer interface {

	// SendRestarted sends a 'restarted' mail.
	SendRestarted(now time.Time)

	// SendReport sends a 'OK' or 'ERR' mail.
	SendReport(now time.Time, ok bool, text string)
}

// NewMailer creates Mailer with a configuration.
func NewMailer(config MailConfig) Mailer {
	return mailerImpl{config}
}

type mailerImpl struct {
	config MailConfig
}

func (mi mailerImpl) SendRestarted(now time.Time) {
	tstamp := now.UTC().Format(time.RFC3339)
	body := fmt.Sprintf("%s\r\n", tstamp)
	mi.send("restarted", body)
}

func (mi mailerImpl) SendReport(now time.Time, ok bool, text string) {
	suffix := ""
	if ok {
		suffix = "OK"
	} else {
		suffix = "ERR"
	}
	tstamp := now.UTC().Format(time.RFC3339)
	body := fmt.Sprintf("%s\r\n\r\n%s\r\n", tstamp, text)
	mi.send(suffix, body)
}

func (mi mailerImpl) send(suffix, body string) {
	subject := mi.config.Subject + " " + suffix
	subject = strings.TrimSpace(subject)
	INFO.Printf("send mail %q to %q\n", subject, mi.config.To)
	auth := smtp.PlainAuth("", mi.config.Username, mi.config.Password, mi.config.Host)
	addr := mi.config.Host + ":25"
	msg := "To: " + mi.config.To + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"\r\n" +
		body + "\r\n" +
		"\r\n"
	err := smtp.SendMail(addr, auth, mi.config.From, []string{mi.config.To}, []byte(msg))
	if err != nil {
		INFO.Printf("cannot send mail: %s\n", err)
		return
	}
	INFO.Printf("mail sent\n")
}
