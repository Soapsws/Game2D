package main

import (
	"math"
)

func TilePosition(p *Player) Tile {
	// Tiles: 256 x 256
	x := p.X
	y := p.Y
	adjustedX := int(math.Floor((x + WorldWidth/2) / TileWorldSize))
	adjustedY := int(math.Floor((y + WorldHeight/2) / TileWorldSize))

	return Tiles[adjustedY][adjustedX]
}
