// numberという変数を数値型で作って好きな数字を代入し、fmtパッケージのPrintfメソッドを使って出力しよう
package main

import "fmt"

func main() {
	var number int = 12
	fmt.Printf("%d", number)
}
