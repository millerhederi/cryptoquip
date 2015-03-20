package main

import (
	"fmt"
	"strings"
	"sort"
)

type SingleWordSolutionSpace struct {
	ciphertext *Word
	plaintextCandidates []*Word
}

type SolutionSpace []*SingleWordSolutionSpace

func (this SolutionSpace) Len() int {
	return len(this)
}

func (this SolutionSpace) Less(i, j int) bool {
	return len(this[i].plaintextCandidates) < len(this[j].plaintextCandidates)
}

func (this SolutionSpace) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func clean(word string) (string, bool) {
	word = strings.ToUpper(word)

	cleanMap := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r
		case r == '.' || r == ';' || r == ',' || r == '"' || r == '?' || r == ':' || r == '!':
			return -1
		}
		return 0
	}

	result := strings.Map(cleanMap, word)

	if strings.ContainsRune(result, 0) {
		return "", false
	}

	return result, true
}

func Solve(dictionary *Dictionary, ciphertext string) {
	ciphertextStrings := strings.Split(ciphertext, " ")

	solutionSpace := make(SolutionSpace, 0, len(ciphertextStrings))
	usedWords := make(map[string]bool, len(ciphertextStrings))

	for _, c := range ciphertextStrings {
		if word, ok := clean(c); ok {
			if !usedWords[word] {
				usedWords[word] = true
				solutionSpace = append(solutionSpace, &SingleWordSolutionSpace{ ciphertext: NewWord(word) })
			}
		}
	}

	for _, s := range solutionSpace {
		(*s).plaintextCandidates = dictionary.QueryBySignature((*s).ciphertext)
	}

	sort.Sort(solutionSpace)

	code := NewCode()
	recursiveSolve(solutionSpace, code, ciphertext)
}

func recursiveSolve(candidates SolutionSpace, code *Code, ciphertext string) {
	if len(candidates) == 0 {
		fmt.Printf("%s", code.Decrypt([]byte(ciphertext)))
		fmt.Println("    ", code.decryptionMap)
		return
	}

	for _, candidate := range (*candidates[0]).plaintextCandidates {
		if keys, ok := code.TryUpdateWithDecryptedWord((*(*candidates[0]).ciphertext).CompressedText, (*candidate).CompressedText); ok {
			recursiveSolve(candidates[1:], code, ciphertext)

			code.RemoveLastNKeys(keys)
		}
	}
}
