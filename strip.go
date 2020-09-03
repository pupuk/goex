package main

import (
	"fmt"

	strip "github.com/grokify/html-strip-tags-go" // => strip
)

func main() {
	html := `<button type="button" class="_1OyPqC _3Mi9q9 _2WY0RL _1YbC5u"><span>赞赏支持</span></button>`
	stripped := strip.StripTags(html)
	fmt.Println(stripped)
}
