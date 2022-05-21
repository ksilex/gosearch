package index

import (
	"strings"
)

type Service struct {
	storage map[string][]int
}

func New() *Service {
	s := &Service{
		storage: make(map[string][]int),
	}
	return s
}

func (serv *Service) Add(s string, i int) {
	words := strings.Fields(s)
	for _, word := range words {
		serv.storage[word] = append(serv.storage[word], i)
	}
}

func (i *Service) Search(s string) []int { return i.storage[s] }
