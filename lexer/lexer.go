package lexer

import (
	"errors"
	"strconv"

	"github.com/thoas/go-funk"
)

//Lex - Parse json and return tokens
//Find patterns of strings, numbers, booleans, nulls, or JSON syntax like left brackets and left braces,
func Lex(input string) []interface{} {
	tokens := lexRecursive(input, 0, "")
	return tokens
}

var openChars = []string{"{", "\"", "["}
var closeChars = []string{"}", "\"", "]"}
var ignoreChars = []string{":", ","}

func lexRecursive(input string, position int, tempValue string) []interface{} {
	result := make([]interface{}, 0)
	if len(input) <= 0 || position > len(input)-1 {
		return result
	}

	//more aobut code points on https://blog.golang.org/strings
	codePoint := string(input[position])

	if funk.Contains(ignoreChars, codePoint) {
		//move on
	} else if funk.Contains(openChars, codePoint) || funk.Contains(closeChars, codePoint) {
		if len(tempValue) > 0 {
			result = append(result, parseValue(tempValue)...)
			tempValue = ""
		}
		if codePoint != "\"" {
			result = append(result, codePoint)
		}
	} else {
		tempValue += codePoint
	}

	position++
	result = append(result, lexRecursive(input, position, tempValue)...)

	return result
}

func parseValue(input string) []interface{} {
	result := make([]interface{}, 0)

	isNull := isNull(input)
	//check if value is null
	if isNull {
		result = append(result, nil)
	} else {
		//check if value is integer
		intValue, err := parseInt(input)
		if err != nil {
			boolValue, err := parseBool(input)
			if err == nil {
				//is nill??
				result = append(result, boolValue)
			} else {
				//if value couldnt be parsed as any of specific parsers than it is a string
				result = append(result, input)
			}
		} else {
			result = append(result, intValue)
		}
	}

	return result
}

func parseInt(input string) (uint64, error) {
	parsed, err := strconv.ParseUint(input, 10, 64)
	return parsed, err
}

func parseBool(input string) (bool, error) {
	if input == "true" {
		return true, nil
	} else if input == "false" {
		return false, nil
	}

	return false, errors.New("Cannot parse as bool")
}

func isNull(input string) bool {
	if input == "null" {
		return true
	}
	return false
}
