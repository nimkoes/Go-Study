package main

import "fmt"

func superAdd(numbers ...int) int {
	fmt.Println("without index")
	for number := range numbers {
		fmt.Print(number)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("with index")
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("ignore index")
	for _, number := range numbers {
		fmt.Println(number)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("without range")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
