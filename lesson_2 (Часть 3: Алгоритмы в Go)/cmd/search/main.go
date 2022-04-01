package main

import (
	"flag"
	"fmt"
	"parser/pkg/crawler"
	"parser/pkg/crawler/spider"
	"parser/pkg/index"
	"sort"
)

func main() {
	spider := spider.New()
	arr := []string{"https://go.dev", "https://golang.org"}
	docs, err := mergeScan(spider, arr, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	idx := index.New()
	for _, doc := range docs {
		idx.Add(doc.Title, doc.ID)
	}

	search := flag.String("s", "", "search for keyword")
	flag.Parse()

	var docIds []int
	if len(*search) > 1 {
		docIds = idx.Search(*search)
	}

	for _, id := range docIds {
		searchDocs(docs, id)
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

	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	return res, nil
}

func searchDocs(docs []crawler.Document, id int) {
	idx := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })
	if idx < len(docs) && docs[idx].ID == id {
		fmt.Printf("found at index %d: %v\n", idx, docs[idx])
	}
}
