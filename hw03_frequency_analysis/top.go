package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type words struct {
	word string
	qty  int
}

var re = regexp.MustCompile("[ \t\n]+")

func Top10(text string) []string {
	if text == "" {
		return make([]string, 0)
	}

	s := re.Split(text, -1)

	m := make(map[string]int)
	for _, v := range s {
		m[v]++
	}

	w := make([]words, 0)
	for key, value := range m {
		w = append(w, words{key, value})
	}

	sort.Slice(w, func(i, j int) bool {
		return (w[i].qty == w[j].qty && w[i].word < w[j].word) || w[i].qty > w[j].qty
	})

	res := make([]string, 0)

	for i := 0; i < 10; i++ {
		res = append(res, w[i].word)
	}

	return res
}
