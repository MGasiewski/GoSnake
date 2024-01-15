package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var snake Snake
var gameBoard GameBoard
var leaderBoard LeaderBoard
var gameOver bool
var score int
var playerName string
var nameEntered bool

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
			score += 100
		}
		gameOver = gameBoard.CheckCollision(&snake)
		if !gameOver {
			gameBoard.UpdateGameBoard(&snake)
		}
	}
	if gameOver && !nameEntered {
		AcceptInput()
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		initializeGame()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if gameOver && !nameEntered {
		ebitenutil.DebugPrint(screen, "GAME OVER. Press the spacebar to restart\n"+
			"Final Score: "+strconv.Itoa(score)+"\n"+
			"Enter your Name and press Enter: ")
		ebitenutil.DebugPrintAt(screen, playerName, 200, 33)
	} else if nameEntered {
		leaderBoard.DrawLeaderBoard(screen)
	} else {
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(score))
		gameBoard.DrawGameboard(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func AcceptInput() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		leaderBoard.AddToLeaderBoard(Score{
			Name:  playerName,
			Value: score,
		})
		playerName = ""
		score = 0
		nameEntered = true
	} else {
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(playerName) > 0 {
			playerName = playerName[:len(playerName)-1]
		}
		runeSlice := ebiten.AppendInputChars([]rune(""))
		playerName += string(runeSlice)
	}
}

func main() {
	ebiten.SetTPS(15)
	initializeGame()
	leaderBoard = LeaderBoard{}
	leaderBoard.Number = 0
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
	nameEntered = false
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
