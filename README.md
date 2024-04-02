<div align="center">
    <br />
    <a href="https://github.com/Oudwins/tailwind-merge-go">
        <img src="https://raw.githubusercontent.com/Oudwins/tailwind-merge-go/master/assets/logo.svg" alt="tailwind-merge-go" height="150px" />
    </a>
</div>

# tailwind-merge-go - Tailwind Merge For Golang

<a href="https://pkg.go.dev/github.com/Oudwins/tailwind-merge-go"><img src="https://pkg.go.dev/badge/github.com//github.com/Oudwins/tailwind-merge-go.svg" alt="Go Reference" /></a>
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/Oudwins/tailwind-merge-go)](https://goreportcard.com/report/github.com/Oudwins/tailwind-merge-go)

Utility function to efficiently merge Tailwind CSS classes in Golang without style conflicts. This library aims to be as close as possible to a 1:1 copy of the original [dcastil/tailwind-merge](https://github.com/dcastil/tailwind-merge/) library written in javascript.

```go
import (
	"fmt"

	twmerge "github.com/Oudwins/tailwind-merge-go"
)

func main() {

	// example usage
	c := twmerge.Merge("px-4 px-10 p-1")
	fmt.Println(c) // "p-1"
}
```

- Supports Tailwind v3.0 up to v3.4
- Support for extending the default configuration
- Support for providing your own caching solution
- Its in 0.1.0, can I use it? Sure! I will personally be deploying this to prod. It's only in pre 1.0 because there some extra features I want to add before the 1.0 release (see roadmap)

## [Why use it?](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/what-is-it-for.md)

- See [tailwind-merge](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/what-is-it-for.md)
- Or Watch this amazing video on it

[Watch this introduction video from Simon Vrachliotis (@simonswiss) ↓ ![The "why" behind tailwind-merge](https://img.youtube.com/vi/tfgLd5ZSNPc/maxresdefault.jpg)](https://www.youtube.com/watch?v=tfgLd5ZSNPc (Watch YouTube video "Tailwind-Merge Is Incredibly Useful — And Here's Why!"))

## [Limitations](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/limitations.md)

- See [tailwind-merge](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/limitations.md)

## Advanced Examples

You might also want to check out the advanced example at `/cmd/examples/advanced`

### Provide Your Own or Extend Default Config

```go
import (
		// Note the import path here is different from the default path. This is so you have access to all the custom functions, structs, etc that are used to build the twmerge config
		twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
)
var TwMerger twmerge.TwMergeFn
func main() {
	// get the default config
	config := twmerge.MakeDefaultConfig()

	// do your modifications here

	// create the merger
	TwMerger = twmerge.CreateTwMerge(config, nil) // config, cache (if nil default will be used)


	// example usage
	m := TwMerger("px-4 px-10", "p-20")
	fmt.Println(m) // output: "p-20"
}
```

### Provide your own Cache

The default cache is a LRU Cache and should be acceptable for most use cases. However, you might want to provide your own cache or modify the default creation parameters. Your cache must implement the interface defined at `/pkg/cache/cache.go`

```go
type ICache interface {
	Get(string) string
	Set(string, string) // key, value
}
```

Here is an example of manually creating the default cache with a custom max capacity

```go
import (
		twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
		lru "github.com/Oudwins/tailwind-merge-go/pkg/lru"
)
var TwMerger twmerge.TwMergeFn
func main() {
	customCapacity := 10000
	cache := lru.make(customCapacity)


	// create the merger
	TwMerger = twmerge.CreateTwMerge(nil, cache) // config, cache (if nil default will be used)

	// example usage
	m := TwMerger("px-4 px-10", "p-20")
	fmt.Println(m) // output: "p-20"
}
```

## Contributing

Checkout the [contributing docs](./CONTRIBUTING.md)

## Roadmap

- Improve cache concurrent performance by locking on a per key basis -> https://github.com/EagleChen/mapmutex
- Build the class map on initialization and have a simple config style
- replace regex with more performant solution
- Move arbitrary value delimeters '[' & ']' to config somehow?
- Plugins & easy plugin api.

## Acknowledgments

- Credit for all the hard work goes to [dcastil/tailwind-merge](https://github.com/dcastil/tailwind-merge/).
  - For the tests I used
  - For the approach and the code. I mostly translated from js to go
  - For the logo
- Big thank you to [tylantz/go-tailwind-merge/](https://github.com/tylantz/go-tailwind-merge/tree/main) for pushing me to finally do this by writing a very interesting version of this same idea (I encourage you to check it out) and for the code to generate a go test file based on tailwind-merge's tests
