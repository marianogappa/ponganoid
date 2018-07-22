package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

func drawGameEvery(state *gameState, interval time.Duration) {
	var drawTicker = time.NewTicker(interval)
	for {
		select {
		case <-drawTicker.C:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			var (
				minX, minY = 0, 1
				maxX, maxY = termbox.Size()
			)
			drawHealth(state.players[0].health, state.maxHealth, state.healthChangeTransient)
			drawHearts(minX, minY, maxX, maxY, state.hearts)
			drawPlayers(minX, minY, maxX, maxY, state.players, state.healthChangeTransient)
			if state.healthChangeTransient > 0 {
				state.healthChangeTransient--
			} else if state.healthChangeTransient < 0 {
				state.healthChangeTransient++
			}
			termbox.Flush()
		}
	}
}

func drawPlayers(minX, minY, maxX, maxY int, players []player, healthChangeTransient int) {
	for i, p := range players {
		if p.x < minX || p.x >= maxX || p.y < minY || p.y >= maxY {
			continue
		}
		var (
			c  = '∆'
			fg = p.c
		)
		if i == 0 {
			c = 'o'
			if healthChangeTransient < 0 {
				fg = termbox.ColorRed
			} else if healthChangeTransient > 0 {
				fg = termbox.ColorGreen
			}
		}
		termbox.SetCell(p.x, p.y, c, fg, termbox.ColorBlack)
	}
}

func drawHearts(minX, minY, maxX, maxY int, hearts []int) {
	for _, h := range hearts {
		if h <= 1 {
			continue
		}
		var x, y = h % maxX, h / maxX
		if x < minX || x >= maxX || y < minY || y >= maxY {
			continue
		}
		termbox.SetCell(x, y, '♥', termbox.ColorGreen, termbox.ColorBlack)
	}
}

func drawHealth(h, mh, transient int) {
	termbox.SetCell(0, 0, '♥', termbox.ColorGreen, termbox.ColorBlack)
	termbox.SetCell(2, 0, '[', termbox.ColorWhite, termbox.ColorBlack)
	var c = termbox.ColorGreen
	if transient > 0 {
		c = termbox.ColorBlue
	} else if transient < 0 {
		c = termbox.ColorRed
	}
	for i := 0; i < h; i++ {
		termbox.SetCell(i+3, 0, ' ', termbox.ColorBlack, c)
	}
	termbox.SetCell(mh+3, 0, ']', termbox.ColorWhite, termbox.ColorBlack)
}
