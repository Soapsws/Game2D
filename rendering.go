package main

import "github.com/hajimehoshi/ebiten/v2"

type Renderer struct {
	DisplayingZone      bool
	DisplayingCrafting  bool
	DisplayingInventory bool
}

func UpdateRenderer(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.R.DisplayingInventory = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.R.DisplayingZone = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.R.DisplayingCrafting = true
	}
}
