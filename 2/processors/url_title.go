package processors

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type html struct {
	Title title `xml:"title"`
}
type title struct {
	Content string `xml:",innerxml"`
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
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return title, err
		} else {
			h := html{}
			err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
			if err != nil {
				fmt.Println("Error parsing html page", err)
				return title, err
			}

			fmt.Println(h.Title.Content)
			return h.Title.Content, nil
			// r, _ := regexp.Compile(`<title>.*<\/title>`)
			// matches := r.FindAll(bytes, 1)
			// if len(matches) > 0 {
			// 	return string(matches[0]), nil
			// }
		}
	}

	return title, err
}
