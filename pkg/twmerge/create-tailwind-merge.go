package twmerge

import (
	"strings"

	lru "github.com/Oudwins/tailwind-merge-go/pkg/cache"
)

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
