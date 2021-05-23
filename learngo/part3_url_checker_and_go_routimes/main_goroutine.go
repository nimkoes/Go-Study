package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	sexyCount("nimkoes")
}

func sexyCount(person string) {
	for i := 0; i < 5; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
