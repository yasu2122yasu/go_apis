package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
    cards := "Q5 H12 H10 H11 H1"
    test := makeIntSlice(cards)
    test2 := makeStringSlice(cards)

    flush := 0

    if strings.Count(cards, test2[0]) == len(test2) {
      flush = 1
    } else {
      flush = 0
    }

    sort.Ints(test)

    straight := 0

    if (test[1] == test[0] +1 && test[2] == test[0] + 2 && test[3] == test[0] +3 && test[4] == test[0] + 4 ||  test[0] == 1 && test[1] == 1){
        straight = 1
    } else {
        straight = 0
    }

    tbt := findDup(test)

    sort.Ints(tbt)

		j := []int{}

		for _,i := range tbt {
			if i != 0 {
			  j = append(j, i)
	   	}
		}

		//二次元スライスを作成する。2ペアなどの判定で利用。

		p := [][]int{{1, 4}, {2, 3}, {1, 1, 3 }, {1, 2, 2}, {1, 1, 1, 2}, {1, 1, 1, 1, 1}}

		if reflect.DeepEqual(j, p[0]){
			fmt.Println("4カードです")
		} else if reflect.DeepEqual(j, p[1]){
			fmt.Println("フルハウスです")
		} else if reflect.DeepEqual(j, p[2]){
			fmt.Println("スリーカードです")
		} else if reflect.DeepEqual(j, p[3]){
			fmt.Println("ツーペアです")
		} else if reflect.DeepEqual(j, p[4]){
			fmt.Println("ワンペアです")
		} else if reflect.DeepEqual(j, p[5]){
			// 比較をするにはスライスを配列に変換する必要がある。testスライスをs配列に変更する。
			s := [5]int{}
			copy(s[:], test)

			if (s == [5]int{1,10,11,12,13} && flush == 1) {
				fmt.Println("ロイヤルストレート")
			} else if (s != [5]int{1,10,11,12,13} && flush == 1 && straight == 1) {
        fmt.Println("ストレートフラッシュ")
		  } else if (flush == 1){
				fmt.Println("フラッシュ")
			} else if (straight == 1){
				fmt.Println("ストレート")
			} else {
				fmt.Println("ノーペア")
			}
		}

}

func makeIntSlice(cards string) []int{
    regex, _ := regexp.Compile("[^0-9| ]")
    stest := strings.Split(regex.ReplaceAllString(cards, ""), " ")

    var ab = []int{}

    for _, i := range stest {
        j, _ := strconv.Atoi(i)
        ab = append(ab, j)
    }

    return ab
}

func makeStringSlice(cards string) []string{
    regex, _ := regexp.Compile("[^A-Z| ]")
    snum := strings.Split(regex.ReplaceAllString(cards, ""), " ")
    return snum
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
	return dup
}

// func judge(j int, test []int, straight, flush int) string {

// }
