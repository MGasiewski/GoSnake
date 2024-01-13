package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var snake Snake
var gameBoard GameBoard
var gameOver bool

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
	if !gameOver {
		ateFood := snake.UpdatePositionAndEatFood(gameBoard.foodX, gameBoard.foodY)
		if ateFood {
			gameBoard.GenerateFood()
		}
		gameOver = gameBoard.CheckCollision(&snake)
		if !gameOver {
			gameBoard.UpdateGameBoard(&snake)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER. Please Restart")
	} else {
		gameBoard.DrawGameboard(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetMaxTPS(15)
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
