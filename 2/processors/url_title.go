package processors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
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
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return title, err
		} else {
			r, _ := regexp.Compile(`<title>.*<\/title>`)
			matches := r.FindAll(bytes, 1)
			if len(matches) > 0 {
				return string(matches[0]), nil
			}
		}
	}

	return title, err
}
