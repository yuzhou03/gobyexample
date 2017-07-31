// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

/**
 * 不固定参数的加法
 */
func sumImproved(items ... int) int {
	var result int
	for _, val := range items {
		result += val
	}
	return result
}

/**
 * 传递数组参数
 */

func sumArray(array [10]int) int {
	var result int
	for _, val := range array {
		result += val
	}

	return result
}

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	p := fmt.Println

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = sumImproved(1, 2, 3, 4, 5)
	p("result:", res)

	arr := [10] int{1, 2, 3, 4, 5}
	p("result:", sumArray(arr))

}
