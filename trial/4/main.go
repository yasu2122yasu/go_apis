/*
Q4.
好きな数字を文字列型の変数で宣言し、出力してください。
その後、文字列型の変数のデータ型を数値型に変換し、出力してください。
最後に、最初に宣言した変数と型変換をした後の変数のデータ型も出力してください。
*/
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
