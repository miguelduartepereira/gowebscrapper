package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error){
	u, err := url.Parse(inputURL)
	
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	u.Path = strings.TrimSuffix(u.Path, "/")
	return fmt.Sprintf("%s%s", strings.ToLower(u.Host), strings.ToLower(u.Path)), nil

}