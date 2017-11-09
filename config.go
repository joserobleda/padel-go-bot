package main

import (
	"strings"
  "github.com/yanzay/tbot"
)

func ConfigDateHandler(message *tbot.Message) {
  // store latest message as a global (OMG!)
  chatId = message.ChatID;

  if message.Vars["date"] == "" {
    buttons := [][]string{
      {"/play Monday", "/play Tuesday", },
      {"/play Wednesday", "/play Thursday"},
    }

    message.ReplyKeyboard("What day?", buttons)

    return;
  }

  dayOfWeek = message.Vars["date"]

  buttons := [][]string{
    {"/at 17:00", "/at 17:30" },
    {"/at 18:00", "/at 18:30"},
  }

  msg := "Cool! I'll will reserve for " + dayOfWeek + "s. Do you want to play at "+ hourToPlay +"?";
  message.ReplyKeyboard(msg, buttons)
}

func ConfigTimeHandler(message *tbot.Message) {
  // store latest message as a global (OMG!)
  chatId = message.ChatID;

  if message.Vars["time"] == "" {
    buttons := [][]string{
      {"/at 17:00", "/at 17:30" },
      {"/at 18:00", "/at 18:30"},
    }

    message.ReplyKeyboard("What hour?", buttons)

    return;
  }

  hourToPlay = message.Vars["time"];

  message.Reply("Cool! Consider it done!")

  ConfigShowHandler(message);
}

func ConfigShowHandler(message *tbot.Message) {
  // store latest message as a global (OMG!)
  chatId = message.ChatID;

  ndate := nextDate(dayOfWeek, hourToPlay);
  tdate := ndate.AddDate(0, 0, -7);

  msgs := []string {
    "I'll will reserve for "+ dayOfWeek +"s at "+ hourToPlay +". Next available date is "+ ndate.Format("02-01-2006") +".",
    "And don't worry, I'll automatically perform the reservation on "+ tdate.Format("02-01-2006") +"'s night",
  }

  message.Reply(strings.Join(msgs, "\n"))
}
