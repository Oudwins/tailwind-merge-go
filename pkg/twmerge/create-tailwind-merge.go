package twmerge

import (
	"strings"

	cache "github.com/Oudwins/tailwind-merge-go/pkg/cache"
	lru "github.com/Oudwins/tailwind-merge-go/pkg/lru"
)

type TwMergeFn func(args ...string) string

func CreateTwMerge(config *TwMergeConfig, cache cache.ICache) TwMergeFn {

	var fnToCall TwMergeFn
	var splitModifiers SplitModifiersFn
	var getClassGroupId GetClassGroupIdfn
	var mergeClassList func(classList string) string

	merger := func(args ...string) string {
		classList := strings.TrimSpace(strings.Join(args, " "))
		if classList == "" {
			return ""
		}
		cached := cache.Get(classList)
		if cached != "" {
			return cached
		}
		// check if in cache
		merged := mergeClassList(classList)
		cache.Set(classList, merged)
		return merged
	}

	init := func(args ...string) string {
		if config == nil {
			config = MakeDefaultConfig()
		}
		if cache == nil {
			cache = lru.Make(config.MaxCacheSize)
		}

		splitModifiers = MakeSplitModifiers(config)

		getClassGroupId = MakeGetClassGroupId(config)

		mergeClassList = MakeMergeClassList(config, splitModifiers, getClassGroupId)

		fnToCall = merger
		return fnToCall(args...)
	}

	fnToCall = init
	return func(args ...string) string {
		return fnToCall(args...)
	}
}

var Merge = CreateTwMerge(nil, nil)
