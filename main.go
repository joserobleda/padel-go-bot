package main

import (
	"log"
	"os"
	"github.com/headzoo/surf/browser"
	"github.com/yanzay/tbot"
	"fmt"
)

var (
	bro *browser.Browser
	dayOfWeek string
	hourToPlay string
	chatId int64
)

func main() {
	fmt.Println("Starting...")

	token := os.Getenv("TELEGRAM_TOKEN")

	// Create new telegram bot server using token
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}

	bro = login();	
	dayOfWeek = "Monday"
	hourToPlay = "18:00"

	bot.HandleFunc("/find {time}", FindHandler)
	bot.HandleFunc("/play", ConfigHandler)
	bot.HandleFunc("/play {time}", ConfigHandler)

	go autoReserve(bot);

	// Start listening for messages
	err = bot.ListenAndServe()
	log.Fatal(err)
}
