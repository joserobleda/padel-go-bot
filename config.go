package main

import (
  "strings"
  "github.com/yanzay/tbot"
)

func ConfigHandler(message *tbot.Message) {
  // store latest message as a global (OMG!)
  chatId = message.ChatID;

  if message.Vars["time"] == "" {
    buttons := [][]string{
      {"/play Monday", "/play Tuesday", },
      {"/play Wednesday", "/play Thursday"},
    }

    message.ReplyKeyboard("What day?", buttons)

    return;
  }

  s := strings.Split(message.Vars["time"], " ")

  if len(s) > 0 {
    dayOfWeek = strings.Title(strings.ToLower(s[0]));
  }

  if len(s) > 1 {
    hourToPlay = s[1];
  }

  message.Reply("Cool! I'll will reserve for " + dayOfWeek + "s at " + hourToPlay)
}
