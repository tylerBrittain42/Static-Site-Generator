package lexer

import (
	"fmt"
)

type tokenType int

const (
	h1 tokenType = iota
	h2
	h3
	h4
	h5
	h6
	ol
	ul
	p
	br
	literal
	bold
	newLine
)

type token struct {
	category   tokenType
	value      string
	lineNumber int
}

// {"# this is a **heading**\n", token{h1, "#", 0}},
func getNextToken(input string) (token, error) {
	fmt.Println(input)
	return token{}, nil
}

// input is a complete lexme
// using startOfLine to seperate block tokens from inline tokens
// TODO: how to handle paragraphs
func getToken(input string, startOfLine bool) (token, error) {
	// TODO: grab until first space
	if startOfLine{
		var tokenType string
		headerCount = 0
		for i, v := range(input) {
			if v == "#" {
			} 
		}

	} else {
		// string literal
		// bold
		// italic
		// TODO: link
		// TODO: image

	}



	
	switch token()
	return token{}, nil
}
