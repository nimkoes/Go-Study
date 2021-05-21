package main

import (
	"fmt"
	"myDict"
)

func main() {
	dictionary := myDict.Dictionary{}

	baseWord := "hello"

	// 값을 추가
	dictionary.Add(baseWord, "First")

	// 저장된 값을 수정
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}

	// 저장된 값을 조회 한 다음 출력
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	fmt.Println()

	// 삭제 테스트 코드 작성
	fmt.Println("========== 삭제 테스트 시작 ==========")
	dictionary.Delete(baseWord)
	word, err = dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
