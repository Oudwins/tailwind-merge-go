package twmerge

import (
	"strings"

	lru "github.com/Oudwins/tailwind-merge-go/pkg/cache"
)

// create the config (just gets the config passed in)

// create the config utils
// LRU cache
// split modifiers
// -> for things like hover:bg-x
// class utils
// -> for splitting classes

// cache get & set

// merge fn
// 1. check cache
// 2. mergeClassList
// 3. set cache

// should this also take a cache directly?
func CreateTwMerge(config *TwMergeConfig, cache lru.Cache) func(args ...string) string {
	if config == nil {
		config = MakeDefaultConfig()
	}
	if cache == nil {
		cache = lru.Make(config.MaxCacheSize)
	}

	splitModifiers := MakeSplitModifiers(config)

	getClassGroupId := MakeGetClassGroupId(config)

	mergeClassList := MakeMergeClassList(config, splitModifiers, getClassGroupId)

	return func(args ...string) string {
		classList := strings.Join(args, " ")
		cached := cache.Get(classList)
		if cached != "" {
			return cached
		}
		// check if in cache
		merged := mergeClassList(classList)
		cache.Set(classList, merged)
		return merged
	}
}

var Merge = CreateTwMerge(nil, nil)
