package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

func mustReadKeyboard(ch chan termbox.Key) {
	defer termbox.Close()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return
			}
			ch <- ev.Key
		case termbox.EventError:
			termbox.Close()
			log.Fatal(ev.Err)
		}
	}
}

type humanActions int

const (
	actionMoveUp humanActions = iota
	actionMoveLeft
	actionMoveRight
	actionMoveDown
)

func interpretHumanActions(keys chan termbox.Key, actions chan humanActions) {
	for key := range keys {
		switch key {
		case termbox.KeyArrowUp:
			actions <- actionMoveUp
		case termbox.KeyArrowLeft:
			actions <- actionMoveLeft
		case termbox.KeyArrowRight:
			actions <- actionMoveRight
		case termbox.KeyArrowDown:
			actions <- actionMoveDown
		default:
		}
	}
}
