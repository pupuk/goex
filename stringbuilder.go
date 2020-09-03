package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	for i := 0; i < 100000; i++ {
		fmt.Fprint(&b, "A")
	}

	var str = b.String()
	fmt.Println(str)
}
