package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.To = []string{"xxx@qq.com"}
	e.Cc = []string{"xxx@qq.com"}
	e.Bcc = []string{"xxx@gmail.com"}
	e.Subject = "Awesome Email"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(`
<h1>Awesome Email</h1>
<p>Text Body is, of course, supported!</p>
`)
	file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal("create file error ")
	}
	file.WriteString("This is a mail attachment")

	e.AttachFile("text.txt")

	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "xxx@gmail.com", "xxx", "smtp.gmail.com"))
	if err != nil {
		log.Fatal("send email failed ", err)
	}
}
