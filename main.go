package main

import (
	"fmt"
	scrape "main/scrape"
)

func main() {
	err := scrape.ScrapeStream("")
	if err != nil {
		fmt.Println(err)
	}
}
