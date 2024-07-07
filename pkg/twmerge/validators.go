package twmerge

import (
	"regexp"
	"strconv"
)

var stringLengths = map[string]bool{
	"px":     true,
	"full":   true,
	"screen": true,
}

var lengthUnitRegex = regexp.MustCompile(`\d+(%|px|r?em|[sdl]?v([hwib]|min|max)|pt|pc|in|cm|mm|cap|ch|ex|r?lh|cq(w|h|i|b|min|max))|\b(calc|min|max|clamp)\(.+\)|^0$`)
var colorFnRegex = regexp.MustCompile(`^(rgba?|hsla?|hwb|(ok)?(lab|lch))\(.+\)$`)

func IsAny(_ string) bool {
	return true
}

func IsNever(_ string) bool {
	return false
}

func IsLength(val string) bool {
	if IsNumber(val) || stringLengths[val] || IsFraction(val) {
		return true
	}
	return false
}

func IsArbitraryLength(val string) bool {
	return GetIsArbitraryValue(val, "length", IsLengthOnly)
}

var arbitraryRegex = regexp.MustCompile(`(?i)^\[(?:([a-z-]+):)?(.+)\]$`)

func IsArbitraryNumber(val string) bool {
	return GetIsArbitraryValue(val, "number", IsNumber)
}

func IsArbitraryPosition(val string) bool {
	return GetIsArbitraryValue(val, "position", IsNever)
}

var sizeLabels = map[string]bool{
	"length": true, "size": true, "percentage": true,
}

func IsArbitrarySize(val string) bool {
	return GetIsArbitraryValue(val, sizeLabels, IsNever)
}

var imageLabels = map[string]bool{
	"image": true, "url": true,
}

func IsArbitraryImage(val string) bool {
	return GetIsArbitraryValue(val, imageLabels, IsImage)
}
func IsArbitraryShadow(val string) bool {
	return GetIsArbitraryValue(val, "", IsShadow)
}

func IsArbitraryValue(val string) bool {
	return arbitraryRegex.MatchString(val)
}

func IsPercent(val string) bool {
	return val[len(val)-1] == '%' && IsNumber(val[:len(val)-1])
}

func IsTshirtSize(val string) bool {
	pattern := regexp.MustCompile(`^(\d+(\.\d+)?)?(xs|sm|md|lg|xl)$`)
	return pattern.MatchString(val)
}

func IsShadow(val string) bool {
	pattern := regexp.MustCompile(`^(inset_)?-?((\d+)?\.?(\d+)[a-z]+|0)_-?((\d+)?\.?(\d+)[a-z]+|0)`)
	return pattern.MatchString(val)
}

func IsImage(val string) bool {
	pattern := regexp.MustCompile(`^(url|image|image-set|cross-fade|element|(repeating-)?(linear|radial|conic)-gradient)\(.+\)$`)
	return pattern.MatchString(val)
}

func IsFraction(val string) bool {
	pattern := regexp.MustCompile(`^\d+\/\d+$`)
	return pattern.MatchString(val)
}

func IsNumber(val string) bool {
	return IsInteger(val) || IsFloat(val)
}

func IsInteger(val string) bool {
	_, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return true
}

func IsFloat(val string) bool {
	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false
	}
	return true
}

func IsLengthOnly(val string) bool {
	return lengthUnitRegex.MatchString(val) && !colorFnRegex.MatchString(val)
}

func GetIsArbitraryValue(val string, label interface{}, testValue func(string) bool) bool {
	res := arbitraryRegex.FindStringSubmatch(val)

	if len(res) > 1 {
		if res[1] != "" {

			if t, ok := label.(string); ok {
				return res[1] == t
			}

			if t, ok := label.(map[string]bool); ok {
				return t[res[1]]
			}
		}

		if len(res) > 2 {
			return testValue(res[2])
		}

	}

	return false
}
