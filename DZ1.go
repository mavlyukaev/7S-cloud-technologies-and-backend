package main

import (
	"fmt"
	"sort"
)

func Sorting(b []int, bool1 bool, bool2 bool) {
	res := []int{}
	even := []int{}
	odd := []int{}
	for i := 0; i < len(b); i++ {
		if b[i]%2 == 0 {
			even = append(even, b[i])
		} else {
			odd = append(odd, b[i])
		}
	}
	if bool1 == true {
		sort.Ints(even)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(even)))
	}
	if bool2 == true {
		sort.Ints(odd)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(odd)))
	}
	res = append(res, even...)
	res = append(res, odd...)
	fmt.Printf("%v", res)
	return
}
func main() {
	buff := []int{1, 2, 3, 4, 5, 6}
	bool1 := false
	bool2 := true
	Sorting(buff, bool1, bool2)
	return
}
