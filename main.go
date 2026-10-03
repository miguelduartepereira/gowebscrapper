package main

import (
	"fmt"
	"os"
)


func main() {
	args := os.Args[1:]
	if len(args) < 1{
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	if len(args) == 1{
		fmt.Printf("starting crawl of: %s\n", args[0])
		pages := make(map[string]int)
		crawlPage(args[0], args[0], pages)
		for normalizedURL, count := range pages{
			fmt.Printf("found page: %s, count: %d\n", normalizedURL, count)
		}
	}


}