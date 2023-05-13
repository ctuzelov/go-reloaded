package reloaded

import (
	"fmt"
	"unicode"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func Filtrate(fileRune []rune) []string {
	var word string
	var arr []string
	IsFirstSing := true
	IsFirstDoub := true
	for i := 0; i < len(fileRune); i++ {
		x := fileRune[i]
		switch true {
		case x == ' ' || x == '\n':
			AddWordToSlice(&word, &arr)
		case x == '(' || (x == '\'' && IsFirstSing) || (x == '"' && IsFirstDoub):
			if x == '\'' {
				RemoveSpacesAfter(&fileRune, i)
				AddWordToSlice(&word, &arr)
				if len(arr) != 0 && i+1 < len(fileRune) {
					last := arr[len(arr)-1]
					if unicode.ToLower(rune(last[len(last)-1])) == ('n') && unicode.ToLower(fileRune[i+1]) == 't' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] += "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					} else if len(fileRune) >= i+2 && unicode.ToLower(fileRune[i+1]) == 's' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] = last + "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					} else if len(fileRune) >= i+3 && string(unicode.ToLower(fileRune[i+1]))+string(unicode.ToLower(fileRune[i+2])) == "ll" {
						switch true {
						case i+3 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+3]):
							arr[len(arr)-1] = last + "'" + "ll"
							fileRune = append(fileRune[:i+1], fileRune[i+3:]...)
							continue
						}
					} else if len(fileRune) >= i+2 && unicode.ToLower(rune(last[len(last)-1])) == ('i') && unicode.ToLower(fileRune[i+1]) == 'm' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] += "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					}
				}
			}
			SwitchIsFirstQuote(x, &IsFirstSing, &IsFirstDoub)
			RemoveSpacesAfter(&fileRune, i)
			// in case if ("' comes together
			if word != "" && (word[len(word)-1] == byte('(') || word[len(word)-1] == byte('"') || word[len(word)-1] == byte('\'')) {
				word += string(x)
				continue
			}
			AddWordToSlice(&word, &arr)
			word += string(x)
		case x == ',' || x == ')' || x == '!' || x == '?' || x == '.' || x == ';' || x == ':' || (x == '\'' && !IsFirstSing) || (x == '"' && !IsFirstDoub):
			if x == '\'' {
				RemoveSpacesAfter(&fileRune, i)
				AddWordToSlice(&word, &arr)
				if len(arr) != 0 && i+1 < len(fileRune) {
					last := arr[len(arr)-1]
					if len(fileRune) >= i+2 && unicode.ToLower(rune(last[len(last)-1])) == ('n') && unicode.ToLower(fileRune[i+1]) == 't' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] += "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					} else if len(fileRune) >= i+2 && unicode.ToLower(fileRune[i+1]) == 's' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] = last + "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					} else if len(fileRune) >= i+3 && string(unicode.ToLower(fileRune[i+1]))+string(unicode.ToLower(fileRune[i+2])) == "ll" {
						switch true {
						case i+3 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+3]):
							arr[len(arr)-1] = last + "'" + "ll"
							fileRune = append(fileRune[:i+1], fileRune[i+3:]...)
							continue
						}
					} else if len(fileRune) >= i+2 && unicode.ToLower(rune(last[len(last)-1])) == ('i') && unicode.ToLower(fileRune[i+1]) == 'm' {
						switch true {
						case i+2 >= len(fileRune):
							fallthrough
						case !unicode.IsLetter(fileRune[i+2]):
							arr[len(arr)-1] += "'" + string(fileRune[i+1])
							fileRune = append(fileRune[:i+1], fileRune[i+2:]...)
							continue
						}
					}
				}
			}
			SwitchIsFirstQuote(x, &IsFirstSing, &IsFirstDoub)
			AddWordToSlice(&word, &arr)
			if len(arr) > 0 {
				arr[len(arr)-1] += string(x)
			} else {
				word += string(x)
			}
			if i == len(fileRune)-1 {
				return arr
			}
		default:
			word += string(x)
		}
	}
	arr = append(arr, word)
	return arr
}

func RemoveSpacesAfter(fileRune *[]rune, i int) {
	for j := i + 1; j != ' ' && j < len(*fileRune); {
		if (*fileRune)[j] == ' ' {
			*fileRune = append((*fileRune)[:j], (*fileRune)[j+1:]...)
		} else {
			break
		}
	}
}

func SwitchIsFirstQuote(x rune, IsFirstSing *bool, IsFirstDoub *bool) {
	switch x {
	case '\'':
		*IsFirstSing = !*IsFirstSing
	case '"':
		*IsFirstDoub = !*IsFirstDoub
	}
}

// Adds word to slice of strings and clears it
func AddWordToSlice(word *string, arr *[]string) {
	if len(*word) >= 1 {
		*arr = append(*arr, *word)
	}
	*word = ""
}
