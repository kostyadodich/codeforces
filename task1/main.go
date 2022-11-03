package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)

	fmt.Println(Remainder(n))

}

func Remainder(n int) string {
	var res string

	if n == 2 {
		res = "NO"
	} else if n%2 == 0 {
		res = "YES"
	} else {
		res = "NO"
	}

	return res
}
