package lexer

import (
	"encoding/json"
	"fmt"
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
		tokens := Lex(string(userJSON))

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

}
