package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Terrain struct {
	Grass *ebiten.Image
	Stone *ebiten.Image
	Water *ebiten.Image

	Tiles [][]Tile

	MapX float64
	MapY float64
}

type Tile struct {
	IDX int
	IDY int

	Type string
}

func InitTiles(M [][]int) [][]Tile {
	T := make([][]Tile, len(M))
	for i := 0; i < len(M); i++ {
		T[i] = make([]Tile, len(M[0]))
		for j := 0; j < len(M[0]); j++ {
			var TileType string
			switch M[i][j] {
			case 0:
				TileType = "Grass"
			case 1:
				TileType = "Stone"
			case 2:
				TileType = "Water"
			}

			NewTile := Tile{
				j,
				i,
				TileType,
			}
			T[i][j] = NewTile
		}
	}
	return T
}
