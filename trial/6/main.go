// 自分の誕生日の下二桁を用意し、変数に代入。15より小さければ、パラメータは0です。15より大きければパラメータは1です。と出力してください。
package main

import "fmt"

func main() {
	var b int = 12
	if b < 15 {
		fmt.Println("パラメータは0です")
	} else {
		fmt.Println("パラメータは1です")
	}
}
