package main

import (
	"flag"

	"github.com/sitemap_builder/sitemap_builder"
)

func main() {
	var url_arg = flag.String("url", "https://spacejam.com/", "usage: url to generate sitemap for")
	flag.Parse()

	sitemap_builder.Build(*url_arg)
	

	// pass url to sitemap builder
	// make request
	// process html to get links
	// go through links, bfs style?, build sitemap
	// generate xml based on sitemap
}