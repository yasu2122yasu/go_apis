/*
問題7の処理をmain関数だけではなく、fruitSelect関数に分けて書いてください（main関数は用意してあります）
*/
package main

import "fmt"

func main() {
	b := fruitSelect()
	fmt.Printf("好きな果物は%sです\n", b)
}

func fruitSelect() string {
	a := "みかん"

	switch a {
	case "ぶどう":
		a = "ぶどう"
	case "りんご":
		a = "りんご"
	case "みかん":
		a = "みかん"
	default:
		fmt.Println("好きな食べ物はありません")
	}
	return a
}
