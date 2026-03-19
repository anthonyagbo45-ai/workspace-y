// // Here, we will write a function that checks whether any rune in a word is a vowel.
// // we will be using the strings.ContainsFunc
// // The strings.ContainsFunc checks if a character(a rune) in string satisfies a given condition

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func check4vowel(word string) bool {
// 	f := func(r rune) bool {
// 		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
// 	}
// 	return strings.ContainsFunc(word, f)
// }

// func main() {
// 	fmt.Println(check4vowel("hello"))
// 	fmt.Println(check4vowel("HELLO")) // it returned false because it is case sensitive.
// 	fmt.Println(check4vowel("Hello"))
// 	fmt.Println(check4vowel("success"))
// }


// Here we will write a function that checks if there is Punctuation in a word

package main

import(
	"fmt"
	"strings"
)

func CheckPunct(words string) bool  {
	f := func (r rune) bool {
		return r =='!' || r == '?' || r == '.' 
	}
	return strings.ContainsFunc(words, f)
}

func main()  {
	fmt.Println(CheckPunct("hello!"))
	fmt.Println(CheckPunct("How are you?"))
	fmt.Println(CheckPunct("my name"))
}