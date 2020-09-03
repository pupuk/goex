package main

import (
	"bufio"
	"fmt"
	"io"
	_ "io/ioutil"
	"os"
	"strings"
)

const cap int = 2

func main() {
	// b, e := ioutil.ReadFile("./md5file.txt")
	// if e != nil {
	// 	fmt.Println("read file error")
	// 	return
	// }
	// fmt.Println(string(b))

	f, err := os.Open("./check.log")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer f.Close()
	r := bufio.NewReader(f)

	//本程序写的是，有cap个连续的请求，出现2+次的timeout的，计数（count）
	count := 0
	tmpCount := 0

	for {
		b, _, c := r.ReadLine()
		if c == io.EOF {
			break
		}

		lineString := string(b)

		if strings.Index(lineString, "timeout") != -1 {
			tmpCount++
			// fmt.Println(tmpCount)
		} else {
			if tmpCount > 0 {
				tmpCount--
			}
		}

		if tmpCount == 5 {
			// fmt.Println("yes")
			count++
		}

	}

	fmt.Println("\t", count)
}
