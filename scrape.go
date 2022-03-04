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

func storeNames(names []NewReleases) {
	json, _ := json.MarshalIndent(names, "", " ")

	_ = ioutil.WriteFile("getNames.json", json, 0644)
}

func main() {
	namesData := make([]NewReleases, 0)
	c := colly.NewCollector()

	c.OnHTML("div.tab_item_name", func(e *colly.HTMLElement) {
		new_release := NewReleases{Name: e.Text}
		namesData = append(namesData, new_release)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://store.steampowered.com/explore/new/")

	storeNames(namesData)
}
