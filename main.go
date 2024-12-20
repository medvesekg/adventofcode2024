package main

import (
	"seehuhn.de/go/ncurses"
)

func main() {
	win := ncurses.Init()
	win.Print("Hello, World!")

	for true {
		ch := win.GetCh()
		if ch == ncurses.KeyUp {
			win.Print("up")
		}
	}
	defer ncurses.EndWin()

}
