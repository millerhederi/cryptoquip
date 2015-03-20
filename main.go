package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	"github.com/davecheney/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	dict, err := NewDictionaryFromFile("en-US.dic")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open("cryptoquips.txt")
	if err != nil {
		panic("unable to open cryptoquips file")
	}
	defer file.Close()

	phrases := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrases = append(phrases, strings.ToUpper(scanner.Text()))
	}

	for _, phrase := range phrases {
		fmt.Println("\n\n\n--------------------------------------------------------------------------------")
		fmt.Println(phrase)
		fmt.Println("--------------------------------------------------------------------------------")

		Solve(dict, phrase)
	}
}