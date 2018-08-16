package main

import (
	"fmt"
	"strings"
)

func Example3() {
	fmt.Println("Example 3")
	fmt.Println("Map array of objects")

	type SomeObject struct {
		foo string
		bar string
	}

	arr := []interface{}{
		SomeObject{
			foo: "foo1",
			bar: "bar1",
		},
		SomeObject{
			foo: "foo2",
			bar: "bar2",
		},
	}

	result := Chain(
		Map(func(obj interface{}) interface{} {
			return obj.(SomeObject).foo
		}),
		Map(func(str interface{}) interface{} {
			return strings.ToUpper(str.(string))
		}),
	)(arr)

	fmt.Println(result)
}
