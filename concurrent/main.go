package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type task struct {
	url   string
	depth uint
}

var (
	wg    sync.WaitGroup
	count int
)

func main() {
	url := flag.String("url", "", "Target URL")
	depth := flag.Uint("depth", 2, "Depth to crawl")
	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	t := task{*url, *depth}

	wg.Add(1)
	go crawl(t)

	wg.Wait()
	fmt.Println("Total page", count)
}

func validateURI(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}

func crawl(t task) {
	fmt.Printf("URL: %s\n", t.url)

	defer wg.Done()

	res, err := http.Get(t.url)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Print(res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
		return
	}

	links := []string{}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if exist && validateURI(href) {
			links = append(links, href)
		}
	})

	fmt.Printf("Links: %v\n", links)

	if t.depth > 0 {
		for _, l := range links {
			newTask := task{l, t.depth - 1}

			wg.Add(1)
			go crawl(newTask)
		}
	}

	count++
}
