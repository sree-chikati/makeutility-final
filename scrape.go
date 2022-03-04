package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Necklace struct {
	Name          string
	OriginalPrice float64
	Image         string
}

// Make template from data
func makeHTML(names []Necklace) {
	// SSF STUFF
	println("Hello")
}

func main() {
	products := []Necklace{}
	var necklaceNames []string
	var necklacePrices []float64
	var necklaceImages []string

	c := colly.NewCollector()

	// Get necklace names
	c.OnHTML("h3.v2-listing-card__title", func(e *colly.HTMLElement) {
		necklaceNames = append(necklaceNames, strings.TrimSpace(e.Text))
	})

	// Get necklace prices
	c.OnHTML("span.currency-value", func(e *colly.HTMLElement) {
		pricefloat, _ := strconv.ParseFloat(e.Text, 32)
		necklacePrices = append(necklacePrices, pricefloat)
	})

	// Get necklace image urls
	c.OnHTML("div.v2-listing-card", func(e *colly.HTMLElement) {
		imgSrc := e.ChildAttr("img", "src")
		necklaceImages = append(necklaceImages, imgSrc)
	})

	// informs you it's visiting the URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// visits given URL
	c.Visit("https://www.etsy.com/c/jewelry/necklaces/pendants?ref=catnav-10855")

	for i := 0; i < len(necklaceNames); i++ {
		necklaceData := Necklace{Name: necklaceNames[i], OriginalPrice: necklacePrices[i], Image: necklaceImages[i]}
		products = append(products, necklaceData)
	}

	// Pass data to HTML generator
	makeHTML(products)
}
