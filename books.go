package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// create collector using Colly package
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse((func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	}))

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://books.toscrape.com/")
}
