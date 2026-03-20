// // // // // // // // // // Here, we will write a function that checks whether any rune in a s is a vowel.
// // // // // // // // // // we will be using the strings.ContainsFunc
// // // // // // // // // // The strings.ContainsFunc checks if a character(a rune) in string satisfies a given condition

// // // // // // // // // package main

// // // // // // // // // import (
// // // // // // // // // 	"fmt"
// // // // // // // // // 	"strings"
// // // // // // // // // )

// // // // // // // // // func check4vowel(s string) bool {
// // // // // // // // // 	f := func(r rune) bool {
// // // // // // // // // 		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
// // // // // // // // // 	}
// // // // // // // // // 	return strings.ContainsFunc(s, f)
// // // // // // // // // }

// // // // // // // // // func main() {
// // // // // // // // // 	fmt.Println(check4vowel("hello"))
// // // // // // // // // 	fmt.Println(check4vowel("HELLO")) // it returned false because it is case sensitive.
// // // // // // // // // 	fmt.Println(check4vowel("Hello"))
// // // // // // // // // 	fmt.Println(check4vowel("success"))
// // // // // // // // // }

// // // // // // // // // Here we will write a function that checks if there is Punctuation in a s

// // // // // // // // package main

// // // // // // // // import(
// // // // // // // // 	"fmt"
// // // // // // // // 	"strings"
// // // // // // // // )

// // // // // // // // func CheckPunct(words string) bool  {
// // // // // // // // 	f := func (r rune) bool {
// // // // // // // // 		return r =='!' || r == '?' || r == '.'
// // // // // // // // 	}
// // // // // // // // 	return strings.ContainsFunc(words, f)
// // // // // // // // }

// // // // // // // // func main()  {
// // // // // // // // 	fmt.Println(CheckPunct("hello!"))
// // // // // // // // 	fmt.Println(CheckPunct("How are you?"))
// // // // // // // // 	fmt.Println(CheckPunct("my name"))
// // // // // // // // }

// // // // // // // // Here we will check if a the first character in a s is a vowel.
// // // // // // // // If it is a vowel, it will  return "an" but if it is not, it will return "a"
// // // // // // // // we will be using the strings.ContainsRune

// // // // // // // package main

// // // // // // // import (
// // // // // // // 	"fmt"
// // // // // // // 	"strings"
// // // // // // // )

// // // // // // // func Is1stVowel(s string) string {
// // // // // // // 	vowels := "aeiouh"
// // // // // // // 	if strings.ContainsRune(vowels, rune(s[0])) {
// // // // // // // 		return "a"
// // // // // // // 	}
// // // // // // // 	return "an"
// // // // // // // }

// // // // // // // func main() {
// // // // // // // 	fmt.Println(Is1stVowel("actor"))
// // // // // // // 	fmt.Println(Is1stVowel("book"))
// // // // // // // 	fmt.Println(Is1stVowel("apple"))
// // // // // // // }

// // // // // // // Here we will be extracting all words into a slice (ignore extra spaces).
// // // // // // // we will be using the strings.Fields

// // // // // package main

// // // // // import(
// // // // // 	"fmt"
// // // // // 	"strings"
// // // // // )

// // // // // func IntoA_Slice(s string) []string  {
// // // // // 	word := strings.Fields(s)
// // // // // 	return word
// // // // // }

// // // // // func main()  {
// // // // // 	fmt.Println(IntoA_Slice(" Go  is   awesome "))
// // // // // }

// // // // // // // here we will be counting the number of words in a text

// // // // package main

// // // // import("fmt"; "strings")

// // // // func ChecknumWord(text string) int  {
// // // // 	  word := strings.Fields(text)
// // // // 		return len(word)
// // // // }

// // // // func main()  {
// // // // 	fmt.Println(ChecknumWord("Go is Awesome"))
// // // // }

// // // // here we want to print each word on a new line

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strings"
// // // )

// // // func PrintEachWord(s string) []string {
// // //       words := strings.Fields(s)
	  
// // // 	  for _, r := range words{
// // // 		fmt.Println(r)
// // // 	  }
// // // 	  return words
	  
// // // 	}

// // // func main() {
// // // 	fmt.Println(PrintEachWord("red, yellow, blue"))
// // // }

// // Here we are spliting the words using the strings.Split package

// // package main

// // import("fmt"; "strings")

// // func SplitColors(s string) []string {
// // 	words := strings.Split(s, ",")
// // 	return words
// // }

// // func main()  {
// // 	fmt.Println(SplitColors("yellow,red,blue,black"))
// // }



// package main

// import("fmt"; "strings")

// func SplitWords(word string) []string {
// 	text := strings.Split(word, "-")
// 	return text
// }

// func main()  {
// 	fmt.Print(SplitWords("1999-03-23"))
// }