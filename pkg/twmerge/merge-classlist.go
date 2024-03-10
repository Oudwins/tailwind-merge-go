package twmerge

import (
	"regexp"
	"slices"
	"strings"
)

const SPLIT_CLASSES_REGEX = `\s+`

var splitPattern = regexp.MustCompile(SPLIT_CLASSES_REGEX)

func MakeMergeClassList(conf *TwMergeConfig, splitModifiers SplitModifiersFn, getClassGroupId GetClassGroupIdfn) func(classList string) string {
	return func(classList string) string {
		classes := splitPattern.Split(strings.TrimSpace(classList), -1)
		unqClasses := make(map[string]string, len(classes))
		resultClassList := ""

		for _, class := range classes {
			baseClass, modifiers, hasImportant, maybePostfixModPosition := splitModifiers(class)

			// there is a postfix modifier -> text-lg/8
			if maybePostfixModPosition != -1 {
				baseClass = baseClass[:maybePostfixModPosition]
			}
			isTwClass, groupId := getClassGroupId(baseClass)
			if !isTwClass {
				resultClassList += class + " "
				continue
			}
			// we have to sort the modifiers bc hover:focus:bg-red-500 == focus:hover:bg-red-500
			modifiers = SortModifiers(modifiers)
			if hasImportant {
				modifiers = append(modifiers, "!")
			}
			unqClasses[groupId+strings.Join(modifiers, string(conf.ModifierSeparator))] = class

			conflicts := conf.ConflictingClassGroups[groupId]
			if conflicts == nil {
				continue
			}
			for _, conflict := range conflicts {
				// erase the conflicts with the same modifiers
				unqClasses[conflict+strings.Join(modifiers, string(conf.ModifierSeparator))] = ""
			}
		}

		for _, class := range unqClasses {
			if class == "" {
				continue
			}
			resultClassList += class + " "
		}
		return strings.TrimSpace(resultClassList)
	}

}

/**
 * Sorts modifiers according to following schema:
 * - Predefined modifiers are sorted alphabetically
 * - When an arbitrary variant appears, it must be preserved which modifiers are before and after it
 */
func SortModifiers(modifiers []string) []string {
	if modifiers == nil || len(modifiers) < 2 {
		return modifiers
	}

	unsortedModifiers := []string{}
	sorted := make([]string, len(modifiers))

	for _, modifier := range modifiers {
		isArbitraryVariant := modifier[0] == '['
		if isArbitraryVariant {
			slices.Sort(unsortedModifiers)
			sorted = append(sorted, unsortedModifiers...)
			sorted = append(sorted, modifier)
			unsortedModifiers = []string{}
			continue
		}
		unsortedModifiers = append(unsortedModifiers, modifier)
	}

	slices.Sort(unsortedModifiers)
	sorted = append(sorted, unsortedModifiers...)

	return sorted
}
