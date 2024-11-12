package scrape

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Cookies struct {
	Name  string
	Value string
}

func AddCookies(req *http.Request, cookies []Cookies) *http.Request {
	for _, cookie := range cookies {
		req.AddCookie(&http.Cookie{Name: cookie.Name, Value: cookie.Value})
	}
	return req
}

func CreateGoqueryDoc(resp *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return doc, nil
}
