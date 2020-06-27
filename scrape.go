package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Post struct {
	Title          string  `json:"title"`
	NumofUpvotes   int 	   `json:"upvotes"`
	User   	       string  `json:"user"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has title attribute call callback
	c.OnHTML("div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE", func(e *colly.HTMLElement) {
		fmt.Printf("Title: %s\n", e.Text)
	})

	// On every a element which has upvotes attribute call callback
	c.OnHTML("div._23h0-EcaBUorIHC-JZyh6J > div > div", func(e *colly.HTMLElement) {
		fmt.Printf("# of Upvotes: %s\n", e.Text)
	})

	// // On every a element which has author attribute call callback
	// c.OnHTML("a._2tbHP6ZydRpjI44J3syuqC._23wugcdiaj44hdfugIAlnX.oQctV4n0yUb0uiHDdGnmE", func(e *colly.HTMLElement) {
	// 	fmt.Printf("user: %s\n", e.Text)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.reddit.com/")
	
}
