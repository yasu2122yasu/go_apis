/*
Q7.
ぶどう、りんご、みかんの中で一番好きな果物を変数aに代入してください。
switch文とcase文を用いて、それぞれの果物が変数aに代入された際に、"好きな果物は〇〇です"と出力されるようプログラムを書いてください。
のcaseにも当てはまらない場合は"好きな果物はありません"と出力してください。
*/
package main

import "fmt"

func main() {
	a := "みかん"

	switch a {
	case "ぶどう":
		fmt.Printf("好きな果物はです%s\n", a)
	case "りんご":
		fmt.Printf("好きな果物はです%s\n", a)
	case "みかん":
		fmt.Printf("好きな果物はです%s\n", a)
	default:
		fmt.Println("好きな食べ物はありません")
	}
}
