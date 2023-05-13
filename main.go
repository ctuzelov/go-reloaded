package main

import (
	"fmt"
	"os"
	"regexp"
	functions "reloaded/functions"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments")
		return
	}
	sampleName := os.Args[1]
	resultName := os.Args[2]
	if resultName == "sample.txt" || resultName == "go.mod" || resultName == "README.md" || resultName == "main.go" {
		fmt.Println("Write correct filename")
		return
	}

	fileData, err := os.ReadFile(sampleName)
	functions.CheckError(err)
	fileRune := []rune(string(fileData))
	resultStr := Formating(functions.Filtrate(fileRune))
	resultStr = strings.Join(functions.FixArticles(strings.Fields(resultStr)), " ")
	err = os.WriteFile(resultName, []byte(resultStr), 0644)
	functions.CheckError(err)
}

func Formating(strSlice []string) string {
	hexPattern := `\(hex\)`
	binPattern := `\(bin\)`
	upPattern := `\(up\)|\(up,$`
	capPattern := `\(cap\)|\(cap,$`
	lowPattern := `\(low\)|\(low,$`
	numPattern := `[^\(+]\)`
	for i := 0; i < len(strSlice); i++ {
		var numMatched bool
		word := strSlice[i]
		numberint := 1
		hexMatched, _ := regexp.MatchString(hexPattern, strSlice[i])
		binMatched, _ := regexp.MatchString(binPattern, strSlice[i])
		if i+1 < len(strSlice) {
			numMatched, _ = regexp.MatchString(numPattern, strSlice[i+1])
			if len(strSlice[i+1]) > 0 {
				if strSlice[i+1][0] == '(' {
					numMatched = false
				}
			}
		}
		upMatched, _ := regexp.MatchString(upPattern, strSlice[i])
		capMatched, _ := regexp.MatchString(capPattern, strSlice[i])
		lowMatched, _ := regexp.MatchString(lowPattern, strSlice[i])
		switch true {
		case hexMatched:
			strSlice = Converting(strSlice, 16, i)
		case binMatched:
			strSlice = Converting(strSlice, 2, i)
		case upMatched:
			if numMatched {
				numberint = ReturnNumb(strSlice, i+1)
				if numberint < 0 {
					continue
				}
				strSlice, i = CorrectWord(strSlice, numberint, i, strings.ToUpper, numMatched)
				continue
			}
			if word[len(word)-1] == ',' {
				continue
			}
			strSlice, i = CorrectWord(strSlice, numberint, i, strings.ToUpper, numMatched)
		case capMatched:
			if numMatched {
				numberint = ReturnNumb(strSlice, i+1)
				if numberint < 0 {
					continue
				}
				strSlice, i = CorrectWord(strSlice, numberint, i, strings.Title, numMatched)
				continue
			}
			if word[len(word)-1] == ',' {
				continue
			}
			strSlice, i = CorrectWord(strSlice, numberint, i, strings.Title, numMatched)
		case lowMatched:
			if numMatched {
				numberint = ReturnNumb(strSlice, i+1)
				if numberint < 0 {
					continue
				}
				strSlice, i = CorrectWord(strSlice, numberint, i, strings.ToLower, numMatched)
				continue
			}
			if word[len(word)-1] == ',' {
				continue
			}
			strSlice, i = CorrectWord(strSlice, numberint, i, strings.ToLower, numMatched)
		}
	}
	strSlice = functions.Filtrate([]rune(strings.Join(strSlice, " ")))
	return strings.Join(strSlice, " ")
}

func Converting(strSlice []string, base int, i int) []string {
	strSlice = RemoveIndex(strSlice, i)
	if i-1 < 0 {
		return strSlice
	}
	a := strSlice[i-1]
	var start int
	var end int = len(a)
	var started int
	var convInt int64
	var err error
	for j := range a {
		if unicode.IsNumber(rune(a[j])) || unicode.IsLetter(rune(a[j])) && base != 2 {
			if started == 0 {
				start = j
				started++
			}
		} else {
			if started == 1 {
				end = j
			}
		}
	}
	if start == 0 && end == len(a) {
		convInt, err = strconv.ParseInt(a, base, 64)
		strSlice[i-1] = strconv.Itoa(int(convInt))
	} else {
		convInt, err = strconv.ParseInt(a[start:end], base, 64)
		strSlice[i-1] = a[:start] + strconv.Itoa(int(convInt)) + a[end:]
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return strSlice
}

func CorrectWord(strSlice []string, numberint int, i int, f func(string) string, flag bool) ([]string, int) {
	if flag {
		strSlice[i] = strSlice[i] + " " + strconv.Itoa(numberint) + ")"
		strSlice = RemoveIndex(RemoveIndex(strSlice, i+1), i)
	} else {
		strSlice = RemoveIndex(strSlice, i)
	}
	for j := numberint; j > 0; j-- {
		if i-j >= 0 {
			strSlice[i-j] = strings.ToLower(strSlice[i-j])
			strSlice[i-j] = f(strSlice[i-j])
		} else {
			fmt.Println("Not enough words")
			break
		}
	}
	return strSlice, i - 1
}

func ReturnNumb(strSlice []string, i int) int {
	numberstr := ""
	for _, x := range strSlice[i] {
		if !unicode.IsDigit(x) {
			break
		}
		numberstr += string(x)
	}
	if numberstr == "" {
		return -1
	}
	numberint, _ := strconv.Atoi(numberstr)
	return numberint
}

func RemoveIndex(strSlice []string, index int) []string {
	if strSlice[index][len(strSlice[index])-1] != ')' {
		parIndex := 0
		for i := range strSlice[index] {
			if strSlice[index][i] == ')' {
				parIndex = i
				break
			}
		}
		strSlice[index] = strSlice[index][parIndex+1:]
		return strSlice
	}
	return append(strSlice[:index], strSlice[index+1:]...)
}

// fixed
