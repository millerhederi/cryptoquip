package main

import (
	"testing"
	"strings"
	"reflect"
)

func TestNewDictionary(t *testing.T) {
	texts := []string{ "cat", "and", "a" }
	dict, err := NewDictionary(strings.NewReader(strings.Join(texts, "\n")))

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(dict.Words) != len(texts) {
		t.Errorf("len(dictionary) == %d, expected %d", len(dict.Words), len(texts))
	}

	for i, word := range dict.Words {
		expectedText := strings.ToUpper(texts[i])
		if word.Text != expectedText {
			t.Errorf("dict[%d] == %s, expected %s", i, word.Text, expectedText)
		}
	}
}

func _TestQueryBySignature(t *testing.T) {
	texts := []string{ "CAT", "CAC", "A" }
	dict, err := NewDictionary(strings.NewReader(strings.Join(texts, "\n")))

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	cases := []struct {
		input string
		result []int
	} {
		{ "XFE", []int{ 0 }},
		{ "XFX", []int{ 1 }},
		{ "JJ", []int{ }},
		{ "R", []int{ 2 }},
	}

	for _, c := range cases {
		result := dict.QueryBySignature(NewWord(c.input))

		expected := make([]*Word, len(c.result))
		for _, index := range c.result {
			expected = append(expected, dict.Words[index])
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QueryBySignature(%s) %v", c.input, len(result))
		}
	}
}