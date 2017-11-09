package main

import (
	"regexp"
	"strconv"
  "strings"
  "time"
  "github.com/PuerkitoBio/goquery"
  "github.com/yanzay/tbot"
)

func ReservationHandler(message *tbot.Message) {
  // store latest message as a global (OMG!)
  chatId = message.ChatID;

  rsvp := getLatestReservation();

  if (rsvp.status == "Anulada") {
    message.Reply("Ooops, you don't have any active reservations");
    
    return;
  }

  msg := "Your latest reservation is for "+ rsvp.date +" in "+ rsvp.track + "\n"
  msg += "Duration: " + rsvp.duration.String() +" - "
  msg += "Price: " + strconv.FormatFloat(rsvp.price, 'f', 2, 64) + "€" +"\n"

  message.Reply(msg)
}

type Reservation struct {
  date string
  price float64
  track string
  duration time.Duration
  status string
}

func getLatestReservation() Reservation {
  reservations := getReservations();

  return reservations[0];
}

func getReservations() []Reservation {
  slice := make([]Reservation, 0)

  bow := login("reservations")
  err := bow.Open("https://canaldeisabel.padelclick.com/bookings/customerzone/reservations")
  if err != nil {
  	panic(err)
  }

  bow.Dom().Find(".czrow").Each(func(_ int, s *goquery.Selection) {
    dateNode := s.Find(".dateHeader")
    infoNode := dateNode.Next().Next()
    priceNode := infoNode.NextFiltered("div")

    dateStr := strings.TrimSpace(dateNode.Text());
    infoStr := strings.TrimSpace(infoNode.Text())
    priceStr := strings.TrimSpace(priceNode.Text());
    priceStr = priceStr[6:len(priceStr)-3]
    priceStr = strings.Replace(priceStr, ",", ".", 1);

    // info parts
    parts := regexp.MustCompile("[\\.,]").Split(infoStr, 5);
    hourStr := parts[0][6:len(parts[0])]
    durationStr := strings.TrimSpace(parts[1])
    durationStr = durationStr[0:len(durationStr)-8]
    duration, _ := time.ParseDuration(durationStr+"m")
    trackStr := strings.TrimSpace(parts[2])
    trackStr = trackStr[3:len(trackStr)]
    statusStr := strings.TrimSpace(parts[3])
    statusStr = statusStr[7:len(statusStr)]

    dateTimeStr := dateStr + " a las " + hourStr;

    price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

    slice = append(slice, Reservation{dateTimeStr, price, trackStr, duration, statusStr})
  })

  return slice
}
