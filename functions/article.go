package reloaded

import (
	"unicode"
)

func FixArticles(s []string) []string {
	for i := range s {
		if s[i] == "a" || s[i] == "A" {
			switch true {
			case i+1 >= len(s):
				break
			case isVowel(rune(s[i+1][0])):
				if s[i] == "A" {
					s[i] = "An"
					continue
				}
				s[i] = "an"
			}
		}
		if s[i] == "an" || s[i] == "An" || s[i] == "AN" || s[i] == "aN" {
			switch true {
			case i+1 >= len(s):
				break
			case !isVowel(rune(s[i+1][0])):
				if string(s[i][:2]) == "An" {
					s[i] = "A"
					continue
				}
				s[i] = "a"
			}
		}
	}
	return s
}

func isVowel(character rune) bool {
	character = unicode.ToLower(character)
	return character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u'
}
