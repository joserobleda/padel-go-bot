package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/headzoo/surf/browser"
)

func checkDate(bro *browser.Browser, date string) (string, string, string) {
	resources := getTracks()

	for _, track := range resources {
		pdr := url.Values{}
		pdr.Set("idResource", track.id)
		pdr.Set("localDatetime", date)
		fmt.Println("Buscando en pista " + track.name + " para " + date)

		err := bro.PostForm("https://"+domain+"/customerZone/newReservation", pdr)
		if err != nil {
			panic(err)
		}

		reservation, _ := bro.Dom().Find("[name=idReservation]").Attr("value")

		title := strings.TrimSpace(bro.Title())

		if title == "Nueva Reserva" {
			fmt.Println("Pista " + track.name + " libre para " + date + " (" + reservation + ")")

			return track.name, track.id, reservation
		}

		div := bro.Dom().Find(".error .message")
		errString := strings.TrimSpace(div.Text())

		fmt.Println("Pista " + track.name + " ocupada " + date + " - " + errString + " (" + title + ")")
	}

	return "", "", ""
}
