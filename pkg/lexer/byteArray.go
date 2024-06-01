package lexer

import "fmt"

type byteSlice struct {
	s         []byte
	index     int
	peakIndex int
	start     int
	length    int
}

/*
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
*/
func (b *byteSlice) getNextLexme() ([]byte, error) {
	firstChar := b.pop(0)
	lexme := []byte{}

	// Header cases
	// TODO: consider only a single exit poin
	if firstChar == '#'{

		
		for b.head() == '#'{
			b.pop()
			lexme = append(lexme, '#')
		// todo: make sure to verify the space
		return lexme
	// unordered list and line break
	} else if firstChar == '-'{
		if b.

func (b *byteSlice) pop(i int) byte {
	popped := b.s[i]
	b.s = append(b.s[:i], b.s[:i+1]...)
	return popped

}

// Index
func (b *byteSlice) getIndex() byte {
	return b.s[b.index]
}

func (b *byteSlice) incrementIndex() {
	b.index++
}

func (b *byteSlice) decrementIndex() {
	b.index--
}

func (b *byteSlice) peak() byte{
	return b.s[1]
}
func (b *byteSlice) head() byte{
	return b.s[0]
}
