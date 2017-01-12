package processors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func GetUrlTitle(urlAddr string) (string, error) {
	var title string
	var err error

	_, urlErr := url.Parse(urlAddr)
	if urlErr != nil {
		return title, urlErr
	}

	res, err := http.Get(urlAddr)

	if err != nil {
		return title, err
	}

	if res.StatusCode != http.StatusOK {
		return title, errors.New(fmt.Sprintf("Invalid status on request %v", res.StatusCode))
	} else {
		htmlStr, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return title, err
		}
		doc, err := html.Parse(strings.NewReader(string(htmlStr)))
		if err != nil {
			return title, err
		}

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "title" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					title = c.Data
					break
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}

	return title, err
}
