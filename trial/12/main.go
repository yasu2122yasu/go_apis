/*
Q12
1〜100までの和を計算するプログラムを書いてください．
*/
package main

import "fmt"

func main() {
	var s int = 0
	for i := 1; i <= 100; i++ {
		s = s + i
	}
	fmt.Printf("%d\n", s)
}
