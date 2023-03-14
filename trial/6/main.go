// 自分の誕生日の下二桁を用意し、変数に代入。15より小さければ、パラメータは0です。15より大きければパラメータは1です。と出力してください。
// if-else文を使ってください。
package main

import "fmt"

func main() {
	var b int = 12
	if b < 15 {
		fmt.Println("月の前半が誕生日です")
	} else {
		fmt.Println("月の後半が誕生日です")
	}
}
