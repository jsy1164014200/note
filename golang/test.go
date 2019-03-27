package main

import (
	"fmt"
	"log"
	"os"
	"github.com/emersion/go-imap"
	// "github.com/emersion/go-message"
	"github.com/emersion/go-imap/client"

    "github.com/go-gomail/gomail"
)

var authCode = os.Getenv("QQAUTH")

func main() {
	
 
	log.Println("Connecting to server...")
 
	// 连接服务器
	c, err := client.DialTLS("imap.qq.com:993", nil)
	if err != nil {
        fmt.Println(err)
    }
	log.Println("连接服务器")
 
	// 结束后退出登录
	defer c.Logout()
 
	// 登录
	//args[1]是用户名，args[2]是imap密码
	if err := c.Login("1164014200@qq.com",authCode); err != nil {
		log.Fatal(err)
	}
	log.Println("登陆邮箱")
 
	// 获取邮箱列表
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
 
	log.Println("邮箱列表:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}
 
	if err := <-done; err != nil {
		log.Fatal(err)
	}
 
	// 选择收件箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}


	from := uint32(1)
    to := mbox.Messages  // uint型 
    // log.Println(to)
	seqset := new(imap.SeqSet)
    seqset.AddRange(from, to) // 声明一个集合
    
    fmt.Println(seqset)
	// attrs := []string{"BODY[]", imap.FlagsMsgAttr}
    messages := make(chan *imap.Message, 10)
    
	done = make(chan error, 1)
	go func() {
        done <- c.Fetch(seqset, []imap.FetchItem{"BODY[]"}, messages)
    }()
    for message := range messages {
		fmt.Println(message.Body)
		// fmt.Println(message.)
    }
	// sendEmail("1164014200@qq.com","1141741348@qq.com","dkfjsdkl","skdfjksdjf",authCode)
}



func sendEmail(from string,to string,title string,contentHTML string,password string) {
    m := gomail.NewMessage() 
    m.SetAddressHeader("From",from,"from")
    m.SetHeader("To",m.FormatAddress(to,"to"))
    m.SetHeader("Subject",title)
    m.SetBody("text/html",contentHTML)
    d := gomail.NewPlainDialer("smtp.qq.com",465,from,password)
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}
