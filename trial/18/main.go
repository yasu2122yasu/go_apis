/*
Q18
Personという名前がついた構造体を宣言してください。中身はfirstName（string型）、lastName（int型）、age（int型）としてください。
その後、この構造体を使って"Wataru Satou 15歳です"と出力してください。
*/
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var a Person
	a.firstName = "Wataru"
	a.lastName = "Satou"
	a.age = 15

	fmt.Printf("%s %s %d歳です。", a.lastName, a.firstName, a.age)
}
