package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(TrimString("東東東東東東東東東東東"))

}

func TrimString(s string) string {
	runes := []rune(s)
	var res string

	if len(runes) <= 10 {
		return string(runes)
	}
	res = string(runes[0]) + strconv.Itoa(len(runes)-2) + string(runes[len(runes)-1])

	return res
}
