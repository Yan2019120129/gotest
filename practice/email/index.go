package email

import (
	"fmt"
	"mime"
	"net/smtp"
	"strings"
)

type Client struct {
	Host     string
	Port     int
	Username string // 发件邮箱
	Password string // SMTP授权码
	Nickname string // 发件人昵称
}

func NewClient(
	host string,
	port int,
	username string,
	password string,
	nickname string,
) *Client {
	return &Client{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Nickname: nickname,
	}
}

func (c *Client) Send(
	to []string,
	subject string,
	body string,
) error {

	auth := smtp.PlainAuth(
		"",
		c.Username,
		c.Password,
		c.Host,
	)

	msg := BuildMessage(
		c.Nickname,
		c.Username,
		to,
		subject,
		body,
	)

	addr := fmt.Sprintf(
		"%s:%d",
		c.Host,
		c.Port,
	)

	return smtp.SendMail(
		addr,
		auth,
		c.Username,
		to,
		msg,
	)
}

func BuildMessage(
	fromName string,
	fromEmail string,
	to []string,
	subject string,
	body string,
) []byte {

	encodedName := mime.BEncoding.Encode(
		"UTF-8",
		fromName,
	)

	encodedSubject := mime.BEncoding.Encode(
		"UTF-8",
		subject,
	)

	var msg strings.Builder

	msg.WriteString(fmt.Sprintf(
		"From: %s <%s>\r\n",
		encodedName,
		fromEmail,
	))

	msg.WriteString(fmt.Sprintf(
		"To: %s\r\n",
		strings.Join(to, ","),
	))

	msg.WriteString(fmt.Sprintf(
		"Subject: %s\r\n",
		encodedSubject,
	))

	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	msg.WriteString("Content-Transfer-Encoding: 8bit\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(body)

	return []byte(msg.String())
}
