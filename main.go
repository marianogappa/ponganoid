package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	mustInitTermbox()
	var (
		keys    = make(chan termbox.Key)
		actions = make(chan humanActions)
		state   = initGameState()
	)
	go drawGameEvery(state, 40*time.Millisecond)
	go playGame(state, actions, 80*time.Millisecond)
	go interpretHumanActions(keys, actions)
	mustReadKeyboard(keys)
}
