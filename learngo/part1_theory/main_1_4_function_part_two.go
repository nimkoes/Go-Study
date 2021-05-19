package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	defer fmt.Println("second message ?")

	fmt.Println("start !")
	length = len(name)
	defer fmt.Println("third ?")
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upper := lenAndUpper("nimkoes")
	fmt.Println(totalLength, upper)
}
