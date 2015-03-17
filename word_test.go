package main

import (
	"testing"
	"reflect"
)

func TestNewWord(t *testing.T) {
	cases := []struct {
		input string
		result Word
	} {
		{ "", Word{ Signature: []byte{}, CompressedText: []byte{}}},
		{ "cat", Word{ Signature: []byte{ 0, 1, 2 }, CompressedText: []byte("cat")}},
		{ "cac", Word{ Signature: []byte{ 0, 1, 0 }, CompressedText: []byte("ca")}},
		{ "cca", Word{ Signature: []byte{ 0, 0, 1 }, CompressedText: []byte("ca")}},
		{ "abcdabacc", Word{ Signature: []byte{ 0, 1, 2, 3, 0, 1, 0, 2, 2 }, CompressedText: []byte("abcd")}},
	}

	for _, c := range cases {
		result := NewWord(c.input)

		if !reflect.DeepEqual(c.result.Signature, result.Signature) {
			t.Errorf("NewWord(%s).Signature = %v, expected %v", c.input, result.Signature, c.result.Signature)
		}

		if !reflect.DeepEqual(c.result.CompressedText, result.CompressedText) {
			t.Errorf("NewWord(%s).CompressedText = %v, expected %v", c.input, result.CompressedText, c.result.CompressedText)
		}
	}
}

func TestAreEqualSignatures(t *testing.T) {
	cases := []struct {
		a *Word
		b *Word
		result bool
	} {
		{ NewWord("cat"), NewWord("dog"), true },
	}

	for _, c := range cases {
		result := c.a.AreEqualSignatures(c.b)

		if result != c.result {
			t.Errorf("AreEqualSignatures(%s, %s) == %v, expected %v", c.a, c.b, result, c.result)
		}
	}
}