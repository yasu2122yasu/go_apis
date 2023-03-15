/*
Q15
fruitsという要素数が3つのスライスを用意し、みかん、りんご、ぶどうという要素を代入して出力してください。
その後、append関数を利用してももという要素を追加して出力してください。
最後に、ぶどうとももの要素を削除して出力してください（方法は問いません。みかんとりんごという要素が出力されていればokです）
命名は自由に決めてください。
*/

package main

import "fmt"

func main() {
	a := []string{"みかん", "りんご", "ぶどう"}
	fmt.Printf("%s\n", a)
	b := append(a, "もも")
	fmt.Printf("%s\n", b)
	b = append(b[:2], b[:0]...)
	fmt.Printf("%s\n", b)
}
