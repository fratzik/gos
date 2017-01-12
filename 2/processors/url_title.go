package processors

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

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

		// b, err := ioutil.ReadAll(res.Body)
		z := html.NewTokenizer(res.Body)

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				// Reaching the end of document
				return title, err
			case tt == html.StartTagToken:
				t := z.Token()

				isTitle := t.Data == "title"
				if isTitle {
					fmt.Println("We found a link!")
					fmt.Printf("%v", t)
				}
			}
		}
		// if err != nil {
		// 	return title, err
		// } else {
		// h := html{}
		// err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
		// if err != nil {
		// 	fmt.Println("Error parsing html page", err)
		// 	return title, err
		// }

		// fmt.Println(h.Title.Content)
		// return h.Title.Content, nil
		// r, _ := regexp.Compile(`<title>.*<\/title>`)
		// matches := r.FindAll(bytes, 1)
		// if len(matches) > 0 {
		// 	return string(matches[0]), nil
		// }
		// }
	}

	return title, err
}
