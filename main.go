package main

import (
	"crypto/tls"
	"flag"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

const (
	//SenderAddr sender addr
	SenderAddr string = "serialt@qq.com"

	//SederPassword sender password
	SederPassword string = "xxxxxxxxx"

	//SMTPServer smtp server
	SMTPServer string = "smtp.qq.com"
	//SMTPPort smtp port
	SMTPPort int = 465
)

//Mailer 邮件配置
type Mailer struct {
	SenderAddr     string
	SenderPassword string
	SMTPServer     string
	SMTPPort       int
}

// NewMailer new一个对象
func NewMailer(sA string, sPwd string, sServer string, sPort int) Mailer {
	return Mailer{
		SenderAddr:     sA,
		SenderPassword: sPwd,
		SMTPServer:     sServer,
		SMTPPort:       sPort,
	}
}

//Recver 邮件信息处理
func Recver() ([]string, string, string, string, string) {
	var recvUser string
	var subject string
	var body string
	var filename string
	var mailtype string
	flag.StringVar(&recvUser, "c", "", "收邮件地址,格式为 a11@qq.com,a22@gmail.com")
	flag.StringVar(&mailtype, "t", "", "邮件发送方式,g 群发邮件,s 一对一发邮件,默认是g")
	flag.StringVar(&subject, "s", "", "邮件主题")
	flag.StringVar(&body, "m", "", "邮件内容")
	flag.StringVar(&filename, "f", "", "添加的附件")

	flag.Parse()
	recvs := strings.Split(recvUser, ",")
	return recvs, subject, body, filename, mailtype
}

//SendMail sendmail
func SendMailByGroup(mailTo []string, subject string, body string, filename string) error {
	my := NewMailer(SenderAddr, SederPassword, SMTPServer, SMTPPort)
	m := gomail.NewMessage()
	m.SetHeader("From", my.SenderAddr)
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if filename != "" {
		m.Attach(filename)
	}
	d := gomail.NewDialer(my.SMTPServer, my.SMTPPort, my.SenderAddr, my.SenderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}

func SendMailByOne(mailTo []string, subject string, body string, filename string) error {
	var err error
	for _, v := range mailTo {
		my := NewMailer(SenderAddr, SederPassword, SMTPServer, SMTPPort)
		m := gomail.NewMessage()
		m.SetHeader("From", my.SenderAddr)
		m.SetHeader("To", v)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)
		if filename != "" {
			m.Attach(filename)
		}
		d := gomail.NewDialer(my.SMTPServer, my.SMTPPort, my.SenderAddr, my.SenderPassword)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		err = d.DialAndSend(m)
	}
	return err
}

func main() {
	mailTo, subject, body, filename, mailtype := Recver()
	switch mailtype {
	case "g", "":
		err := SendMailByGroup(mailTo, subject, body, filename)
		if err != nil {
			log.Println("Send mail failed", err)
			return
		}
		log.Println("Send mail successfully!")
	case "s":
		err := SendMailByOne(mailTo, subject, body, filename)
		if err != nil {
			log.Println("Send mail failed", err)
			return
		}
		log.Println("Send mail successfully!")

	default:

		log.Println("the send mailtye is unknown", mailtype)
	}

}
