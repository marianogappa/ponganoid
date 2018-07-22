package main

import (
	"fmt"
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
)

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
				maxX, maxY = termbox.Size() // resizes cause Size() to change
			)

			// Heart management
			var human = &state.players[0]
			for i, h := range state.hearts { // map would have been more performant, but would panic (goroutines)
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

			// Collision management
			var screen = make([]bool, maxX*(maxY+1)) // player position cache makes collision checks constant time
			for _, p := range state.players {        // map would have been more performant, but would panic (goroutines)
				screen[p.y*maxX+p.x] = true
			}
			for i := range state.players {
				var p = &state.players[i]
				if frameCount%p.speed == 0 { // individual speed enables each player to move every "speed" frames
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
