package main

import (
  "strings"
  "github.com/headzoo/surf/browser"
  "fmt"
)

func reserve(bro *browser.Browser, date string) (string) {
  name, resource := checkDate(bro, date)

  if resource == "" {
    return "No hay pistas";
  }
  // fmt.Println(bro.Body());

  rsrvefm, _ := bro.Form("form#newreservation")
  rsrvefm.Input("duration", "90")
  if rsrvefm.Submit() != nil {
     panic("form#newreservation error")
  }

  paymentfm, err := bro.Form("form#paymentForm")
  if (err != nil) {
    fmt.Println(bro.Body());
    return err.Error();
  }

  paymentfm.Input("idPaymentMethod", "968")
  if paymentfm.Submit() != nil {
     panic("form#paymentForm error")
  }

  if strings.TrimSpace(bro.Title()) == "Reserva Confirmada" {
    return "Pista " + name + " reservada!";
  }
  
  return "No he podido reservar! Parece que no hay pasta o algo así!";
}
