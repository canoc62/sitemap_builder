package main

import (
	"flag"
	"log"
	"os"

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

	sitemapHrefs := builder.Build()
	sitemap_builder.ToXml(sitemapHrefs)
}