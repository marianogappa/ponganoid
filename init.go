package main

import (
	"math/rand"

	termbox "github.com/nsf/termbox-go"
)

type player struct {
	x, y, dx, dy, speed, health int
	c                           termbox.Attribute
}

type gameState struct {
	players               []player
	hearts                []int
	maxHealth             int
	healthChangeTransient int // used to make a visual effect on health change
}

func initGameState() *gameState {
	var (
		enemies    = make([]player, 0)
		minX, minY = 0, 1
		maxX, maxY = termbox.Size()
		hearts     = make([]int, 0)
	)

	for i := 0; i < int(float64(maxX*maxY)*0.01); i++ {
		var x, y = randXY(minX, minY, maxX, maxY)
		enemies = append(enemies, player{
			x:      x,
			y:      y,
			dx:     1 + (-2 * rand.Intn(2)),
			dy:     1 + (-2 * rand.Intn(2)),
			speed:  2,
			health: 0,
			c:      termbox.ColorRed,
		})
		x, y = randXY(minX, minY, maxX, maxY)
		hearts = append(hearts, y*maxX+x)
	}

	return &gameState{
		players: append([]player{ // player 0 is the human player
			{x: 1, y: 1, dx: 1, dy: 1, speed: 1, health: 10, c: termbox.ColorWhite}},
			enemies...),
		hearts:                hearts,
		maxHealth:             20,
		healthChangeTransient: 0,
	}
}

func randXY(minX, minY, maxX, maxY int) (int, int) {
	for {
		var x, y = rand.Intn(maxX), rand.Intn(maxY-1) + 1
		if (x+y)%2 == 0 { // Otherwise position can never be reached
			return x, y
		}
	}
}
