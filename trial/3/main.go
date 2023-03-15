/*
Q3.
string型の変数を宣言して、"NPの設立は"を代入し、
int型の変数を宣言して、2000を代入し、
bool型の変数を宣言して、trueを代入し、
fmtパッケージのPrintfメソッドを使って
"NPの設立は2000年です:true"
と出力してください。
*/
package main

import "fmt"

func main() {
	var a string = "NPの設立は"
	var b int = 2000
	var c bool = true
	fmt.Printf("%s%d年です:%v", a, b, c)
}
