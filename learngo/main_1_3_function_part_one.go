package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nimkoes")
	fmt.Println(totalLength, upperName)

	myTotalLength, _ := lenAndUpper("xxxelppa")
	fmt.Println(myTotalLength)

	repeatMe("nimkoes", "test", "nico", "go", "java")
}
