package main

import (
	"github.com/yanzay/tbot"
	"time"
	"fmt"
)

func autoReserve(server *tbot.Server) {
	fmt.Println("Auto reserve enabled");

	location, _ := time.LoadLocation("Europe/Madrid")

	for date := range time.Tick(time.Second) {
		date = date.In(location);

		currentTime := date.Format("15:04:05")
		newAvailableDate := date.AddDate(0, 0, 6);
		newAvailableDay := newAvailableDate.Weekday().String();

		if (newAvailableDay != dayOfWeek) {
			// fmt.Println("Tracks are not ready yet for next " + dayOfWeek)
			continue
		}

		if (currentTime != "00:00:02") {
			// fmt.Println(currentTime + " is not the right time to reserve")
			continue
		}

		bro = login();
		reserveDate := newAvailableDate.Format("02-01-2006") + " " + hourToPlay
		startMessage := "Begin reservation for next " + dayOfWeek + " at " + hourToPlay + " (" + reserveDate + ")";

		fmt.Println(startMessage);
		if chatId != 0 {
			server.Send(chatId, startMessage);
		}

		resultMessage := reserve(bro, reserveDate)
		fmt.Println(resultMessage);
		if chatId != 0 {
			server.Send(chatId, resultMessage);
		}
	}
}