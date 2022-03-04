package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type NewReleases struct {
	Name string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	namesData := make([]NewReleases, 0)
	var names []string
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("div.tab_item_name", func(e *colly.HTMLElement) {
		names := append(names, e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://store.steampowered.com/explore/new/")
	for i := 0; i < len(names); i++ {
		namesData := NewReleases{Name: names[i]}
	}
}
