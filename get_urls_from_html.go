package main

import (
	"strings"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		parsedURL, err := url.Parse(href)
		if err != nil {
			return
		}

		if parsedURL.Host == "" {
			parsedURL = baseURL.ResolveReference(parsedURL)
		}

		urls = append(urls, parsedURL.String())
	})

	return urls, nil
}