package sitemap_builder

import (
	"errors"
	"fmt"
	"net/http"
	"log"
	"os"


	"github.com/html_link_parser/parser"
)

type visitedQueue []string

func Build(url string) []string {
	urlList := []string{}
	urlVisitTracker := make(map[string]bool)
	var urlQueue visitedQueue
	urlQueue[0] = url

	for len(urlQueue) > 0 {
		currUrl = urlQueue.dequeue()

		createSitemapList(currUrl, &urlList, &urlQueue, urlVisitTracker)
	}

	return urlList
}

func (urlQueue *visitedQueue) dequeue() string {
	url := (*urlQueue)[0]
	
	if len(*urlQueue) == 1 {
		*urlQueue = (*urlQueue)[:0]
	} else {
		*urlQueue = (*urlQueue)[1:]
	}

	return url
}

func (urlQueue *visitedQueue) addUrl(url string) {
	*urlQueue = append(*urlQueue, url)
}

// func Build(url string) {
	// // urlList := []string{}
	// resp, err := http.Get(url)

	// if err != nil {
		// log.Fatalln(err)
		// os.Exit(1)
	// }
	// defer resp.Body.Close()

	// err = validateHTMLResponse(resp)
	// if err != nil {
		// log.Fatalln(err)
		// os.Exit(1)
	// }

	// links, err := parser.ProcessHTML(resp.Body)

	// if err != nil {
		// log.Fatalln(err)
	// }

	// fmt.Println(links)
// }

func createSitemapList(url string, urlList *[]string, urlQueue *visitedQueue, map[string]bool visitTracker) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if validateHTMLResponse(resp) == false {
		return
	}

	links, err := parser.ProcessHTML(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	links = filterLinks(links)

	for _, link := range(links) {
		if _, ok := visitTracker[]; !ok {
			visitTracker[link] = true
			*urlList = append(*urlList, link)
			visitedQueue.addUrl(link)
		}
	}

	fmt.Println(links)
}

// func validateHTMLResponse(resp *http.Response) error {
	// content_type := resp.Header.Get("Content-Type")
	// fmt.Println(content_type)
	// if content_type[:9] != "text/html" {
		// return errors.New(fmt.Sprintf("Content-Type is not 'text/html', it is %s", content_type))
	// }

	// return nil
// }

func validateHTMLResponse(resp *http.Response) bool {
	content_type := resp.Header.Get("Content-Type")

	if content_type[:9] != "text/html" {
		return false
	}

	return true
}

func filterLinks(links parser.Links) parser.Links {


	return nil
}