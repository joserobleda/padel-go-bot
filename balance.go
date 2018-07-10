package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/yanzay/tbot"
)

func BalanceHandler(message *tbot.Message) {
	// store latest message as a global (OMG!)
	chatId = message.ChatID

	err, balance := getBalance()
	if err != nil {
		message.Reply("Ooops, I can't get your balance... something weird happens")

		return
	}

	msg := "You have " + strconv.FormatFloat(balance, 'f', 2, 64) + "€ in your padelclick wallet"

	message.Reply(msg)
}

func getBalance() (error, float64) {
	bow := login("balance")
	err := bow.Open("https://" + domain + "/customerzone/vouchers")
	if err != nil {
		panic(err)
	}

	nodes := bow.Dom().Find(".defaultRow .numericCell")
	if nodes.Length() == 0 {
		return errors.New("Balance not found"), 0
	}

	balanceStr := nodes.First().Text()
	balanceStr = strings.Replace(balanceStr, ",", ".", 1)
	balance, _ := strconv.ParseFloat(strings.TrimSpace(balanceStr), 64)

	return nil, balance
}
