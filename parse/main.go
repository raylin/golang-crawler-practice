package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	url := flag.String("url", "", "Target URL")
	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("URL: %s\n", *url)

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Links")
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if exist && validateURI(href) {
			fmt.Println(href)
		}
	})
}

func validateURI(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}
