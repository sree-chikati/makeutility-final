package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type NewReleases struct {
	Name string
}

// function makes a file with the NewReleased names
func storeNames(names []NewReleases) {
	json, _ := json.MarshalIndent(names, "", " ")
	_ = ioutil.WriteFile("output.json", json, 0644)
}

func main() {

	// stores an array of new release structs
	namesData := make([]NewReleases, 0)
	c := colly.NewCollector()

	c.OnHTML("div.tab_item_name", func(e *colly.HTMLElement) {
		// creates NewRelease struct with e.Text for Name attribute
		new_release := NewReleases{Name: e.Text}

		// appends to the array namesData
		namesData = append(namesData, new_release)
	})

	// informs you it's visiting the URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// visits given URL
	c.Visit("https://store.steampowered.com/explore/new/")

	// calls storeNames with array of NewReleases
	storeNames(namesData)
}
