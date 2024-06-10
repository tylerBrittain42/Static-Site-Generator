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

// reset for next token. returns int of next index and error
func (b *byteSlice) SetNext() (int, error) {
	nextIndex := b.startingIndex + b.length
	if nextIndex >= len(b.s) {
		return 0, errors.New("Out of bounds")
	}
	b.startingIndex = b.startingIndex + b.length
	b.length = 0
	return b.startingIndex, nil
}
