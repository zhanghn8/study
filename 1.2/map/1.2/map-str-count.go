package main

import (
	"fmt"
	"sort"
)

type frequency struct {
	char string
	fre  int
}

// 计算字符出现频率，并按照频率降序排序
func frequencies(s string) []frequency {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	a := make([]frequency, 0, len(m))
	for c, f := range m {
		a = append(a, frequency{char: c, fre: f})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].fre > a[j].fre })
	return a
}
func main() {
	str := "hiilovegogogosdfsodfosdof"
	fmt.Println(str)
	f := frequencies(str)
	fmt.Println(f)
}
