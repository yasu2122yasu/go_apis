// スライスの要素に不要な要素があったときに処理を止めるコード
package main

import "fmt"

func main() {
	// 排除したいコードをパラメータに含めてテストをする
	y := []string{"Q", "D", "Q", "T", "A"}

	var test2 string = suitValidation(y)

	fmt.Printf("%s", test2)

}

func suitValidation(y []string) string {
	var t string
	// 排除したい文字がなければ正しいパラメータを表示するだけ
	for _, n := range y {
		switch n {
		case "Q":
			fmt.Printf("%s\n", n)
		case "D":
			fmt.Printf("%s\n", n)
		case "S":
			fmt.Printf("%s\n", n)
		case "H":
			fmt.Printf("%s\n", n)
			//異常系ならpanic()で止める
		default:
			panic("柄が異なります")
		}
	}
	return t
}
