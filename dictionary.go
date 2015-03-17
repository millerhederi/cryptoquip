package main

import (
	"os"
	"strings"
	"bufio"
	"io"
)

type Dictionary struct {
	Words []*Word
}

func NewDictionaryFromFile(fname string) (*Dictionary, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return NewDictionary(file)
}

func NewDictionary(reader io.Reader) (*Dictionary, error) {
	dictionary := Dictionary{
		Words: make([]*Word, 0),
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		dictionary.Words = append(dictionary.Words, NewWord(strings.ToUpper(scanner.Text())))
	}

	return &dictionary, scanner.Err()
}

func (this *Dictionary) QueryBySignature(word *Word) []*Word {
	results := make([]*Word, 0)

	for _, w := range this.Words {
		if (*w).AreEqualSignatures(word) {
			results = append(results, w)
		}
	}

	return results
}