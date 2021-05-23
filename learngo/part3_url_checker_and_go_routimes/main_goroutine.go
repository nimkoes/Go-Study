package main

import (
	"fmt"
	"time"
)

func main() {

	// 길이 2 의 문자열 배열 생성
	people := [5]string{"nico", "nimkoes", "go", "java", "spring"}

	// bool 타입을 주고받을 수 있는 channel 생성, c 는 임의의 이름으로 사용 가능
	c := make(chan string)

	// 반복문을 실행 하면서 두 개의 goroutine 을 실행
	for _, person := range people {
		// channel 을 같이 전달
		go isSexy(person, c)
	}

	fmt.Println("waiting... ")

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	fmt.Println("DONE !")
}

func isSexy(person string, c chan string) {
	// 5초 동안 스레드를 멈춘다.
	time.Sleep(time.Second * 3)

	// channel 로 bool 값을 전달한다.
	c <- person + " is sexy"
}
