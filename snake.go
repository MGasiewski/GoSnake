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

func (s *Snake) UpdatePositionAndEatFood(foodX, foodY int) bool {
	newLink := Link{
		x: s.SnakeHead.x,
		y: s.SnakeHead.y,
	}
	switch s.d {
	case UP:
		newLink.y -= 1
	case DOWN:
		newLink.y += 1
	case LEFT:
		newLink.x -= 1
	case RIGHT:
		newLink.x += 1
	}
	prevHead := s.SnakeHead
	s.SnakeHead = &newLink
	s.SnakeHead.l = prevHead
	if s.SnakeHead.x == foodX && s.SnakeHead.y == foodY {
		return true
	} else {
		snakePtrPtr := &s.SnakeHead.l
		for *snakePtrPtr != nil {
			if (*snakePtrPtr).l != nil {
				snakePtrPtr = &(*snakePtrPtr).l
			} else {
				*snakePtrPtr = nil
			}
		}
		return false
	}
}
