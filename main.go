package main

import (
	"coding-katas/minsteps"
	"fmt"
)

func main() {
	input := 1000000

	result := minsteps.CountTab(input)
	fmt.Println(result)
}
