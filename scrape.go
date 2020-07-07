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
	Title          string
	NumofUpvotes   string
	// User   	       string
}

type Posts struct {
	Posts string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

func jsonRedditFile(filename string, e Posts) {
	jsonPost, err := json.Marshal(e)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("output.json", jsonPost, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := colly.NewCollector()
	var posts [] Posts

	c.OnHTML(".rpBJOHq2PR60pnwJlUyP0", func(e *colly.HTMLElement) {
		// // On every a element which has title attribute call callback
		// e.ForEach("div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE", func(title *colly.HTMLElement) {
		// 	posts = append(Title: title.Text)
		// })

		// // On every a element which has upvotes attribute call callback
		// e.ForEach("div._23h0-EcaBUorIHC-JZyh6J > div > div", func(upvotes *colly.HTMLElement) {
		// 	posts = append(NumofUpvotes: upvotes.Text)
		// })

		post := Posts{Posts: e.Text}

		j, _ := json.Marshal(post)
		_ = ioutil.WriteFile("output.json", j, os.ModePerm)

		for _, post := range posts {
			jsonRedditFile("output.json", post)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.reddit.com/")
}


// * Doesn't work

// c.OnHTML(".rpBJOHq2PR60pnwJlUyP0", func(e *colly.HTMLElement) {
// 	// On every a element which has title attribute call callback
// 	e.ForEach(".rpBJOHq2PR60pnwJlUyP0", func(i int, wrap *colly.HTMLElement) {
// 		posts.Title := wrap.ChildText("div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE")
// 		posts.NumOfUpvotes := wrap.ChildText("div._23h0-EcaBUorIHC-JZyh6J > div > div")
		
// 		post := Posts{Posts: wrap.Text}
// 		posts = append(post, posts)

// 		j, _ := json.Marshal(post)
// 		_ = ioutil.WriteFile("output.json", j, os.ModePerm)

// 		for _, post := range posts {
// 			jsonRedditFile("output.json", post)
// 		}
// 	})

// })