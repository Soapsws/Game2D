package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Terrain struct {
	Grass *ebiten.Image
	Stone *ebiten.Image
	Water *ebiten.Image

	MapX float64
	MapY float64
}
