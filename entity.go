package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X, Y  float64
	Type  string
	image *ebiten.Image
}

// Possible non-moving entities: rocks, trees/bushes
// Add moving entities separately
