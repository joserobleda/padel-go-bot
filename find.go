package main

import (
  "strings"
  "github.com/yanzay/tbot"
)

func FindHandler(message *tbot.Message) {
  s := strings.Split(message.Vars["time"], " ")

  if len(s) != 2 {
    message.Reply("Specify the day and the hour");

    return;
  }

  day := strings.Title(strings.ToLower(s[0]));
  hour := s[1];

  message.Reply("Looking for track on next " + day + " at " + hour + "...");

  date := prepareDate(day, hour).Format("02-01-2006 15:04");
  name, resource := checkDate(bro, date)

  if resource == "" {
    message.Reply("No hay pistas")

    return
  }

  message.Reply("Pista " + name + " libre!")
}