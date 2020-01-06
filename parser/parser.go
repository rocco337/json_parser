package parser

import (
	"errors"
	"fmt"
	"reflect"
)

var arrayOpen = "["
var arrayClose = "]"
var objectOpen = "{"
var objectClose = "}"
var colon = ":"

//Parse - from lsit of tokens detect objects, arrays
func Parse(tokens []interface{}, returnObject interface{}) error {

	if tokens[0] == objectOpen {
		parseObject(tokens, 1, returnObject)
		return nil

	} else if tokens[0] == arrayOpen {
		parseArray(tokens, 1, returnObject)
		return nil
	}

	return errors.New("Cannot parse object")
}

func parseObject(tokens []interface{}, i int, returnObject interface{}) (index int) {
	for objectClose != tokens[i] {
		if tokens[i] == colon {
			fieldName := tokens[i-1]
			fieldValue := tokens[i+1]

			if fieldName == objectOpen {
				i++ //skip opening tag
				parsedObject := parseObject(tokens, i, returnObject)
				setValueByFieldName(returnObject, fmt.Sprintf("%v", fieldName), parsedObject)
			} else if fieldName == arrayOpen {
				i++ //skip opening tag
				parseArray(tokens, i, returnObject)
				//setValueByFieldName(returnObject, fmt.Sprintf("%v", fieldName), parsedArray)
			} else {
				setValueByFieldName(returnObject, fmt.Sprintf("%v", fieldName), fieldValue)
			}
		}

		i++
	}

	return i
}

func parseArray(tokens []interface{}, i int, returnObject interface{}) (index int) {
	result := make([]interface{}, 0)

	for arrayClose != tokens[i] {
		if tokens[i] == arrayOpen {
			i++ //skip opening tag

			//todo - fix me
			parsedArray := make([]int, 0)
			i = parseArray(tokens, i, &parsedArray)
			result = append(result, parsedArray)
		} else if tokens[i] == objectOpen {
			i++ //skip opening tag

			//get type of element in slice
			elemType := reflect.TypeOf(returnObject).Elem().Elem()

			//create pointer of object to fill
			parsedObject := reflect.New(elemType).Interface()
			i = parseObject(tokens, i, parsedObject)

			valueFromPointer := reflect.ValueOf(parsedObject).Elem().Interface()

			result = append(result, valueFromPointer)
		} else {
			result = append(result, tokens[i])
		}

		i++
	}

	setArrayValue(returnObject, result)
	return i
}

func setValueByFieldName(targetObject interface{}, targetFieldName string, value interface{}) {
	element := reflect.ValueOf(targetObject).Elem()
	fieldName := element.FieldByName(targetFieldName)

	valueReflected := reflect.ValueOf(value)
	fieldName.Set(valueReflected)
}

func setArrayValue(targetObject interface{}, values []interface{}) {
	object := reflect.ValueOf(targetObject)

	//if it is regular array, just use normal append
	if isArrayOrSlice(targetObject) {
		for _, val := range values {
			valueReflected := reflect.ValueOf(val)
			reflect.Append(object, valueReflected)
		}
	} else {
		object := object.Elem()

		for _, val := range values {
			valueReflected := reflect.ValueOf(val)
			object.Set(reflect.Append(object, valueReflected))
		}
	}

}

func isArrayOrSlice(object interface{}) bool {
	for t := reflect.TypeOf(object); ; {
		switch t.Kind() {
		case reflect.Array, reflect.Slice:
			return true
		default:
			return false
		}
	}
}
