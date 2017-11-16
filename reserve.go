package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/headzoo/surf/browser"
	"github.com/yanzay/tbot"
)

func ReserveHandler(message *tbot.Message) {
	date := message.Vars["date"]

	if "no" == date {
		message.Reply("Ok, perfect")

		return
	}

	message.Reply("Trying to reserve...")

	resultMessage, resultCode := reserve(bow, date)

	fmt.Println("result:" + resultMessage)
	message.Reply(resultMessage)

	if resultCode == 0 {
		message.Reply("Enjoy!! and remember to add the " + date + " to your calendar!")
	}
}

func reserve(bro *browser.Browser, date string) (string, int) {
	name, resource, reservation := checkDate(bro, date)

	if resource == "" {
		return "Oh! Looks like there are no available tracks! Sorry!!", 1
	}

	fmt.Println(bro.Url(), bro.StatusCode())

	pdr := url.Values{}
	pdr.Set("date", date)
	pdr.Set("duration", "90")
	pdr.Set("idReservation", reservation)
	pdr.Set("idResource", resource)
	pdr.Set("paymentPending", "")
	err := bro.PostForm("https://canaldeisabel.padelclick.com/customerZone/newReservationPost", pdr)

	if err != nil {
		panic(err)
	}

	paymentfm, err := bro.Form("form#paymentForm")
	if err != nil {
		fmt.Println(bro.Url(), bro.Body())
		return err.Error(), 2
	}

	paymentfm.Input("idPaymentMethod", "968")
	if paymentfm.Submit() != nil {
		panic("form#paymentForm error")
	}

	if strings.TrimSpace(bro.Title()) == "Reserva Confirmada" {
		return "Track " + name + " reserved!", 0
	}

	return "I can't reserve! Maybe you need to put more money? Sorry!!", 3
}
