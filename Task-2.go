package main
import (
	"fmt"
	"strings"
)

func isDelimiter(r rune) bool {
	delimiters := " .,?!"
	return strings.ContainsRune(delimiters, r)
}

func getWordFrequency(s string) map[string]int {
	words := strings.FieldsFunc(s, isDelimiter)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[strings.ToLower(word)] += 1
	}

	return frequency
}

func isPalindrome(s string) bool{
	halfLen := int(len(s) / 2)
	for i := 0; i <= halfLen; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(getWordFrequency("Hello, hello, world World WORLD "))
	fmt.Println(isPalindrome("woow"))
	fmt.Println(isPalindrome("oow"))
	fmt.Println(isPalindrome("wow"))
	fmt.Println(isPalindrome("bow"))
	fmt.Println(isPalindrome("slow"))
}