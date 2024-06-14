package main

import (
	"os"
	"github.com/01-edu/z01"
)

// reverseWords reverses the words in the input string.
func reverseWords(s string) string {
	words := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			words = append(words, s[start:i])
			start = i + 1
		}
	}
	words = append(words, s[start:]) // Append the last word

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join words into a single string with spaces
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	reversed := reverseWords(input)

	for _, r := range reversed {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
