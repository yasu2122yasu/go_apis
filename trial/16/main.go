// main関数内でaというスライスを用意し，100,200,300を追加してください．
// その後、calculateSum関数を作って、aの合計を計算して出力するプログラムを書いてください。
// その後、calculateAverage関数を作って、aの平均を計算して出力するプログラムを書いてください。

package main

import "fmt"

func main() {
	a := []int{100, 200, 300}
	calculateSum(a)
	calculateAverage(a)
}

func calculateSum(a []int) {
	var s int = 0
	for _, i := range a {
		s = s + i
	}
	fmt.Printf("%d\n", s)
	return
}

func calculateAverage(a []int) {
	var s int = 0
	for _, i := range a {
		s = s + i
	}
	var d = s / len(a)
	fmt.Printf("%d\n", d)
	return
}
