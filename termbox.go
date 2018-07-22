package main

import (
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
