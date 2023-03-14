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

	u := makeStringSlice(p)

	flush := judgeFlush(p, u)

	straight := judgeStraight(num)

	dup := findDup(num)

	j := sliceToArray(dup)

	a := judge(j, num, straight, flush)

	json.NewEncoder(w).Encode(a)

	err := Unauthorized.New("ある認証の処理内で返されたエラー")
	fmt.Println(statusCode(err)) // 401
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
		// 正常系
	} else {
		// 配布されたコードが5枚でなければ処理を止める

	}

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
		// 正常系
	} else {
		// 配布されたコードが5枚でなければ処理を止めて、400エラーを出す
		err := Unauthorized.New("入力値が異なります")
		fmt.Println(statusCode(err))
	}

	var u []string = suitValidation(p)

	return u
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

func suitValidation(y []string) []string {
	// 排除したい文字がなければ正しいパラメータを表示するだけ
	for _, n := range y {
		switch n {
		case "Q":
			// 正常系
		case "D":
			// 正常系
		case "S":
			// 正常系
		case "H":
			// 正常系
		default:
			panic("柄が異なります")
		}
	}
	return y
}

// エラー処理
func statusCode(err error) int {
	switch GetType(err) {
	case ConnectionFailed:
		return http.StatusInternalServerError // 500
	case Unauthorized:
		return http.StatusUnauthorized // 401
	default:
		return http.StatusBadRequest // 400
	}
}
