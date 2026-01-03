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
	i := 0
	for i < n {
		x := rand.Intn(4096) - 2048
		y := rand.Intn(4096) - 2048
		bc := 0
		if x < 100 && x > -100 {
			bc += 1
		}
		if y < 100 && y > -100 {
			bc += 1
		}
		if bc < 2 {
			coords[i] = Point{float64(x), float64(y)}
			i++
		}
	}
	return coords
}

// for circles
func CollisionDetectorCircle(c1x, c1y, c2x, c2y, r1, r2 float64) bool {
	// takes two centers and radii
	dx := c1x - c2x
	dy := c1y - c2y
	return dx*dx+dy*dy <= (r1+r2)*(r1+r2)
}

// Circles vs Boxes
func CollisionDetectorHybrid(cx, cy, r, bx, by, hw, hh float64) bool {
	closestX := math.Max(bx-hw, math.Min(cx, bx+hw))
	closestY := math.Max(by-hh, math.Min(cy, by+hh))

	dx := cx - closestX
	dy := cy - closestY

	return dx*dx+dy*dy <= r*r
}

func DistanceCalculator(x1, y1, x2, y2, threshold float64) bool {
	return math.Sqrt(((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))) <= threshold
}
