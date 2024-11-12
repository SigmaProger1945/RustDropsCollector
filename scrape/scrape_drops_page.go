package scrape

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeRustDropsPage() ([]string, error) {
	/*rust_page_cookies := []Cookies{
		{
			Name:  ".AspNetCore.Session",
			Value: "CfDJ8PHckxzFqzFBjGGqMKqwNnBQ9w34Bl4/UalNO1ORGkC62fKiKEfbLuGD712RkXOOZlM7Hs7UX+4WyEoadsUpXl3Nh6rnMAiZHu6vgWT/ekygLjz/Qc9f34cQ8jdMXORMXqH/cDQO4AOxwg6nV9/z+pKpCtZQSwA5JSufCH6i0xAf",
		},
	}*/
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://twitch.facepunch.com/#drops", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := CreateGoqueryDoc(resp)
	if err != nil {
		return nil, err
	}
	var avalibleStreams []string
	doc.Find(".drop-box.is-live").Each(func(_ int, s *goquery.Selection) {
		s.Find(".header-container").Each(func(i int, q *goquery.Selection) {
			q.Find("a").Each(func(_ int, o *goquery.Selection) {
				href, exists := o.Attr("href")
				if exists {
					avalibleStreams = append(avalibleStreams, href)
				}
			})
		})
	})
	return avalibleStreams, nil

}
