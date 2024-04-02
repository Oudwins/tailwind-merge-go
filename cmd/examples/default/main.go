package main

import (
	"fmt"

	twmerge "github.com/Oudwins/tailwind-merge-go"
)

func main() {
	// mainly used for manual tests

	// example usage
	m := twmerge.Merge("px-4 px-10", "p-20")
	fmt.Println(m)
}
