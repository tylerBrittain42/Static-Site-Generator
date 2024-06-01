package byteArray

import (
	"errors"
)

type byteSlice struct {
	s             []byte
	startingIndex int
	length        int
}

// Increment the length
func (b *byteSlice) Add() error {
	if b.length+1 < len(b.s) {
		b.length++
		return nil
	}
	return errors.New("length will be out of bounds")
}

// Return value at index startingIndex + length
func (b *byteSlice) Peak() (byte, error) {
	if b.length+1 < len(b.s) {
		return b.s[b.startingIndex+b.length+1], nil

	}
	return 0, errors.New("length+1 will be out of bounds")
}

// Return subslice specified by startingIndex + length
func (b *byteSlice) Get() ([]byte, error) {
	if b.length == 0 {
		return []byte{}, errors.New("attempted to return empty string")
	}
	return b.s[b.startingIndex : b.startingIndex+b.length], nil
}

// Return char of length
//func (b *byteSlice) head() byte {
// reset for next word. CONSIDER using index + length instead due to weirdness with size 1
//func (b *byteSlice) SetNext() byte {
//
// func (b *byteSlice) getNextLexme() ([]byte, error) {
// 	firstChar := b.pop(0)
// 	lexme := []byte{}
//
// 	// Header cases
// 	// TODO: consider only a single exit poin
// 	if firstChar == '#'{
//
//
// 		for b.head() == '#'{
// 			b.pop()
// 			lexme = append(lexme, '#')
// 		// todo: make sure to verify the space
// 		return lexme
// 	// unordered list and line break
// 	} else if firstChar == '-'{
// }
/*
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


   func (b *byteSlice) head() byte {
   	return b.s[0]
   }
*/
