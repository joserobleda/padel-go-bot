package main;

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/headzoo/surf/browser"
)

func checkDate(bro *browser.Browser, date string) (string, string) {
	resources := map[string]string{
	  "#2": "1478",
	  "#3": "1479",
	  "#4": "1480",
	  "#5": "1481",
	  "#1": "1477", // dont like this one
	}
  
	for name, id := range resources {
	  pdr := url.Values{}
	  pdr.Set("idResource", id);
	  pdr.Set("localDatetime", date);
	  // fmt.Println("Buscando pistas para " + date);
  
	  err := bro.PostForm("https://canaldeisabel.padelclick.com/customerZone/newReservation", pdr)
	  if err != nil {
			panic(err)
	  }
	
		title := strings.TrimSpace(bro.Title())

	  if title == "Nueva Reserva" {
			fmt.Println("Pista " + name + " libre para " + date);
  
			return name, id;
	  } else {
			fmt.Println("Pista " + name + " ocupada " + date + " ("+ title +")");
	  }
	}
  
	return "", ""
  }
  