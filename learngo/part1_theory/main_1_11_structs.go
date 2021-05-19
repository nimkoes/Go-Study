package main

import "fmt"

// Go 의 struct 에는 constructor (생성자) 가 없다.
// 우리가 constructor 를 직접 실행해 줘야 한다. -> 생성자가 없다고 했는데 실행해줘야 한다는 말을 이해하지 못함.
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	myFood := []string{"kimchi", "ramen"}

	// struct 사용 방법 1, 순서대로 타입에 맞게 값을 입력 (권장하지 않음)
	nimkoes_ver1 := person{"nimkoes", 17, myFood}
	nimkoes_ver2 := person{"nimkoes", 21, []string{"kimchi_2", "ramen_2"}}

	// struct 사용 방법 2, key 를 명시하는 방법 (권장하는 방법)
	nimkoes_ver3 := person{name: "nk", favFood: []string{"I", "eat", "something"}}

	// 방법 1 과 방법 2 를 혼용할 수 없다. (mixture of field:value and value elements in struct literalcompilerMixedStructLit)
	// nimkoes_ver4 := person{name: "nk", []string{"I", "eat", "something"}}

	fmt.Println(nimkoes_ver1.name, nimkoes_ver1.age, nimkoes_ver1.favFood)
	fmt.Println(nimkoes_ver1)
	fmt.Println(nimkoes_ver2)

	fmt.Println(nimkoes_ver3)
}
