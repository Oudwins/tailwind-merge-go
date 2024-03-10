package twmerge

import (
	"regexp"
	"strings"
)

type GetClassGroupIdfn func(string) (isTwClass bool, groupId string)

func MakeGetClassGroupId(conf *TwMergeConfig) GetClassGroupIdfn {
	var getClassGroupIdRecursive func(classParts []string, i int, classMap *ClassPart) (isTwClass bool, groupId string)
	getClassGroupIdRecursive = func(classParts []string, i int, classMap *ClassPart) (isTwClass bool, groupId string) {
		if i >= len(classParts) {
			if classMap.ClassGroupId != "" {
				return true, classMap.ClassGroupId
			}

			return false, ""
		}

		if classMap.NextPart != nil {
			nextClassMap := classMap.NextPart[classParts[i]]
			isTw, id := getClassGroupIdRecursive(classParts, i+1, &nextClassMap)
			if isTw {
				return isTw, id
			}
		}

		if classMap.Validators != nil && len(classMap.Validators) > 0 {
			remainingClass := strings.Join(classParts[i:], string(conf.ClassSeparator))

			for _, validator := range classMap.Validators {
				if validator.Fn(remainingClass) {
					return true, validator.ClassGroupId
				}
			}

		}
		return false, ""
	}

	var arbitraryPropertyRegex = regexp.MustCompile(`^\[(.+)\]$`)

	getGroupIdForArbitraryProperty := func(class string) (bool, string) {
		if arbitraryPropertyRegex.MatchString(class) {
			arbitraryPropertyClassName := arbitraryPropertyRegex.FindStringSubmatch(class)[1]
			property := arbitraryPropertyClassName[:strings.Index(arbitraryPropertyClassName, ":")]

			if property != "" {
				// I use two dots here because one dot is used as prefix for class groups in plugins
				return true, "arbitrary.." + property
			}
		}

		return false, ""
	}

	return func(baseClass string) (isTwClass bool, groupdId string) {

		classParts := strings.Split(baseClass, string(conf.ClassSeparator))
		// remove first element if empty for things like -px-4
		if len(classParts) > 0 && classParts[0] == "" {
			classParts = classParts[1:]
		}
		isTwClass, groupId := getClassGroupIdRecursive(classParts, 0, &conf.ClassGroups)
		if isTwClass {
			return isTwClass, groupId
		}

		return getGroupIdForArbitraryProperty(baseClass)
	}

}
