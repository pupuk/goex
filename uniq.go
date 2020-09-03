package main

import "fmt"

func main() {
	arr := [...]int{2, 5, 8, 3, 5, 2, 5, 9, 10, 3, 3, 5, 8, 3, 100, 3, 200, 8, 1, 2, 5, 6, 2, 5, 8, 3, 5, 2, 5, 9, 10, 3, 3, 5, 8, 3, 100, 3, 200, 8, 1, 2, 5, 6, 2, 5, 8, 3, 5, 2, 5, 9, 10, 3, 3, 5, 8, 3, 100, 3, 200, 8, 1, 2, 5, 6, 2, 5, 8, 3, 5, 2, 5, 9, 10, 3, 3, 5, 8, 3, 100, 3, 200, 8, 1, 2, 5, 6, 2, 5, 8, 3, 5, 2, 5, 9, 10, 3, 3, 5, 8, 3, 100, 3, 200, 8, 1, 2, 5, 6}

	r := []int{}

	for _, v := range arr {
		if !exist(v, r) {
			r = append(r, v)
		}
	}

	fmt.Println(r)
}

func exist(search int, arr []int) bool {
	if arr == nil {
		return false
	}

	for _, v := range arr {
		if search == v {
			return true
		}
	}
	return false
}
