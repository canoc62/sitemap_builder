package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sitemap_builder/sitemap_builder"
)

func main() {
	var urlArg = flag.String("url", "https://spacejam.com/", "usage: url to generate sitemap for")
	var depthArg = flag.Int("depth", 3, "usage: number of levels of links to generate sitemap")
	flag.Parse()

	builder, err := sitemap_builder.NewBuilder(*urlArg, *depthArg)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	go spinner(100 * time.Millisecond)

	urlList := builder.Build()

	fmt.Println(urlList)
	

	// pass url to sitemap builder
	// make request
	// process html to get links
	// go through links, bfs style?, build sitemap
	// generate xml based on sitemap
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}