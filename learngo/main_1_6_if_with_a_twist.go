package main

import "fmt"

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}

	return true
}

func canIDrinkVer2(age int) bool {
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkVer2(16))
}
