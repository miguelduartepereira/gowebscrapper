package main

import "net/url"

type PageData struct {
    URL            string   `json:"url"`
    Heading        string   `json:"heading"`
    FirstParagraph string   `json:"first_paragraph"`
    OutgoingLinks  []string `json:"outgoing_links"`
    ImageURLs      []string `json:"image_urls"`
}


func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}
	outgoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		return PageData{}
	}
	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		return PageData{}
	}
	return PageData{
		URL: pageURL,
		Heading: getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks: outgoingLinks,
		ImageURLs: imageURLs,
	}
}