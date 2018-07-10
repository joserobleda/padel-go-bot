package main

import (
	"fmt"
	"log"
	"os"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
)

var (
	bow *browser.Browser
)

func login(reason string) *browser.Browser {
	if nil == bow {
		bow = surf.NewBrowser()
	}

	bow.SetUserAgent(agent.Chrome())

	err := bow.Open("https://" + domain + "/customer/login")
	if err != nil {
		panic(err)
	}

	div := bow.Dom().Find(".c-login-user .c-login-user__text")

	if div.Length() == 1 {
		fmt.Println("Logged in as " + div.Text() + " [" + reason + "]")

		return bow
	}

	username := os.Getenv("LOGIN")
	pass := os.Getenv("PASSWORD")

	fmt.Println("Login with username: " + username + " [" + reason + "]")

	// Log in to the site.
	fm, _ := bow.Form("form#customerLogin")
	fm.Input("email", username)
	fm.Input("password", pass)
	fm.Input("keepSession", "true")
	if fm.Submit() != nil {
		panic(err)
	}

	span := bow.Dom().Find("span.metadataSubtitle")

	if span.Length() != 1 {
		log.Fatal("Cannot login")
	}

	return bow
}
