package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	if len(word) == 0 {
		return false
	}
	word = strings.ToLower(word)
	for i, c := range word {
		last := len(word) - 1 - i
		if string(c) != string(word[last]) {
			return false
		}
		if i == last {
			break
		}
	}

	return true
}

func main() {
	fmt.Println("Digite uma palavra: ")
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("The word '%s' is a Palidrome ? %t ", word, isPalindrome(word))

}
