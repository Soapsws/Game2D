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
		// s := fmt.Sprintf("Entity Coord: %d, %d", x-2048, y-2048)
		// fmt.Println(s)
		coords[i] = Point{float64(x) - 2048, float64(y) - 2048}
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
