package main

import "fmt"

func main() {

	fmt.Println(LargestOddNumber("1033390"))

}

func LargestOddNumber(num string) string {
	var res string
	cymbols := map[string]struct{}{
		"0": {},
		"2": {},
		"4": {},
		"6": {},
		"8": {},
	}

	for i := len(num) - 1; i >= 0; i-- {
		_, ok := cymbols[string(num[i])]
		if !ok {
			res += num[:i+1]
			return res
		}
	}

	return ""
}
