package twmerge

type SplitModifiersFn = func(string) (baseClass string, modifiers []string, hasImportant bool, maybePostfixModPosition int)

func MakeSplitModifiers(conf *TwMergeConfig) SplitModifiersFn {
	separator := conf.ModifierSeparator

	return func(className string) (string, []string, bool, int) {
		modifiers := []string{}
		modifierStart := 0
		bracketDepth := 0
		// used for bg-red-500/50 (50% opacity)
		maybePostfixModPosition := -1

		for i := 0; i < len(className); i++ {
			char := rune(className[i])

			if char == '[' {
				bracketDepth++
				continue
			}
			if char == ']' {
				bracketDepth--
				continue
			}

			if bracketDepth == 0 {
				if char == separator {
					modifiers = append(modifiers, className[modifierStart:i])
					modifierStart = i + 1
					continue
				}

				if char == conf.PostfixModifier {
					maybePostfixModPosition = i
				}
			}
		}

		baseClassWithImportant := className[modifierStart:]
		hasImportant := baseClassWithImportant[0] == byte(conf.ImportantModifier)
		var baseClass string
		if hasImportant {
			baseClass = baseClassWithImportant[1:]
		} else {
			baseClass = baseClassWithImportant
		}

		// fix case where there is modifier & maybePostfix which causes maybePostfix to be beyond size of baseClass!
		if maybePostfixModPosition != -1 && maybePostfixModPosition > modifierStart {
			maybePostfixModPosition -= modifierStart
		}

		return baseClass, modifiers, hasImportant, maybePostfixModPosition

	}
}
