/*
Q14
1分ごとに2倍になっていく栗まんじゅうがあります．1つの栗まんじゅうは10分後にいくつになっているかを求めるプログラムを書いてください
*/
package main

import "fmt"

func main() {
	var s int = 1
	var k int = 1
	for k = 1; k <= 10; k++ {
		s = s * 2
	}
	fmt.Printf("%d\n", s)
}
