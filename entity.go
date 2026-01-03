package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X, Y         float64
	Type         string
	image        *ebiten.Image
	Interactable bool
}

// Possible non-moving entities: rocks, trees/bushes
// Add moving entities separately

func (e *Entity) Interact() {
	if e.Interactable == false {
		return
	}

}
