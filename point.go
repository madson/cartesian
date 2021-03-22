package main

type point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type pointDistance struct {
	Point    point `json:"point"`
	Distance int64 `json:"distance"`
}

func (p1 point) distanceFrom(p2 point) int64 {
	distanceX := p1.X - p2.X
	distanceY := p1.Y - p2.Y

	if distanceX < 0 {
		distanceX = -distanceX
	}

	if distanceY < 0 {
		distanceY = -distanceY
	}

	return distanceX + distanceY
}
