package twmerge

type TwMergeConfig struct {
	// defaults should be good enough
	// hover:bg-red-500 -> :
	ModifierSeparator rune
	// bg-red-500 -> -
	ClassSeparator rune
	// !bg-red-500 -> !
	ImportantModifier rune
	// used for bg-red-500/50 (50% opacity) -> /
	PostfixModifier rune
	// optional
	Prefix string

	// CACHE
	MaxCacheSize int
	// I couldn't figure out what they use the theme for
	// Theme TwTheme

	// This is a large map of all the classes and their validators -> see default-config.go
	ClassGroups ClassPart

	// class group with conflict + conflicting groups -> if "p" is set all others are removed
	// p: ['px', 'py', 'ps', 'pe', 'pt', 'pr', 'pb', 'pl']
	ConflictingClassGroups ConflictingClassGroups
}

// type TwTheme struct {
// }

type ClassGroupValidator struct {
	Fn           func(string) bool
	ClassGroupId string
}
type ClassPart struct {
	NextPart     map[string]ClassPart
	Validators   []ClassGroupValidator
	ClassGroupId string
}

type ConflictingClassGroups map[string][]string
