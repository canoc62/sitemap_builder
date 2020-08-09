package sitemap_builder

import (
	"errors"
	"fmt"
	"net/http"
	"log"
	"os"


	"github.com/html_link_parser/parser"
)

func Build(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = validateHTMLResponse(resp)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	links, err := parser.ProcessHTML(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(links)
}

func validateHTMLResponse(resp *http.Response) error {
	content_type := resp.Header.Get("Content-Type")
	fmt.Println(content_type)
	if content_type[:9] != "text/html" {
		return errors.New(fmt.Sprintf("Content-Type is not 'text/html', it is %s", content_type))
	}

	return nil
}