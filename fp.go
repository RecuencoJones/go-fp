package main

func Filter(fn func(value interface{}) bool) func(arr interface{}) interface{} {
	return func(arr interface{}) interface{} {
		result := []interface{}{}

		for _, value := range arr.([]interface{}) {
			if fn(value) {
				result = append(result, value)
			}
		}

		return result
	}
}

func Map(fn func(value interface{}) interface{}) func(arr interface{}) interface{} {
	return func(arr interface{}) interface{} {
		result := []interface{}{}

		for _, value := range arr.([]interface{}) {
			result = append(result, fn(value))
		}

		return result
	}
}

func Reduce(fn func(accum interface{}, value interface{}) interface{}, initialValue interface{}) func(arr interface{}) interface{} {
	return func(arr interface{}) interface{} {
		result := initialValue

		for _, value := range arr.([]interface{}) {
			result = fn(result, value)
		}

		return result
	}
}

func Chain(fns ...func(interface{}) interface{}) func(interface{}) interface{} {
	return func(arr interface{}) interface{} {
		result := arr

		for _, fn := range fns {
			result = fn(result)
		}

		return result
	}
}
