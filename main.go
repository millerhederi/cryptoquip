package main

import (
	"strings"
	"fmt"

	"github.com/davecheney/profile"
)

type CandidateMap struct {
	cyphertext *Word
	plaintextCandidates []*Word
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	
	dict, err := NewDictionaryFromFile("en-US.dic")
	if err != nil {
		fmt.Println(err)
		return
	}

	cryptogram := strings.Split("TJRBA AKFRXO KX EWKXOA R WAIAIYAWAS FJKF K WRXS RL K FAWWRYBA FJRXO FE FKLFA", " ")

	cryptogramWords := make([]*Word, len(cryptogram))
	for i, cyphertext := range cryptogram {
		cryptogramWords[i] = NewWord(cyphertext)
	}

	candidates := make([]CandidateMap, len(cryptogramWords))
	for i, word := range cryptogramWords {
		candidates[i] = CandidateMap{ word, dict.QueryBySignature(word) }
	}

	solve(candidates)
}

func solve(candidates []CandidateMap) {
	code := NewCode()

	recursiveSolve(candidates, code)
}

func recursiveSolve(candidates []CandidateMap, code *Code) {
	if len(candidates) == 0 {
		fmt.Println(code.decryptionMap)
		return
	}

	for _, candidate := range candidates[0].plaintextCandidates {
		if keys, ok := code.TryUpdateWithDecryptedWord((*candidates[0].cyphertext).CompressedText, (*candidate).CompressedText); ok {
			recursiveSolve(candidates[1:], code)

			code.RemoveKeys(keys)
		}
	}
}