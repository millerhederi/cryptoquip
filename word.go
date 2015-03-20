package main

import (
	"fmt"
	"strconv"
)

type Word struct {
	Text string
	Signature []byte
	CompressedText []byte
}

func NewWord(text string) *Word {
	word := Word{ 
		Text: text,
		Signature: make([]byte, len(text)),
		CompressedText: make([]byte, 0),
	}

	byteValueMap := make(map[byte]byte)

	for i := 0; i < len(text); i++ {
		if text[i] < 'A' || text[i] > 'Z' {
			panic(fmt.Sprintf("%s, %v, %v", text, []byte(text), strconv.QuoteToASCII(text)))
		}

		value, ok := byteValueMap[text[i]]

		if !ok {
			value = byte(len(word.CompressedText))
			byteValueMap[text[i]] = value
			word.CompressedText = append(word.CompressedText, text[i])
		}

		word.Signature[i] = value
	}

	return &word
}

func (this *Word) AreEqualSignatures(other *Word) bool {
	if len((*this).Signature) != len((*other).Signature) {
		return false
	}

	for i := 0; i < len((*this).Signature); i++ {
		if (*this).Signature[i] != (*other).Signature[i] {
			return false
		}
	}

	return true
}

func (this *Word) String() string {
	return this.Text
}