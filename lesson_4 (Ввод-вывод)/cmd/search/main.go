package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"
	"parser/pkg/crawler"
	"parser/pkg/crawler/spider"
	"parser/pkg/index"
	"sort"
)

const filename = "docs"

func main() {
	spider := spider.New()
	arr := []string{"https://go.dev", "https://golang.org"}
	docs := []crawler.Document{}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		docs, err = mergeScan(spider, arr, 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Create(filename)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = store(f, docs)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		docs, err = load(f)
		if err != nil {
			fmt.Println(err)
			return
		}
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
		if ind, doc, ok := searchDocs(docs, id); ok {
			fmt.Printf("found at index %d: %v\n", ind, doc)
		}
	}
}

func store(w io.Writer, docs []crawler.Document) error {
	enc := gob.NewEncoder(w)
	if err := enc.Encode(docs); err != nil {
		return err
	}
	return nil
}

func load(r io.Reader) (docs []crawler.Document, err error) {
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&docs); err != nil {
		return nil, err
	}
	return docs, nil
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

func searchDocs(docs []crawler.Document, id int) (int, crawler.Document, bool) {
	idx := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })
	if idx < len(docs) && docs[idx].ID == id {
		return idx, docs[idx], true
	}
	return 0, crawler.Document{}, false
}
