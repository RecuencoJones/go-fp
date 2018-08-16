package main

import "fmt"

func Example1() {
	fmt.Println("Example 1")
	fmt.Println("Sum of even numbers in array")

	arr := []interface{}{1, 2, 3, 4}

	filteredArr := Filter(func(value interface{}) bool {
		return value.(int)%2 == 0
	})(arr)

	fmt.Println(filteredArr)

	mappedArr := Map(func(value interface{}) interface{} {
		return value.(int) * 2
	})(filteredArr)

	fmt.Println(mappedArr)

	reducedRes := Reduce(func(accum interface{}, value interface{}) interface{} {
		return accum.(int) + value.(int)
	}, 0)(mappedArr)

	fmt.Println(reducedRes)
}
