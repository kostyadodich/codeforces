package main

import "fmt"

func main() {
	remainder(8)
	remainder(5)
}

func remainder(n int) {
	if n%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
