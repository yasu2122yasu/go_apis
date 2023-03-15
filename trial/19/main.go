/*
Q19
string型の変数aを宣言してください。
その後、strconvパッケージのAtoiメソッドを使ってaのデータ型をstring型からint型に変換してください
その際、aの中身が数値でなければ"データ型が数値型ではありません"、aの中身が数値ならその数値を出力してください。
ただしerror型を使用して実装してください。
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = "テスト"
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
