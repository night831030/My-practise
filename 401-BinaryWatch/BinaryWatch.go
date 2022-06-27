package main

import (
	f "fmt"
	b "math/bits"
	r "math/rand"
	t "time"
)

func main() {
	x := roll(11) - 1
	f.Print(readBinaryWatch(1))
	println(x)

}

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if b.OnesCount8(h)+b.OnesCount8(m) == turnedOn {
				ans = append(ans, f.Sprintf("%d:%02d", h, m))

			}
		}
	}
	return ans
}

func roll(x int) int { //骰子
	r.Seed(t.Now().UnixNano())
	return r.Intn(100000000000000)%x + 1
}
