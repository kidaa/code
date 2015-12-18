package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CARDS = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39}

func main() {
	CARDS = append(CARDS, CARDS...)
	CARDS = append(CARDS, CARDS...)
	draw()
	fmt.Println(len(CARDS), CARDS)
	draw()
	fmt.Println(len(CARDS), CARDS)
	draw()
	fmt.Println(len(CARDS), CARDS)

}
func draw() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	for i := range CARDS {
		j := rand.Intn(i + 1)
		CARDS[i], CARDS[j] = CARDS[j], CARDS[i]
	}
}
