package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	words := []string{"aaaasd", "a", "aab", "aaabcd",
		"ef", "cssssssd", "fdz", "kf", "zc",
		"lklklklklklklklkl", "l"}

	//Search "a"

	filtered1 := Filter(words, func(word string) bool {
		return strings.HasPrefix(word, "a")
	})

	//Slicer
	sort.Slice(filtered1, func(i, j int) bool {
		return len(filtered1[i]) > len(filtered1[j])
	})

	words = words[4:]
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	//fmt.Println(filtered1)

	//Append

	words = append(filtered1, words...)
	fmt.Println(words)

}

//Filter func with for

func Filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
