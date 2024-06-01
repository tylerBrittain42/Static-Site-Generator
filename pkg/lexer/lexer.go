package lexer

import "os"

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

type Token struct {
	category   tokenType
	value      string
	lineNumber int
}

type locationInfo struct {
	column int
	row    int
	len    int
}

func getFileContents(filePath string) ([]byte, error) {
	// opening(and closing) the file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// reading the file
	fStat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	b := make([]byte, fStat.Size())
	f.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil

}

/*
func Scan(b *[]byte) ([]Token, error) {
	location := locationInfo{0, 0, 0}
	for i := 0; i < len(*b); {
		location.scanToken()

	}
	return []Token{}, nil

}

func (*locationInfo) scanToken() (string, err) {
	return "foo", nil
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
	if startOfLine {
		var tokenType string
		headerCount = 0
		for i, v := range input {
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

	return token{}, nil
}
*/
