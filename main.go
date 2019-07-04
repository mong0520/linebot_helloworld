package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var SSLCertPath = "/path/to/ssl"
var SSLPrivateKeyPath = "/path/to/key"
var bot *linebot.Client

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	secret := os.Getenv("LINE_SECRET")
	token := os.Getenv("LINE_TOKEN")
	port := os.Getenv("PORT")
	runMode := os.Getenv("RUN_MODE")

	// init linebot
	bot, err = linebot.New(secret, token)
	if err != nil {
		log.Println(err)
	}

	// init web hook
	http.HandleFunc("/callback", callbackHandler)
	addr := fmt.Sprintf(":%s", port)
	if strings.ToLower(runMode) == "https" {
		log.Printf("Secure listen on %s with \n", addr)
		err := http.ListenAndServeTLS(addr, SSLCertPath, SSLPrivateKeyPath, nil)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Listen on %s\n", addr)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Panic(err)
		}
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		// Handle message event
		if event.Type == linebot.EventTypeMessage {
			log.Printf("Receieve Event Type = %s from User [%s], or Room [%s] or Group [%s]\n",
				event.Type, event.Source.UserID, event.Source.RoomID, event.Source.GroupID)

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Println("Text = ", message.Text)
				replyTextMessage(event, "I got a text message: ["+message.Text+"] from User ID: "+event.Source.UserID)
			}
		} else if event.Type == linebot.EventTypePostback {
			log.Println("got a postback event")
		} else {
			log.Printf("got a %s event\n", event.Type)
		}
	}
}

func replyTextMessage(event *linebot.Event, text string) {
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
		log.Println(err)
	}
}
