package main

import (
	"github.com/mong0520/linebot-ptt-beauty/models"
	"log"
	"os"
	"http"
)

func main() {
	var err error
	meta = m
	secret := os.Getenv("ChannelSecret")
	token := os.Getenv("ChannelAccessToken")
	bot, err = linebot.New(secret, token)
	if err != nil {
		log.Println(err)
	}
	//log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	//port := "8080"
	addr := fmt.Sprintf(":%s", port)
	runMode := os.Getenv("RUNMODE")
	log.Printf("Run Mode = %s\n", runMode)
	if strings.ToLower(runMode) == ModeHttps {
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
		if event.Type == linebot.EventTypeMessage {
			log.Printf("Receieve Event Type = %s from User [%s](%s), or Room [%s] or Group [%s]\n",
				event.Type, userDisplayName, event.Source.UserID, event.Source.RoomID, event.Source.GroupID)

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Println("Text = ", message.Text)
				// textHander(event, message.Text)			
			}
		} else if event.Type == linebot.EventTypePostback {
			log.Println("got a postback event")
			log.Println(event.Postback.Data)
			// postbackHandler(event)

		} else {
			log.Printf("got a %s event\n", event.Type)
		}
	}
}