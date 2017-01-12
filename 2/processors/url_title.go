package processors

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func GetUrlTitle(urlAddr string) (string, error) {
	var title string
	var err error

	resp, err := http.Head(urlAddr)
	if err != nil {
		return title, err
	}

	if resp.ContentLength > 1*1024*1024 {
		return title, errors.New(fmt.Sprintf("Dropping request for url %s due to large size: %d", urlAddr, resp.ContentLength))
	}

	resp, err = http.Get(urlAddr)

	if resp.StatusCode != http.StatusOK {
		return title, errors.New(fmt.Sprintf("Invalid status on request %v", resp.StatusCode))
	}
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("%v\n", err)
	}

	title, ok := traverse(doc)

	if ok {
		return title, nil
	}

	return title, err
}
