package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TOTAL = 108
	WAN   = 3
	TONG  = 2
	TIAO  = 1
	HAND  = 13
)

var CARDS = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39}

type Cards struct {
	List []int
}

func main() {
	CARDS = append(CARDS, CARDS...)
	CARDS = append(CARDS, CARDS...)
	aw := draw()
	left, hand1, hand2, hand3, hand4 := slice(aw)
	fmt.Println(aw, len(aw))
	fmt.Println(left, len(left))
	fmt.Println(hand1, len(hand1))
	fmt.Println(hand2, len(hand2))
	fmt.Println(hand3, len(hand3))
	fmt.Println(hand4, len(hand4))

}
func (this *Cards) get() int {

}
func (this *Cards) slice(list []int) (left []int, hand1 []int, hand2 []int, hand3 []int, hand4 []int) {
	hand1 = list[0:HAND]
	hand2 = list[HAND : HAND*2]
	hand3 = list[HAND*2 : HAND*3]
	hand4 = list[HAND*3 : HAND*4+1]
	left = list[HAND*4+1 : TOTAL]
	return
}
func (this *Cards) draw() []int {
	t := time.Now()
	rand.Seed(t.UnixNano())
	for i := range CARDS {
		j := rand.Intn(i + 1)
		CARDS[i], CARDS[j] = CARDS[j], CARDS[i]
	}
	d := make([]int, TOTAL, TOTAL)
	copy(d, CARDS)
	return d
}
