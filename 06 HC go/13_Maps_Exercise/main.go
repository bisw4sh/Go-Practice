package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make( map[string]int )
	keys := strings.Fields(s)

	for _ , item := range keys {
		_ , ok := m[item]
			
	if ok {
		m[item] = m[item] + 1
	}else {
		m[item] = 1
	}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
