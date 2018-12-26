package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	bytes, err := ioutil.ReadAll(res.Body)
	str := string(bytes)

	fmt.Printf("Body: %s\n", str)

}
