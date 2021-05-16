package main

import "fmt"

func main() {
	a := 2
	b := &a

	fmt.Println(a, b)
	fmt.Println(*b)

	// 실행 결과 예상
	myValue := 10
	anotherValue := &myValue
	*anotherValue = 21
	myValue = 17

	fmt.Println(myValue, *anotherValue)
	fmt.Println((&myValue == anotherValue), (myValue == *anotherValue))
}
