package helper

import (
	_ "gopkg.in/mail.v2"
)

var Client *MailClient

type MailClient struct {
	handel *mail.Dialer
	name   string
}

type CCList struct {
	Email string
	Name  string
}
type Info struct {
	To      []string
	CC      CCList
	Subject string
	Content string
	Attach  string
}

func (c *MailClient) Run(name, passwd, host string, port int) {
	c.handel = mail.NewDialer(host, port, name, passwd)
	c.handel.StartTLSPolicy = mail.MandatoryStartTLS
	c.name = name
}

func (c *MailClient) Send(info Info) error {
	m := mail.NewMessage()
	m.SetHeader("From", c.name)
	m.SetHeader("To", info.To...)

	m.SetAddressHeader("Cc", info.CC.Email, info.CC.Name)

	m.SetHeader("Subject", info.Subject)
	m.SetBody("text/html", info.Content)

	if err := c.handel.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
