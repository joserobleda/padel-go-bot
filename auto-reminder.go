package main

import (
	"fmt"
	"time"

	"github.com/yanzay/tbot"
)

func autoReminder(server *tbot.Server) {
	fmt.Println("Auto reminder enabled")

	location, _ := time.LoadLocation("Europe/Madrid")

	for date := range time.Tick(time.Minute) {
		date = date.In(location)
		currentTime := date.Format("15:04")

		if currentTime != "22:45" {
			continue
		}

		err, rsvp := getLatestActiveReservation()

		if err != nil {
			fmt.Println(err)
			break
		}

		remainingTime := time.Until(rsvp.date)
		remainingHours := remainingTime.Hours()

		if remainingHours > 24 {
			continue
		}

		fmt.Println("Remaining time: " + remainingTime.String())

		if chatId != 0 {
			msg := "Eeeey!!\n"
			msg += "There is a match tomorrow at " + rsvp.date.Format("15:04") + "! "
			msg += "Don't forget to get your bag!!"

			server.Send(chatId, msg)
		}
	}
}
