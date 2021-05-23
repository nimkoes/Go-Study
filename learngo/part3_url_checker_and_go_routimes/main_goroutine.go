package main

import (
	"fmt"
	"time"
)

func main() {

	// 길이 2 의 문자열 배열 생성
	people := [2]string{"nico", "nimkoes"}

	// bool 타입을 주고받을 수 있는 channel 생성, c 는 임의의 이름으로 사용 가능
	c := make(chan bool)

	// 반복문을 실행 하면서 두 개의 goroutine 을 실행
	for _, person := range people {
		// channel 을 같이 전달
		go isSexy(person, c)
	}

	// main function 은 channel 로부터 받는 값을 기다린다.
	// 기다린다는 것은 스레드를 종료하지 않는다는 것을 의미한다.
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	// 5초 동안 스레드를 멈춘다.
	time.Sleep(time.Second * 5)
	fmt.Println(person)

	// channel 로 bool 값을 전달한다.
	c <- true
}
