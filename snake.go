package main

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type Snake struct {
	d         Direction
	SnakeHead *Link
	size      int
}

type Link struct {
	x int
	y int
	l *Link
}

func (s *Snake) EatFood() {

}

func (s *Snake) UpdatePosition() {
	//move snake
	switch s.d {
	case UP:
		s.SnakeHead.y -= 1
	case DOWN:
		s.SnakeHead.y += 1
	case LEFT:
		s.SnakeHead.x -= 1
	case RIGHT:
		s.SnakeHead.x += 1
	}

	//adjust if out of bounds
	if s.SnakeHead.x >= 64 {
		s.SnakeHead.x = 63
	} else if s.SnakeHead.x < 0 {
		s.SnakeHead.x = 0
	}
	if s.SnakeHead.y >= 64 {
		s.SnakeHead.y = 63
	} else if s.SnakeHead.y < 0 {
		s.SnakeHead.y = 0
	}
}
