package main

import (
	"flag"
	"fmt"
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
}
