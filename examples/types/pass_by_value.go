package main

import "fmt"

func main() {
	pass_by_value()
	pass_by_reference()
}
func pass_by_value() {
	var a = [3] int{1, 2, 3}
	var b = a //Go中的数组和基本类型没有区别，是很纯粹的值类型
	b[1]++
	fmt.Println(a, b)
}

func pass_by_reference(){
	var a = [3] int{1,2,3}
	var b = &a
	b[1]++
	fmt.Println(a,*b)
}