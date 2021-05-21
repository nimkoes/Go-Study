package main

import (
	"fmt"
	"myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First Word"}

	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
