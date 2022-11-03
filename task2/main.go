package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	input := make([]string, 0, n)
	for range make([]struct{}, n) {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	for _, s := range input {
		fmt.Println(TrimString(s))
	}

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
