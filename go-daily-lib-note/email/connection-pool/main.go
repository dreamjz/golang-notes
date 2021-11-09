package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	ch := make(chan *email.Email, 10)
	pool, err := email.NewPool("smtp.gmail.com:587", 4, smtp.PlainAuth("", "xxx@gmail.com", "xxx", "smtp.gmail.com"))
	if err != nil {
		log.Fatal("create pool failed:", err)
	}
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := pool.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email: %v sent error:%v\n", e, err)
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		e := email.NewEmail()
		e.To = []string{"xxx@qq.com"}
		e.Cc = []string{"xxx@qq.com"}
		e.Bcc = []string{"xxx@gmail.com"}
		e.Subject = "Awesome Email:" + strconv.Itoa(i+1)
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte(`
<h1>Awesome Email</h1>
<p>Text Body is, of course, supported!</p>
`)
		ch <- e
	}
	close(ch)
	wg.Wait()
}
