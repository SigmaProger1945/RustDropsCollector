package scrape

import (
	"fmt"
	"io"
	"net/http"
)

func ScrapeStream(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	cookies := []Cookies{
		{
			Name:  "auth-token",
			Value: "yhu7mg9whbpp47bwfrvlwmj5e1es52",
		},
	}
	client := &http.Client{}
	req = AddCookies(req, cookies)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buffer := make([]byte, 8192)
	StreamIsLive := true

	for StreamIsLive {

		fmt.Println("Connected to stream, starting to read...")
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Stream ended.")
				StreamIsLive = false
			}
			fmt.Println("Error reading stream:", err)
		}

		fmt.Printf("Received %d bytes of data.\n", n)
		fmt.Print(string(buffer[:n]))
	}
	return nil
}
