package sitemap_builder

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/html_link_parser/parser"
)

type visitedQueue []string

type Builder struct {
	BaseUrl string
	UrlList *[]string
	seen map[string]bool
	depth int
}

func NewBuilder(url string, depth int) (Builder, error) {
	baseUrl, err := getBaseUrl(url)

	if err != nil {
		return Builder{}, err
	}

	urlList := []string{}
	seen := make(map[string]bool)

	return Builder { baseUrl, &urlList, seen, depth }, nil
}

func getBaseUrl(urlArg string) (string, error) {
	resp, err := http.Get(urlArg)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	content_type := resp.Header.Get("Content-Type")
	if validateHTMLResponse(content_type) == false {
		return "", errors.New(fmt.Sprintf("Content type of requested URL is %s", content_type))
	}

	reqUrl := resp.Request.URL
	baseUrl := &url.URL { Scheme: reqUrl.Scheme, Host: reqUrl.Host }
	base := baseUrl.String()
	
	return base, nil
}

func (builder Builder) Build() []string {
	var urlQueue visitedQueue
	urlQueue = append(urlQueue, builder.BaseUrl)

	counter := builder.depth

	for len(urlQueue) > 0 && counter > 0 {
		counter--
		levelLength := len(urlQueue)
		for levelLength > 0 {
			currUrl := urlQueue.dequeue()
			builder.createSitemapList(currUrl, &urlQueue)
			levelLength--
		}
	}

	return *builder.UrlList
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

func (builder Builder) createSitemapList(url string, urlQueue *visitedQueue) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if !validateHTMLResponse(resp.Header.Get("Content-Type")) {
		log.Fatalln(fmt.Sprintf("Invalid Content-Type for resource at: %s, will skip processing.", url))
		return
	}

	links, err := parser.ProcessHTML(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	hrefs := filterLinks(&links, builder.BaseUrl)

	for _, href := range(hrefs) {
		if _, ok := builder.seen[href]; !ok {
			builder.seen[href] = true
			*builder.UrlList = append(*builder.UrlList, href)
			urlQueue.addUrl(href)
		}
	}
}

func validateHTMLResponse(respHeader string) bool {
	return strings.HasPrefix(respHeader, "text/html")
}

func filterLinks(links *parser.Links, baseUrl string) []string {
	hrefs := []string{}

	for _, link := range(*links) {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			hrefs = append(hrefs, baseUrl+link.Href)
		case !strings.HasPrefix(link.Href, "http"):
			hrefs = append(hrefs, baseUrl+"/"+link.Href)
		case strings.HasPrefix(link.Href, baseUrl):
			hrefs = append(hrefs, link.Href)
		}
	}

	return hrefs
}