package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	secret := "0f3c8f7307814b7b9475cda0e45fd1aa"
	nonce := getNonce()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) //int64 to string

	fmt.Println(nonce, secret, timestamp)

	arr := []string{nonce, secret, timestamp}
	sort.Strings(arr) //字典升序
	bs := []byte(strings.Join(arr, ""))
	arr20 := sha1.Sum(bs)

	fmt.Println(hex.EncodeToString(arr20[:]))
}

func getNonce() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
