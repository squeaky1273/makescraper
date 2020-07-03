package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
	"github.com/gocolly/colly"
)

type Post struct {
	Title          string  `json:"title"`
	NumofUpvotes   string  `json:"upvotes"`
	// User   	       string  `json:"user"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	posts := scrapeReddit()
	serializePosts(posts)
}

func scrapeReddit() {
	c := colly.NewCollector()

	var posts Posts

	c.OnHTML(".rpBJOHq2PR60pnwJlUyP0", func(wrapper *colly.HTMLElement) {
		// On every a element which has title attribute call callback
		wrapper.ForEach("div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE", func(i int, title *colly.HTMLElement) {
			if i < 10 {
				posts.Posts = append(posts.Posts, Post{Title: title.Text})
			}
		})

		// On every a element which has upvotes attribute call callback
		wrapper.ForEach("div._23h0-EcaBUorIHC-JZyh6J > div > div", func(k int, upvotes *colly.HTMLElement) {
			if k < 10 {
				posts.Posts[k].NumofUpvotes = upvotes.Text
			}

		})
		// // On every a element which has author/user attribute call callback
		// c.ForEach("#t3_hi3oht > div._1poyrkZ7g36PawDueRza-J._11R7M_VOgKO1RJyRSRErT3 > div.BiNC74axuTz66dlnEgicY > div.cZPZhMe-UCZ8htPodMyJ5 > div._3AStxql1mQsrZuUIFP9xSg.nU4Je7n-eSXStTBAPMYt8 > div:nth-child(4) > a", func(e *colly.HTMLElement) {
		// 	fmt.Printf("User: %s\n", e.Text)
		// })
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.reddit.com/")
}

func serializePosts(posts Post) {
	j, err := json.Marshal(posts)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("output.json", j, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}