package main

import (
  "fmt"
  "os"
  "github.com/headzoo/surf"
  "github.com/headzoo/surf/browser"
  "github.com/headzoo/surf/agent"
)

func login() *browser.Browser {
	bow := surf.NewBrowser()
  bow.SetUserAgent(agent.Chrome())

  err := bow.Open("https://canaldeisabel.padelclick.com/customer/login")
  if err != nil {
  	panic(err)
  }

  username := os.Getenv("USERNAME")
  pass := os.Getenv("PASSWORD")

  // Log in to the site.
  fm, _ := bow.Form("form#customerLogin")
  fm.Input("email", username)
  fm.Input("password", pass)
  if fm.Submit() != nil {
     panic(err)
  }

  // pd := url.Values{}
  // 	pd.Set("date", "18/9/2017");
  //
  // err = bow.PostForm("https://canaldeisabel.padelclick.com/customerzone/timetable?type=56", pd)
  // if err != nil {
  //   panic(err)
  // }

  // script := strings.TrimSpace(bow.Find("script").Last().Text());
  // jsonString := strings.TrimPrefix(strings.SplitN(script, ";", 3)[0], "var timetable = ")
  //
  // var re = regexp.MustCompile(`resources:{(.+),resources`)
  // jsonString = re.FindString(jsonString)

  // var re = regexp.MustCompile(`([a-zA-Z]+):`)
  // jsonString = re.ReplaceAllString(jsonString, `"$1":`)
  //
  // re = regexp.MustCompile(`([0-9]+):{`)
  // jsonString = re.ReplaceAllString(jsonString, `"$1":{`)
  // errMessage := strings.TrimSpace(bow.Find(".error .message").Text())


  // var f interface{}
  // err = json.Unmarshal([]byte(jsonString), &f)
  // if err != nil {
  //   panic(err)
  // }
  //
  fmt.Println("login done")

  // checkDate(bow, "12/10/2017 12:00")
  //
  // fmt.Println(bow.Body())

  return bow;
}
