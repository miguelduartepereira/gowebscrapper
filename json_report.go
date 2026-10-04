package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	pagesData := make([]PageData, 0, len(pages))
	for _, k := range keys {
		pagesData = append(pagesData, pages[k])
	}

	data, err := json.MarshalIndent(pagesData, "", "  ")
	if err != nil {
		return err
	}

	os.WriteFile(filename, data, 0644)
	return nil
}