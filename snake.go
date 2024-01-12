package main

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type Snake struct {
	d Direction
	x int
	y int
}

func (s *Snake) EatFood() {

}

func (s *Snake) UpdatePosition() {
	switch s.d {
	case UP:
		s.y -= 1
	case DOWN:
		s.y += 1
	case LEFT:
		s.x -= 1
	case RIGHT:
		s.x += 1
	}

	if s.x >= 64 {
		s.x = 63
	} else if s.x < 0 {
		s.x = 0
	}

	if s.y >= 64 {
		s.y = 63
	} else if s.y < 0 {
		s.y = 0
	}
}
