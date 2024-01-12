package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SNAKE_SEGMENT = 1
	FOOD          = 2
)

var black = color.RGBA{0, 0, 0, 255}
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

type GameBoard struct {
	Cells [64][64]int
}

func (gb *GameBoard) DrawGameboard(screen *ebiten.Image) {
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			if gameBoard.Cells[i][j] == 1 {
				vector.DrawFilledRect(screen, float32(i*10), float32(j*10), 10, 10, green, false)
			} else if gameBoard.Cells[i][j] == 2 {
				vector.DrawFilledRect(screen, float32(i*10), float32(j*10), 10, 10, white, false)
			}
		}
	}
}

func (gb *GameBoard) UpdateGameBoard(s *Snake) {
	gameBoard.Cells[s.x][s.y] = SNAKE_SEGMENT
}

func (gb *GameBoard) InitializeGameBoard(s *Snake) {
	gameBoard.Cells[s.x][s.y] = SNAKE_SEGMENT
	foodX := rand.Intn(64)
	foodY := rand.Intn(64)
	gameBoard.Cells[foodX][foodY] = FOOD
}
