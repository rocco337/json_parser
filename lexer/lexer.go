package lexer

//Lex - Parse json and return tokens
func Lex(input string) []string {
	tokens := make([]string, 0)
	tokens = append(tokens, input)
	return tokens
}
