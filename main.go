package main

import (
	"fmt"

	"json_parser/lexer"
)

func main() {
	fmt.Println(lexer.Lex("test"))
}
