package main

import (
	"fmt"
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func drawGameEvery(gs *gameState, interval time.Duration) {
	var drawTicker = time.NewTicker(interval)
	for {
		select {
		case <-drawTicker.C:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			var (
				minX, minY = 0, 1
				maxX, maxY = termbox.Size()
			)
			drawHealth(gs.players[0].health, gs.maxHealth, gs.healthChangeTransient)
			for _, h := range gs.hearts {
				if h <= 1 {
					continue
				}
				var x, y = h % maxX, h / maxX
				if x < minX || x >= maxX || y < minY || y >= maxY {
					continue
				}
				termbox.SetCell(x, y, '♥', termbox.ColorGreen, termbox.ColorBlack)
			}
			for i, p := range gs.players {
				if p.x < minX || p.x >= maxX || p.y < minY || p.y >= maxY {
					continue
				}
				var (
					c  = '∆'
					fg = p.c
				)
				if i == 0 {
					c = 'o'
					if gs.healthChangeTransient < 0 {
						fg = termbox.ColorRed
					} else if gs.healthChangeTransient > 0 {
						fg = termbox.ColorGreen
					}
				}
				termbox.SetCell(p.x, p.y, c, fg, termbox.ColorBlack)
			}
			if gs.healthChangeTransient > 0 {
				gs.healthChangeTransient--
			} else if gs.healthChangeTransient < 0 {
				gs.healthChangeTransient++
			}
			termbox.Flush()
		}
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

func playGame(state *gameState, actions chan humanActions, interval time.Duration) {
	var (
		gameTicker  = time.NewTicker(interval)
		frameCount  = 0
		humanPlayer = &state.players[0]
	)
	for {
		select {
		case action := <-actions:
			switch action {
			case actionMoveUp:
				humanPlayer.dy = -1
			case actionMoveLeft:
				humanPlayer.dx = -1
			case actionMoveRight:
				humanPlayer.dx = 1
			case actionMoveDown:
				humanPlayer.dy = 1
			}
		case <-gameTicker.C:
			var (
				minX, minY = 0, 1
				maxX, maxY = termbox.Size()
				screen     = make([]bool, maxX*(maxY+1))
				human      = &state.players[0]
			)
			for _, p := range state.players {
				screen[p.y*maxX+p.x] = true
			}
			for i, h := range state.hearts {
				if h == human.y*maxX+human.x {
					state.hearts[i] = -1
					human.health++
					state.healthChangeTransient = 5
					if human.health == state.maxHealth {
						termbox.Close()
						fmt.Println("You WIN! 😍")
						os.Exit(0)
					}
				}
			}
			for i := range state.players {
				var p = &state.players[i]
				if frameCount%p.speed == 0 {
					if i == 0 && (screen[(p.y+p.dy)*maxX+(p.x+p.dx)] || screen[p.y*maxX+(p.x+p.dx)] || screen[(p.y+p.dy)*maxX+p.x]) {
						p.health--
						state.healthChangeTransient = -5
						if p.health == 0 {
							termbox.Close()
							fmt.Println("You LOSE! 😭")
							os.Exit(0)
						}
					}
					if p.x+p.dx < minX || p.x+p.dx >= maxX || screen[p.y*maxX+(p.x+p.dx)] || screen[(p.y+p.dy)*maxX+(p.x+p.dx)] {
						p.dx *= -1
					}
					if p.y+p.dy < minY || p.y+p.dy >= maxY || screen[(p.y+p.dy)*maxX+p.x] || screen[(p.y+p.dy)*maxX+(p.x+p.dx)] {
						p.dy *= -1
					}
					p.x += p.dx
					p.y += p.dy
				}
			}
			frameCount++
		}
	}
}
