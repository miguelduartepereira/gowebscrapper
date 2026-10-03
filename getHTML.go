package main

import (
	"fmt"
	"net/http"
	"io"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400{
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)	
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("Content-Type is not text/html: %s", resp.Header.Get("Content-Type"))
	}
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil


}