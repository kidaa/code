package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	TOTAL = 108
	WAN   = 2
	TONG  = 1
	TIAO  = 0
	HAND  = 13
	SEAT  = 4
)

var CARDS = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
}

type Cards struct {
	TableCards []byte
	Hand1      []byte
	Hand2      []byte
	Hand3      []byte
	Hand4      []byte
	Remainder1 []byte
	Remainder2 []byte
	Remainder3 []byte
	Remainder4 []byte
}
type sortByte []byte

func (s sortByte) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByte) Len() int {
	return len(s)
}
func main() {
	cards := &Cards{}
	cards.Deal()
	//fmt.Println(cards.Get())
	a := []byte{0x21, 0x21, 0x12, 0x13, 0x13, 0x14, 0x14, 0x16, 0x16, 0x18, 0x18, 0x19, 0x19}
	//a := []byte{0x21, 0x21, 0x21, 0x12, 0x12, 0x12, 0x14, 0x14, 0x14, 0x17, 0x17, 0x17, 0x19, 0x19}
	//a := []byte{0x21, 0x21, 0x22, 0x22, 0x23, 0x23, 0x14, 0x14, 0x15, 0x15, 0x16, 0x16, 0x17, 0x17}
	fmt.Println(cards.ting(a))
	fmt.Println(cards.Patterns(0x21))
}
func (this *Cards) Deal() {
	this.TableCards = this.draw()
	this.Hand4 = this.TableCards[0:HAND]
	this.Hand2 = this.TableCards[HAND : HAND*2]
	this.Hand3 = this.TableCards[HAND*2 : HAND*3]
	this.Hand1 = this.TableCards[HAND*3 : HAND*4+1]
	this.TableCards = this.TableCards[HAND*4+1 : TOTAL]
}
func (this *Cards) Patterns(card byte) int {
	return int(card >> 4)
}

func (this *Cards) ting(original []byte) bool {
	sort.Sort(sortByte(original))
	le := len(original)
	for n := 0; n < le-1; n++ {
		if original[n] == original[n+1] {
			list := make([]byte, le, le)
			copy(list, original)
			list[n] = 0x00
			list[n+1] = 0x00
			for i := 0; i < le; i++ {
				for j := i + 1; j < le; j++ {
					for k := j + 1; k < le; k++ {
						if list[i] > 0 && ((list[i]+1) == list[j] && (list[j]+1) == list[k]) || (list[i] == list[j] && list[j] == list[k]) {
							list[i] = 0x00
							list[j] = 0x00
							list[k] = 0x00
						}
					}
				}
			}
			arr := make([]byte, 0, le)
			for i := 0; i < le; i++ {
				if list[i] > 0 {
					arr = append(arr, list[i])
				}
			}
			if len(arr) == 1 {
				return true
			} else if len(arr) == 2 && (arr[0]+1 == arr[1] || arr[0] == arr[1]) {
				return true
			}
		}
	}

	list := make([]byte, le, le)
	copy(list, original)
	for n := 0; n < le-1; n++ {
		if original[n] == original[n+1] {
			list[n] = 0x00
			list[n+1] = 0x00
		}
	}
	count := 0
	for i := 0; i < le; i++ {
		if list[i] > 0 {
			count = count + 1
		}
	}
	if count == 1 {
		return true
	}
	return false
}
func (this *Cards) hu(original []byte) bool {
	sort.Sort(sortByte(original))
	le := len(original)
	for n := 0; n < le-1; n++ {
		if original[n] == original[n+1] {
			list := make([]byte, le, le)
			copy(list, original)
			list[n] = 0x00
			list[n+1] = 0x00
			for i := 0; i < le; i++ {
				for j := i + 1; j < le; j++ {
					for k := j + 1; k < le; k++ {
						if list[i] > 0 && ((list[i]+1) == list[j] && (list[j]+1) == list[k]) || (list[i] == list[j] && list[j] == list[k]) {
							list[i] = 0x00
							list[j] = 0x00
							list[k] = 0x00
						}
					}
				}
			}
			num := false
			for i := 0; i < le; i++ {
				if list[i] > 0 {
					num = true
					break
				}
			}
			if !num {
				return true
			}
		}
	}
	list := make([]byte, le, le)
	copy(list, original)
	for n := 0; n < le-1; n++ {
		if original[n] == original[n+1] {
			list[n] = 0x00
			list[n+1] = 0x00
		}
	}
	num := false
	for i := 0; i < le; i++ {
		if list[i] > 0 {
			num = true
			break
		}
	}
	if !num {
		return true
	}

	return false
}
func (this *Cards) Out(seat int, card byte) {
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
func (this *Cards) Get() byte {
	var card byte = 0
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
func (this *Cards) draw() []byte {
	t := time.Now()
	rand.Seed(t.UnixNano())
	d := make([]byte, TOTAL, TOTAL)
	copy(d, CARDS)

	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
