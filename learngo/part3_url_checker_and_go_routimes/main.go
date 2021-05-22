package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 사용자 정의 error
var errRequestFailed = errors.New("Request failed")

func main() {

	// url 접속 결과를 담을 비어있는 map 선언
	// 방법 1
	// results := map[string]string{}

	// 방법 2
	// var results = map[string]string{}

	// 방법 3 :: make 는 map 을 만들어주는 함수
	var results = make(map[string]string)

	// 이 방법은 panic: assignment to entry in nil map 발생
	// var results map[string]string

	// 접속을 시도 할 url 목록
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		"https://xxxelppa.tistory.com/",
		"https://nimkoes.github.io/",
	}

	// 반복문을 사용하여 각 url 에 접속 시도
	for _, url := range urls {
		result := "OK"

		err := hitURL(url)

		// function 실행 결과 err 가 있으면 메시지 출력
		if err != nil {
			// err 가 nil 이 아니면 result 를 FAILED 처리
			result = "FAILED"
		}

		// url 에 대한 GET request 결과를 map 에 저장
		results[url] = result
	}

	fmt.Println()

	// 실행 결과 출력
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {

	// 현재 request 시도 하는 url 출력
	fmt.Println("Checking:", url)

	// Go reference 참고하여 url 에 Get 요청
	resp, err := http.Get(url)

	// err 가 있거나 http 응답 코드가 400 과 같거나 큰 경우 예외 처리
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
