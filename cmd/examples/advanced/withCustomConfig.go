package main

import (
	"fmt"

	lru "github.com/Oudwins/tailwind-merge-go/pkg/lru"
	twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
)

var TwMerger twmerge.TwMergeFn

func main() {
	// get the default config
	config := twmerge.MakeDefaultConfig()

	// make cache
	cache := lru.Make(10000)

	// do your modifications here

	// create the merger
	TwMerger = twmerge.CreateTwMerge(config, cache)

	// example usage
	m := TwMerger("px-4 px-10", "p-20")
	fmt.Println(m) // output: "p-20"
}
