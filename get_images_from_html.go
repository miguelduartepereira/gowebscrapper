package main
import (
	"strings"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)
func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var images []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}

		parsedURL, err := url.Parse(src)
		if err != nil {
			return
		}

		if parsedURL.Host == "" {
			parsedURL = baseURL.ResolveReference(parsedURL)
		}

		images = append(images, parsedURL.String())
	})

	return images, nil
}