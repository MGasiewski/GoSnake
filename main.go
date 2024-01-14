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
var score int

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && !(snake.d == DOWN) {
		snake.d = UP
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && !(snake.d == UP) {
		snake.d = DOWN
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && !(snake.d == RIGHT) {
		snake.d = LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && !(snake.d == LEFT) {
		snake.d = RIGHT
	}
	if !gameOver {
		ateFood := snake.UpdatePositionAndEatFood(gameBoard.foodX, gameBoard.foodY)
		if ateFood {
			gameBoard.GenerateFood()
			score += 1
		}
		gameOver = gameBoard.CheckCollision(&snake)
		if !gameOver {
			gameBoard.UpdateGameBoard(&snake)
		}
	}
	if gameOver && ebiten.IsKeyPressed((ebiten.KeySpace)) {
		initializeGame()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER. Press the spacebar to restart")
	} else {
		gameBoard.DrawGameboard(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetTPS(15)
	initializeGame()
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Snakey snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func initializeGame() {
	snake = Snake{
		d: RIGHT,
		SnakeHead: &Link{
			x: 32,
			y: 32,
		},
	}
	gameOver = false
	score = 0
	snake.size = 1
	gameBoard = GameBoard{}
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			gameBoard.Cells[i][j] = 0
		}
	}
	gameBoard.InitializeGameBoard(&snake)
}
