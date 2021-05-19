package main

import "fmt"

func main() {
	sample_map := map[string]string{"key": "value", "name": "nimkoes"}

	sample_map["add_key"] = "add_value" // 새로운 값을 추가
	sample_map["key"] = "modify_value"  // 기존의 값을 수정 (덮어쓰기)

	// range 를 사용 한 반복
	for key, value := range sample_map {
		fmt.Println(key, value)
	}
	fmt.Println()

	// value 만 사용하고 싶을 때
	for _, v := range sample_map {
		fmt.Println(v)
	}
	fmt.Println()

	// key 만 사용하고 싶을 때
	for k, _ := range sample_map {
		fmt.Println(k)
	}
	fmt.Println()

}
