package main

import (
	"math"
	"math/rand"
)

func TilePosition(p *Player) Tile {
	// Tiles: 256 x 256
	x := p.X
	y := p.Y
	adjustedX := int(math.Floor((x + WorldWidth/2) / TileWorldSize))
	adjustedY := int(math.Floor((y + WorldHeight/2) / TileWorldSize))

	return Tiles[adjustedY][adjustedX]
}

type Point struct {
	X float64
	Y float64
}

func RandomCoordGenerator(n int) []Point {
	coords := make([]Point, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(4096)
		y := rand.Intn(4096)
		coords[i] = Point{float64(x), float64(y)}
	}
	return coords
}
