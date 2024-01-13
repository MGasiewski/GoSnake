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
	foodX int
	foodY int
	snake *Snake
}

func (gb *GameBoard) DrawGameboard(screen *ebiten.Image) {
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			if gb.Cells[i][j] == 1 {
				vector.DrawFilledRect(screen, float32(i*10), float32(j*10), 10, 10, green, false)
			} else if gb.Cells[i][j] == 2 {
				vector.DrawFilledRect(screen, float32(i*10), float32(j*10), 10, 10, white, false)
			}
		}
	}
}

func (gb *GameBoard) UpdateGameBoard(s *Snake) {
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			gb.Cells[i][j] = 0
		}
	}
	segmentPtr := s.SnakeHead
	for segmentPtr != nil {
		gb.Cells[segmentPtr.x][segmentPtr.y] = SNAKE_SEGMENT
		segmentPtr = segmentPtr.l
	}
	gb.Cells[gb.foodX][gb.foodY] = FOOD
}

func (gb *GameBoard) CheckCollision(s *Snake) bool {
	snakeLink := s.SnakeHead.l
	for snakeLink != nil {
		if snakeLink.x == s.SnakeHead.x && snakeLink.y == s.SnakeHead.y {
			return true
		}
		snakeLink = snakeLink.l
	}
	if s.SnakeHead.x >= 64 || s.SnakeHead.x < 0 {
		return true
	}
	if s.SnakeHead.y >= 64 || s.SnakeHead.y < 0 {
		return true
	}
	return false
}

func (gb *GameBoard) InitializeGameBoard(s *Snake) {
	gb.Cells[s.SnakeHead.x][s.SnakeHead.y] = SNAKE_SEGMENT
	gameBoard.GenerateFood()
	gb.Cells[gb.foodX][gb.foodY] = FOOD
}

func (gb *GameBoard) GenerateFood() {
	gb.Cells[gb.foodX][gb.foodY] = 0
	gb.foodX = rand.Intn(64)
	gb.foodY = rand.Intn(64)
}
