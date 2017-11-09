package main

import (
  "fmt"
  "os"
  "log"
  "github.com/headzoo/surf"
  "github.com/headzoo/surf/browser"
  "github.com/headzoo/surf/agent"
)

func login(reason string) *browser.Browser {
	bow := surf.NewBrowser()
  bow.SetUserAgent(agent.Chrome())

  err := bow.Open("https://canaldeisabel.padelclick.com/customer/login")
  if err != nil {
  	panic(err)
  }

  username := os.Getenv("LOGIN")
  pass := os.Getenv("PASSWORD")

  fmt.Println("Login with username: " + username + " ["+ reason +"]");

  // Log in to the site.
  fm, _ := bow.Form("form#customerLogin")
  fm.Input("email", username)
  fm.Input("password", pass)
  if fm.Submit() != nil {
     panic(err)
  }

  span := bow.Dom().Find("span.metadataSubtitle");
  
  if span.Length() != 1 {
    log.Fatal("Cannot login")
  }

  return bow;
}
