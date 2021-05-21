package main

import (
	"fmt"
	"myDict"
)

func main() {
	dictionary := myDict.Dictionary{}

	word := "hello"
	definition := "Greeting"

	// 첫번째 값 추가
	err := dictionary.Add(word, definition)

	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	// 같은 값을 한번 더 추가해서 결과 확인
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
