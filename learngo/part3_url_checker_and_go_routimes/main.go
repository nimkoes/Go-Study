package main

import (
	"errors"
	"fmt"
	"net/http"
)

// channel 을 통해 주고 받을 데이터 타입으로 사용 할 struct 선언
type result struct {
	url    string
	status string
}

// 사용자 정의 error
var errRequestFailed = errors.New("Request failed")

func main() {

	// url 접속 결과를 담을 비어있는 map 선언
	results := make(map[string]string)

	// channel 생성
	c := make(chan result)

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
		go hitURL(url, c)
	}

	// 실행 결과 출력
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan<- result) {

	// 현재 request 시도 하는 url 출력
	fmt.Println("Checking:", url)

	// Go reference 참고하여 url 에 Get 요청
	resp, err := http.Get(url)

	// result struct 의 status 값으로 사용 할 변수 선언
	status := "Ok"

	// err 가 있거나 http 응답 코드가 400 과 같거나 큰 경우 예외 처리
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- result{url: url, status: status}
}
