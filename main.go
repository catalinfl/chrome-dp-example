package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
)

type LocationData struct {
	A string `json:"a"`
	B string `json:"b,omitempty"`
}

var data []LocationData

var htmlLink1 string = "https://earthquake.usgs.gov/earthquakes/map/?extent=3.25021,-154.77539&extent=62.18601,-35.24414"
var htmlLink2 string = "https://infp.ro/"

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var location []string

	err := chromedp.Run(ctx, fetchHTML(htmlLink1, &location))
	if err != nil {
		panic(err)
	}

	for _, loc := range location {
		if strings.Contains(loc, "L-ai simţit? Lasă-ne un feedback!") {
			parts := strings.SplitN(loc, "L-ai simţit? Lasă-ne un feedback!", 2)
			a, b := parts[0], parts[1]

			data = append(data, LocationData{A: a, B: b})
		} else {
			data = append(data, LocationData{A: loc})
		}
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	os.WriteFile("earthquake.json", jsonData, 0644)

	if err != nil {
		panic(err)
	}
}

// func fetchHTML(url string, location *[]string, time *[]string) chromedp.Tasks {
// 	return chromedp.Tasks{
// 		chromedp.Navigate(url),
// 		chromedp.WaitVisible(`body`, chromedp.ByQuery),
// 		chromedp.Evaluate(`Array.from(document.querySelectorAll('h6.header')).map(n => n.innerText)`, location),
// 		chromedp.Evaluate(`Array.from(document.querySelectorAll('span.time')).map(n => n.innerText)`, time),
// 	}
// }

func fetchHTML(url string, location *[]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(htmlLink2),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('tr[title="Detalii cutremur"]')).map(n => n.innerText)`, location),
	}
}
