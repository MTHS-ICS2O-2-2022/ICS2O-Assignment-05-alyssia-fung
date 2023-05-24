package main

import "fmt"

func reverseWord(word string) string {
	reversed := ""
	// The use of len statment
	// https://www.includehelp.com/golang/len-function-with-examples.aspx
	for ver := len(word) - 1; ver >= 0; ver-- {
		reversed += string(word[ver])
	}
	return reversed
}

func main() {
	// Replace "Phone" with the word you want to reverse
	word := "tacocat"
	reversedWord := reverseWord(word)
	fmt.Println(reversedWord) // Output: edoc
	fmt.Println("\nDone.")
}
