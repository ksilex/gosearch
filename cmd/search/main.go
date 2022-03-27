package main

import (
	"flag"
	"fmt"
	"parser/pkg/crawler"
	"parser/pkg/crawler/spider"
	"strings"
)

func main() {
	spider := spider.New()
	arr := []string{"https://go.dev", "https://golang.org"}
	docs, err := mergeScan(spider, arr, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	search := flag.String("s", "", "search for keyword")
	flag.Parse()

	if len(*search) > 1 {
		searchDocs(docs, *search)
	}
}

func mergeScan(s *spider.Service, urls []string, lvl int) ([]crawler.Document, error) {
	res, err := s.Scan(urls[0], lvl)

	if err != nil {
		return nil, err
	}

	for _, url := range urls[1:] {
		i, err := s.Scan(url, lvl)

		if err != nil {
			return nil, err
		}

		res = append(res, i...)
	}

	return res, nil
}

func searchDocs(docs []crawler.Document, word string) {
	res := make([]crawler.Document, 0)
	for _, doc := range docs {
		if strings.Contains(doc.URL, word) || strings.Contains(doc.Title, word) {
			res = append(res, doc)
		}
	}

	fmt.Printf("search result: %v\n", res)
}
