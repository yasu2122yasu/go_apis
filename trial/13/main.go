/*
Q13
1〜30までの数字を順番に表示し，もしも3の倍数だったら「3の倍数です」と表示するプログラムを書いてください
*/

package main

import "fmt"

func main() {
	var i int
	for i = 1; i <= 30; i++ {
		fmt.Println(i)
		if i%3 == 0 {
			fmt.Println("3の倍数です")
		}
	}
}
