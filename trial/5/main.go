/*
Q5.
数値型の変数a, bを用意して任意の数値を代入し、
a,bの和：c
a,bの差：d
a,bの積：e
a,bの商：f
a,bの余り：g
をそれぞれ出力してください。
*/
package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int = 50
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
