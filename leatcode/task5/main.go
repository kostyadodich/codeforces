package main

import "fmt"

func main() {
	fmt.Println(LengthOfLastWord("Hello World "))
}

func LengthOfLastWord(s string) int {
	var counter int

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			counter++
		}
		if i != len(s)-1 {
			if s[i] == ' ' && s[i+1] != ' ' {
				return counter
			}
		}
	}

	return counter
}
