package webrcon_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/TexHik620953/go-webrcon"
)

func Test(t *testing.T) {
	conn, err := webrcon.Connect(context.Background(), "ip-addr:port", "rcon-password")
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := conn.ListPlayers()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(resp)

	conn.OnFeedback(func(f *webrcon.Feedback) {
		log.Printf("feedback: %v", f)
	})
	conn.OnReport(func(f *webrcon.Report) {
		log.Printf("report: %v", f)
	})
	conn.OnMessage(func(f *webrcon.Message) {
		log.Printf("msgs: %v", f)
	})

	<-time.After(time.Hour)
}
