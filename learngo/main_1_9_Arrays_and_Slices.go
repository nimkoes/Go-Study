package main

import "fmt"

func main() {
	// array 사용 방법 1
	example_array_ver1 := [5]string{"arr_nimkoes", "arr_go", "arr_java"}

	// array 사용 방법 2
	example_array_ver2 := [...]string{"kim", "lee", "park", "choi"}

	// slice 사용 방법
	example_slice := []string{"nimkoes", "go", "java"}
	example_slice = append(example_slice, "new_elem") // 배열 요소를 추가

	fmt.Println(example_array_ver1)
	fmt.Println(example_array_ver2)
	fmt.Println(example_slice)
}
