package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func scrapeRutracker() {
	c := colly.NewCollector()

	// Find and extract the href value in ".magnet-link"
	c.OnHTML(".magnet-link", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		fmt.Println("Magnet Link:", href)
	})

	// Visit the website
	err := c.Visit("https://rutracker.org/forum/viewtopic.php?t=6447215")
	if err != nil {
		log.Fatal(err)
	}
}