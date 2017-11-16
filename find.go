package main

import (
	"strings"

	"github.com/yanzay/tbot"
)

func FindHandler(message *tbot.Message) {
	s := strings.Split(message.Vars["time"], " ")

	if len(s) != 2 {
		message.Reply("Specify the day and the hour")

		return
	}

	day := strings.Title(strings.ToLower(s[0]))
	hour := s[1]

	message.Reply("Looking for track on next " + day + " at " + hour + "...")

	date := nextAvailableDate(day, hour).Format("02-01-2006 15:04")
	bro := login("find")
	name, resource, rsvp := checkDate(bro, date)

	if resource == "" {
		message.Reply("No available tracks")

		return
	}

	message.Reply("Track " + name + " is free! Your rsvp id is " + rsvp + "\nYou can reserve this track for 10 min")

	buttons := [][]string{
		{"/reserve no"},
		{"/reserve " + date},
	}

	msg := "*Do you want me to reserve it now?*"
	message.ReplyKeyboard(msg, buttons, tbot.OneTimeKeyboard, tbot.WithMarkdown)
}
