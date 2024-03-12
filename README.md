# tailwind-merge-go

Utility function to efficiently merge Tailwind CSS classes in Golang without style conflicts. This library aims to be as close as possible to a 1:1 copy of the original [dcastil/tailwind-merge](https://github.com/dcastil/tailwind-merge/) library written in javascript.

```go
import (
	"fmt"

	"github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
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

## [Why use it?](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/what-is-it-for.md)

- See [tailwind-merge](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/what-is-it-for.md)

## [Limitations](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/limitations.md)

- See [tailwind-merge](https://github.com/dcastil/tailwind-merge/blob/v2.2.1/docs/limitations.md)

## Roadmap

- Write contributing docs
- Improve current docs
- Improve cache concurrent performance by locking on a per key basis -> https://github.com/EagleChen/mapmutex
- Split code into multiple pkgs so in the twmerge pkg there is only the Merge & CreateTailwindMerge functions
- Build the class map on initialization and have a simple config style
- replace regex with more performant solution
- Move arbitrary value delimeters '[' & ']' to config somehow?

## Acknowledgments

- Credit for all the hard work goes to [dcastil/tailwind-merge](https://github.com/dcastil/tailwind-merge/).
  - For the tests I used
  - For the approach and the code. I mostly translated from js to go
- Big thank you to [tylantz/go-tailwind-merge/](https://github.com/tylantz/go-tailwind-merge/tree/main) for pushing me to finally do this by writing a very interesting version of this same idea (I encourage you to check it out) and for the code to generate a go test file based on tailwind-merge's tests
