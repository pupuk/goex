package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// MD5文件
	// f, err := os.Open("md5file.txt")
	f, err := os.Open("md5.zip")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Println(f) //&{0xc00009a120} 是内存地址，是指针

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))

	r := md5.Sum([]byte("111111"))
	fmt.Printf("%x\n", r)

}
