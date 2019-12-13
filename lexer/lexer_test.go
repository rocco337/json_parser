package lexer

import "testing"

func TestLexer(t *testing.T) {

	t.Run("Lex should return tokens", func(t *testing.T) {
		tokens := Lex("{ \"userId\": 1}")
		if tokens[0] != "{" {
			t.Fail()
		}
	})

}
