package main

import "fmt"

func main() {
	fmt.Println(Domino(11, 1))

}

func Domino(width, height int) int {
	s := width * height

	return s / 2
}
