package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	TOTAL = 108
	WAN   = 3
	TONG  = 2
	TIAO  = 1
	HAND  = 13
	SEAT  = 4
)

var CARDS = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39}

type Cards struct {
	TableCards []int
	Hand1      []int
	Hand2      []int
	Hand3      []int
	Hand4      []int
	Remainder1 []int
	Remainder2 []int
	Remainder3 []int
	Remainder4 []int
}

func init() {
	CARDS = append(CARDS, CARDS...)
	CARDS = append(CARDS, CARDS...)
}
func main() {
	cards := &Cards{}
	cards.Deal()
	//fmt.Println(cards.TableCards, len(cards.TableCards))
	//fmt.Println(cards.Hand1, len(cards.Hand1))
	//fmt.Println(cards.Hand2, len(cards.Hand2))
	//fmt.Println(cards.Hand3, len(cards.Hand3))
	//fmt.Println(cards.Hand4, len(cards.Hand4))
	//fmt.Println(cards.Get())
	a := []int{11, 11, 11, 12, 12, 13, 14, 15, 16, 17, 18, 19, 19, 19}
	fmt.Println(cards.hu(a))
}
func (this *Cards) Deal() {
	this.TableCards = this.draw()
	this.Hand4 = this.TableCards[0:HAND]
	this.Hand2 = this.TableCards[HAND : HAND*2]
	this.Hand3 = this.TableCards[HAND*2 : HAND*3]
	this.Hand1 = this.TableCards[HAND*3 : HAND*4+1]
	this.TableCards = this.TableCards[HAND*4+1 : TOTAL]
}
func (this *Cards) hu(list []int) bool {
	sort.Ints(list)
	le := len(list)
	fmt.Println(list)
	for n := 0; n < le-1; n++ {
		jiang := 0
		if list[n] == list[n+1] {
			jiang = list[n]
			list[n] = 0
			list[n+1] = 0
			for i := 0; i < le; i++ {
				for j := i + 1; j < le; j++ {
					for k := j + 1; k < le; k++ {
						if ((list[i]+1) == list[j] && (list[j]+1) == list[k]) || (list[i] > 0 && list[i] == list[j] && list[j] == list[k]) {
							list[i] = 0
							list[j] = 0
							list[k] = 0
						}
					}
				}
			}
			num := 0
			for i := 0; i < le; i++ {
				if list[i] > 0 {
					num = list[i]
				}
			}
			if num == 0 {
				return true
			}
			list[n] = jiang
			list[n+1] = jiang
		}

	}
	return false
}
func (this *Cards) Out(seat int, card int) {
	if seat == 1 {
		this.Remainder1 = append(this.Remainder1, card)
	} else if seat == 2 {

		this.Remainder2 = append(this.Remainder3, card)
	} else if seat == 3 {

		this.Remainder3 = append(this.Remainder3, card)
	} else if seat == 4 {

		this.Remainder4 = append(this.Remainder4, card)
	} else {

	}
}
func (this *Cards) Get() int {
	card := 0
	l := len(this.TableCards)
	if l > 0 {
		card = this.TableCards[l-1]
		this.TableCards = this.TableCards[0 : l-1]
	}
	return card
}

func (this *Cards) LeftCount() int {
	return len(this.TableCards)
}
func (this *Cards) draw() []int {
	t := time.Now()
	rand.Seed(t.UnixNano())
	d := make([]int, 0, TOTAL)
	copy(d, CARDS)

	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
