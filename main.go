package main

import (
	"log"
	"os"
	"fmt"
	"github.com/yanzay/tbot"
)

var (
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

	login("initial")
	dayOfWeek = "Monday"
	hourToPlay = "18:00"

	bot.HandleFunc("/find {time}", FindHandler)
	bot.HandleFunc("/play", ConfigDateHandler)
	bot.HandleFunc("/play {date}", ConfigDateHandler)
	bot.HandleFunc("/at", ConfigTimeHandler)
	bot.HandleFunc("/at {time}", ConfigTimeHandler)
	bot.HandleFunc("/when", ConfigShowHandler)
	bot.HandleFunc("/rsvp", ReservationHandler)
	bot.HandleFunc("/money", BalanceHandler)

	go autoReserve(bot);

	// Start listening for messages
	err = bot.ListenAndServe()
	log.Fatal(err)
}
