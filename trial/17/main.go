// int型のポインタ変数aを用意する
// int型の変数bを用意して任意の数字を代入する。
// bをaに代入する（変数aとbではデータ型が異なるので注意すること）
// aに代入された値とアドレスを出力する
package main

import "fmt"

func main() {
	var a *int
	b := 7
	a = &b

	fmt.Println(*a)
	fmt.Println(a)
}
