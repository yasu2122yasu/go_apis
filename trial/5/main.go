// 任意の数字を文字列型で宣言し、出力する。データ型を数値型に変更して、もう一度宣言する。最後にnum1,num2のデータ型を出力すること
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
