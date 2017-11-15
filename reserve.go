package main

import (
	"fmt"
	"strings"

	"github.com/headzoo/surf/browser"
)

func reserve(bro *browser.Browser, date string) (string, int) {
	name, resource := checkDate(bro, date)

	if resource == "" {
		return "Oh! Looks like there are no available tracks! Sorry!!", 1
	}

	rsrvefm, _ := bro.Form("form#newreservation")
	rsrvefm.Input("duration", "90")
	if rsrvefm.Submit() != nil {
		panic("form#newreservation error")
	}

	paymentfm, err := bro.Form("form#paymentForm")
	if err != nil {
		fmt.Println(bro.Body())
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
