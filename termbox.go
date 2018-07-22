package main

import (
	"fmt"
	"log"

	termbox "github.com/nsf/termbox-go"
)

func mustInitTermbox() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func printTb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func logf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	printTb(0, 0, termbox.ColorWhite, termbox.ColorBlack, s)
}
