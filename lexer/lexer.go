package lexer

import (
	"errors"
	"strconv"

	"github.com/thoas/go-funk"
)

var jsonSyntax = []string{":", ",", "{", "[", "}", "]"}
var ignoreChars = []string{"\t", "\b", "\n", "\r", " "}
var jsonQuote = "\""

//Lex - Parse json and return tokens
//Find patterns of strings, numbers, booleans, nulls, or JSON syntax like left brackets and left braces,
func Lex(input string) ([]interface{}, error) {
	result := make([]interface{}, 0)
	if len(input) <= 0 {
		return result, errors.New("Empty input")
	}

	tempValue := ""
	position := 0
	for position < len(input) {
		//more aobut code points on https://blog.golang.org/strings
		codePoint := string(input[position])

		//its a string, read until end of string
		if codePoint == jsonQuote {
			position++
			//add everything betwen quotes
			tempValue, position = readUntilCharacterReached(input, position, func(c string) bool { return string(c) == jsonQuote })
			result = append(result, tempValue)

			//position points on quote character,increase counter to skip it
			position++
		} else if funk.Contains(jsonSyntax, codePoint) {
			//add json syntax character and increase counter
			result = append(result, codePoint)
			position++
		} else if funk.Contains(ignoreChars, codePoint) {
			position++
		} else {
			tempValue, position = readUntilCharacterReached(input, position, func(c string) bool { return funk.Contains(jsonSyntax, c) || funk.Contains(ignoreChars, c) })

			parsedValue, err := parseValue(tempValue)
			if err != nil {
				return nil, err
			}
			result = append(result, parsedValue)
			//position is currently on breaking character, do not increase it, rather let it to trough loop
		}
	}

	return result, nil
}

func readUntilCharacterReached(input string, position int, isBreakingChar characterReached) (value string, newPosition int) {
	tempValue := ""

	for !isBreakingChar(string(input[position])) {
		tempValue += string(input[position])
		position++
	}

	return tempValue, position
}

type characterReached func(char string) bool

func parseValue(input string) (interface{}, error) {
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
				return nil, errors.New("Cannot parse value: " + input)
			}
		} else {
			result = append(result, intValue)
		}
	}

	return result[0], nil
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
