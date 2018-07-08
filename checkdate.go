package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/headzoo/surf/browser"
)

type track struct {
	id   string
	name string
}

func checkDate(bro *browser.Browser, date string) (string, string, string) {
	var resources [5]track
	resources[2] = track{"1480", "#4"}
	resources[4] = track{"1477", "#1"}
	resources[1] = track{"1479", "#3"}
	resources[0] = track{"1478", "#2"}
	resources[3] = track{"1481", "#5"}

	for _, track := range resources {
		pdr := url.Values{}
		pdr.Set("idResource", track.id)
		pdr.Set("localDatetime", date)
		fmt.Println("Buscando en pista " + track.name + " para " + date)

		err := bro.PostForm(domain+"/customerZone/newReservation", pdr)
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
