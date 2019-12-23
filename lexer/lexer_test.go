package lexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLexer(t *testing.T) {

	t.Run("Lex - should return tokens", func(t *testing.T) {
		user := struct {
			Userid              int
			Roles               []string
			IsActive            bool
			ShouldResetPassword bool
			ParentUserID        string
		}{
			1,
			[]string{"Admin", "User", "Operator"},
			true,
			false,
			"null",
		}
		userJSON, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Cannot serialize test data")
		}
		tokens, _ := Lex(string(userJSON))

		expectedResult := []string{
			"{",
			"Userid", "1",
			"Roles", "[",
			"Admin", "User", "Operator",
			"]",
			"IsActive", "true",
			"ShouldResetPassword", "false",
			"ParentUserID", "null",
			"}",
		}

		if len(expectedResult) != len(tokens) {
			t.Errorf("Result has different length than expected result. Expected: %+q Got: %+q", expectedResult, tokens)
		}

		for i := 0; i < len(expectedResult); i++ {
			if expectedResult[i] == "null" && tokens[i] == nil {
				//its ok, move on
			} else if expectedResult[i] != fmt.Sprintf("%v", tokens[i]) {
				t.Errorf("Values differ: Expected: %s Got: %v", expectedResult[i], tokens[i])
			}
		}
	})

	t.Run("Lex - read glossary.json", func(t *testing.T) {
		jsonData, _ := ioutil.ReadFile("testdata/glossary.json")
		tokens, _ := Lex(string(jsonData))

		if tokens[1] != "value" {
			t.Errorf("Wrong value on line 1, expected: value, 154 got: %v, %v", tokens[1], tokens[2])
		}

		if tokens[3] != "glossary" {
			t.Errorf("Wrong value on line 3, expected: glossary got: %v", tokens[3])
		}

		if tokens[20] != "Standard Generalized Markup Language" {
			t.Errorf("Wrong value on line 3, expected: StandardGeneralizedMarkupLanguage Got: %v", tokens[20])
		}
	})

	// t.Run("Lex - invalid json", func(t *testing.T) {
	// 	tokens,_ := Lex("{ aa-234 }")
	// })
}
