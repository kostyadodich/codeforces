package main

import "fmt"

func main() {
	fmt.Println(Remainder(8))
	fmt.Println(Remainder(5))

}

func Remainder(n int) string {
	res := ""
	if n%2 == 0 {
		res = "YES"
		return res
	} else {
		res = "NO"
	}

	return res
}
