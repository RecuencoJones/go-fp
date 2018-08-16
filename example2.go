package main

import "fmt"

func Example2() {
	fmt.Println("Example 2")
	fmt.Println("Sum of even numbers in array (with chain)")

	arr := []interface{}{1, 2, 3, 4}

	result := Chain(
		Filter(func(value interface{}) bool {
			return value.(int)%2 == 0
		}),
		Map(func(value interface{}) interface{} {
			return value.(int) * 2
		}),
		Reduce(func(accum interface{}, value interface{}) interface{} {
			return accum.(int) + value.(int)
		}, 0),
	)(arr)

	fmt.Println(result)
}
