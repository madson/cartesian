package main

import "math/big"

type point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type pointDistance struct {
	Point    point  `json:"point"`
	Distance string `json:"distance"`
}

func (p1 point) distanceFrom(p2 point) *big.Int {
	distanceX := big.NewInt(0)
	distanceY := big.NewInt(0)
	result := big.NewInt(0)

	x1 := big.NewInt(0).SetInt64(p1.X)
	x2 := big.NewInt(0).SetInt64(p2.X)

	y1 := big.NewInt(0).SetInt64(p1.Y)
	y2 := big.NewInt(0).SetInt64(p2.Y)

	distanceX.Sub(x1, x2)
	distanceY.Sub(y1, y2)

	distanceX.Abs(distanceX)
	distanceY.Abs(distanceY)

	result.Add(distanceX, distanceY)

	return result
}
