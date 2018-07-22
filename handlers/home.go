package handlers

import (
	"fmt"
	"github.com/cwise/tfsa/libhttp"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type WebData struct {
	Message string
	Unused  int
	Current int
	Monthly int
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}

func PostHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	r.ParseForm()

	minimum := 5500 / 12
	unused, _ := strconv.Atoi(r.Form.Get("unused"))
	monthly, _ := strconv.Atoi(r.Form.Get("monthly"))
	current, _ := strconv.Atoi(r.Form.Get("current"))
	message := ""

	if monthly < minimum {
		message = "Your monthly payment is less than the minimum amount. You'll never catch up at that rate."
	} else {
		year := time.Now().Year()
		month := time.Now().Month()
		date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		accumulated := minimum * int(month)
		target := unused + accumulated - current

		for target > 0 {
			target = target + minimum - monthly
			date = date.AddDate(0, 1, 0)
		}

		message = fmt.Sprintf("You will catch up on contributions by %s %d", date.Month(), date.Year())
	}

	wd := WebData{
		Message: message,
		Unused:  unused,
		Monthly: monthly,
		Current: current,
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, &wd)
}
