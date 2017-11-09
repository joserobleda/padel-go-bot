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

		if (currentTime != "00:00:01") {
			// fmt.Println(currentTime + " is not the right time to reserve")
			continue
		}

		if chatId != 0 {
			msg := "Hi there, it's time!\n"
			msg += "I'm going to reserve for next "+ dayOfWeek +" ("+ newAvailableDate.Format("2006-01-02") +")"

			server.Send(chatId, msg);
		}

		bro := login("auto-reserve");
		date, _ = time.Parse("2006-01-02 15:04", newAvailableDate.Format("2006-01-02") +" "+ hourToPlay)

		for _, timeRange := range TimesRanges {
			rdate := date.Add(timeRange)

			reserveTimeStr := rdate.Format("15:04")
			startMessage := "I'm trying to reserve for "+ reserveTimeStr +"...";

			fmt.Println(startMessage);
			if chatId != 0 {
				server.Send(chatId, startMessage);
      }

			reserveDateStr := rdate.Format("02-01-2006 15:04")
			resultMessage, resultCode := reserve(bro, reserveDateStr)

			fmt.Println(resultMessage);
			if chatId != 0 {
				server.Send(chatId, resultMessage);

				if resultCode == 0 {
					server.Send(chatId, "Enjoy!! and remember to add the "+ reserveDateStr +" to your calendar!");
				}
			}

			// 1 = no available tracks, so try again with next time range
			if resultCode != 1 {
				break;
			}
		}
	}
}