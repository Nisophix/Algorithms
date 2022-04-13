package main

import (
	"fmt"
)

func main() {
	dec := "the quick brown fox JUMPS OVER THE LAZY DOG 1234567890!@#$%^&*()_+=-"
	enc := caesarEnc(dec, 6)
	fmt.Printf("original: %v\ncaesar: %v\noriginal again: %v\n", dec, enc, caesarDec(enc, 6))
}

func caesarEnc(text string, key rune) string {
	runeText := []rune(text)
	if key > 26 {
		key = key % 26
	}
	for i, ch := range runeText {
		if ch >= 'a' && ch <= 'z'-key || ch >= 'A' && ch <= 'Z'-key {
			ch = ch + key
		} else if ch > 'z'-key && ch <= 'z' || ch > 'Z'-key && ch <= 'Z' {
			ch = ch + key - rune(26)
		}
		runeText[i] = ch
	}
	return string(runeText)
}

func caesarDec(text string, key rune) string {
	return caesarEnc(text, 26 - key)
	
}