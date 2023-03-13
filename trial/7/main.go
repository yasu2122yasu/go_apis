// 好きな果物を3つ選び（頭の中で）、その中で一番好きな果物を変数に代入してください。その後、switch文とcase文を用いて、"好きな果物は〇〇です"と出力してください。
// 好きな果物がない人もいると思うので、どのcaseにも当てはまらない場合は"好きな果物はありません"と出力してください。
// 出力はされませんが、残りの2つの果物を選んだときの処理と、defaultの処理も書いてください。

package main

import "fmt"

func main() {
	a := "みかん"

	switch a {
	case "ぶどう":
		fmt.Printf("%s\n", a)
	case "りんご":
		fmt.Printf("%s\n", a)
	case "みかん":
		fmt.Printf("%s\n", a)
	default:
		fmt.Println("好きな食べ物はありません")
	}
}
