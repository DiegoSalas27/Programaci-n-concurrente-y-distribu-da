package main

import (
	"fmt"
)

const (
	K = 4
	MAX = 9
)

func main() {
	msg := "aassddddjjsjjjakksdggggshhhahhsdsss"
	inC := make(chan rune)
	pipe := make(chan rune)
	outC := make(chan rune) //runes son chars unicode
	go compress(inC, pipe)
	go output(pipe, outC)
	go func() {
		for _, c := range msg { //foreach
			inC <- c
		}
		close(inC)
	}()

	for c := range outC {
		fmt.Printf("%c", c)
	}
}

func compress(inC, pipe chan rune) {
	n := 0
	previous := <- inC
	for c := range inC {
		if c == previous && n < MAX - 1 {
			n++
		} else {
			if n > 0 {
				pipe <- rune(n+1+48)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
	close(pipe)
}

func output(pipe, outC chan rune) {
	m := 0
	for c := range pipe  {
		outC <- c
		m++
		if m >= K {
			outC <- '\n'
			m = 0
		} 
	}
	close(outC)
}