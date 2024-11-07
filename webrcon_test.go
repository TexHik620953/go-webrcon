package webrcon_test

import (
	"context"
	"fmt"
	"go-webrcon"
	"log"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	conn, err := webrcon.Connect(context.Background(), os.Getenv("RCON_ADDR"), os.Getenv("RCON_PWD"))
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
