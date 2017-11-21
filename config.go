package main

import (
	"strings"

	"github.com/yanzay/tbot"
)

func ConfigDateHandler(message *tbot.Message) {
	// store latest message as a global (OMG!)
	chatId = message.ChatID

	if message.Vars["date"] == "" {
		buttons := [][]string{
			{"/play Monday", "/play Tuesday"},
			{"/play Wednesday", "/play Thursday"},
		}

		message.ReplyKeyboard("What day?", buttons)

		return
	}

	dayOfWeek = message.Vars["date"]

	buttons := [][]string{
		{"/at 17:00", "/at 17:30"},
		{"/at 18:00", "/at 18:30"},
	}

	msg := "Cool! I'll reserve for " + dayOfWeek + "s. *Do you want to play at " + hourToPlay + "?*"
	message.ReplyKeyboard(msg, buttons, tbot.OneTimeKeyboard, tbot.WithMarkdown)
}

func ConfigTimeHandler(message *tbot.Message) {
	// store latest message as a global (OMG!)
	chatId = message.ChatID

	if message.Vars["time"] == "" {
		buttons := [][]string{
			{"/at 17:00", "/at 17:30"},
			{"/at 18:00", "/at 18:30"},
		}

		message.ReplyKeyboard("What hour?", buttons, tbot.OneTimeKeyboard)

		return
	}

	hourToPlay = message.Vars["time"]

	message.Reply("Cool! Consider it done!")

	ConfigShowHandler(message)
}

func ConfigShowHandler(message *tbot.Message) {
	// store latest message as a global (OMG!)
	chatId = message.ChatID

	ndate := nextDateForSchedule(dayOfWeek, hourToPlay)
	tdate := ndate.AddDate(0, 0, -7)

	msgs := []string{
		"I'll will reserve for *" + dayOfWeek + "s at " + hourToPlay + "*. Next available date for schedule is *" + ndate.Format("02-01-2006") + "*.",
		"_And don't worry, I'll automatically perform the reservation on " + tdate.Format("02-01-2006") + "'s night_",
	}

	message.Reply(strings.Join(msgs, "\n"), tbot.WithMarkdown)
}
