package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var snake Snake
var gameBoard GameBoard

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		snake.d = UP
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		snake.d = DOWN
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		snake.d = LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		snake.d = RIGHT
	}
	snake.UpdatePosition()
	gameBoard.UpdateGameBoard(&snake)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameBoard.DrawGameboard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	snake = Snake{
		d: RIGHT,
		SnakeHead: &Link{
			x: 32,
			y: 32,
		},
	}
	snake.size = 1
	gameBoard = GameBoard{}
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			gameBoard.Cells[i][j] = 0
		}
	}
	gameBoard.InitializeGameBoard(&snake)
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Snakey snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
