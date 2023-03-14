// 任意の数字を文字列型で宣言し、出力する。データ型を数値型に変更して、もう一度宣言する。最後にnum1,num2のデータ型を出力すること
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 string = "355"
	fmt.Println(num1)
	var num2 int
	num2, _ = strconv.Atoi(num1)
	fmt.Println(num2)
	fmt.Printf("%T,	%T", num1, num2)
}
