package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Poker struct {
	Title string `json:"title"`
}

var pokers []Poker

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var poker Poker

	_ = json.NewDecoder(r.Body).Decode(&poker)

	p := poker.Title

	num := makeIntSlice(p)

	test2 := makeStringSlice(p)

	flush := judgeFlush(p, test2)

	straight := judgeStraight(num)

	dup := findDup(num)

	j := sliceToArray(dup)

	a := judge(j, num, straight, flush)

	json.NewEncoder(w).Encode(a)
}

func main() {
	// Initiate Router
	r := mux.NewRouter()

	r.HandleFunc("/", process).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func makeIntSlice(cards string) []int {
	r, _ := regexp.Compile("[^0-9| ]")

	q := strings.Split(r.ReplaceAllString(cards, ""), " ")

	s := [5]string{}
	copy(s[:], q)

	if len(s) == 5 {
		fmt.Printf("ok %T\n", s)
	} else {
		fmt.Printf("bad %T\n", s)
		panic("配布されたカードの数が5枚ではありません")
	}

	fmt.Printf("テスト%s\n", s)

	fmt.Printf("%T\n", s)
	fmt.Println(cap(s))

	var ab = []int{}

	for _, i := range q {
		j, _ := strconv.Atoi(i)
		ab = append(ab, j)
	}

	sort.Ints(ab)

	return ab
}

func makeStringSlice(cards string) []string {
	r, _ := regexp.Compile("[^A-Z| ]")

	p := strings.Split(r.ReplaceAllString(cards, ""), " ")

	if len(p) == 5 {
		fmt.Printf("ok %T\n", p)
	} else {
		fmt.Printf("bad %T\n", p)
		panic("配布されたカードの数が5枚ではありません")
	}

	// p := removeImproperNumberOfSheets(s)

	return p
}

func judgeFlush(cards string, test2 []string) int {
	flush := 0

	if strings.Count(cards, test2[0]) == len(test2) {
		flush = 1
	} else {
		flush = 0
	}

	return flush

}

func judgeStraight(num []int) int {
	straight := 0

	if num[1] == num[0]+1 && num[2] == num[0]+2 && num[3] == num[0]+3 && num[4] == num[0]+4 || num[0] == 1 && num[1] == 1 {
		straight = 1
	} else {
		straight = 0
	}
	return straight
}

func findGreat(array []int) int {
	greatest := 0
	for i := 0; i < len(array); i++ {
		if array[i] > greatest {
			greatest = array[i]
		}
	}
	return greatest + 1
}

func findDup(array []int) []int {
	dup := make([]int, findGreat(array))

	for i := 0; i < len(array); i++ {
		dup[array[i]] += 1
	}

	sort.Ints(dup)
	return dup
}

func sliceToArray(dup []int) []int {
	j := []int{}

	for _, i := range dup {
		if i != 0 {
			j = append(j, i)
		}
	}
	return j
}

func judge(j []int, num []int, straight, flush int) string {
	//二次元スライスを作成する。2ペアなどの判定で利用。

	p := [][]int{{1, 4}, {2, 3}, {1, 1, 3}, {1, 2, 2}, {1, 1, 1, 2}, {1, 1, 1, 1, 1}}
	var a string = ""

	if reflect.DeepEqual(j, p[0]) {
		a = "4カード"
	} else if reflect.DeepEqual(j, p[1]) {
		a = "フルハウス"
	} else if reflect.DeepEqual(j, p[2]) {
		a = "スリーカード"
	} else if reflect.DeepEqual(j, p[3]) {
		a = "ツーペア"
	} else if reflect.DeepEqual(j, p[4]) {
		a = "ワンペア"
	} else if reflect.DeepEqual(j, p[5]) {
		// 比較をするにはスライスを配列に変換する必要がある。numスライスをs配列にコピーする。
		s := [5]int{}
		copy(s[:], num)

		if s == [5]int{1, 10, 11, 12, 13} && flush == 1 {
			a = "ロイヤルストレートフラッシュ"
		} else if s != [5]int{1, 10, 11, 12, 13} && flush == 1 && straight == 1 {
			a = "ストレートフラッシュ"
		} else if flush == 1 {
			a = "フラッシュ"
		} else if straight == 1 {
			a = "ストレート"
		} else {
			a = "ノーペア"
		}
	}
	return a
}

// 文字列型、数値型のスライス共通で、4枚以下・6枚以上の異常系を取り除く関数。スライス→関数→スライスにする必要がある。
// func removeImproperNumberOfSheets(a []string) [5]string {
// 	b := [5]string{}
// 	copy(b[:], a)
// 	fmt.Printf("b1%T\n", b)
// 	if cap(b) != 5 {
// 		panic("与えられたカードの枚数が5枚ではない。")
// 	}

// 	return b
// }
