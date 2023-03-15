// りんご、ぶどう、みかんという要素が入った文字列型のスライスを定義し、出力してください。
package main

import "fmt"

func main() {
	var fruits = []string{"りんご", "ぶどう", "みかん"}
	for _, i := range fruits {
		fmt.Println(i)
	}
}
