package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	X, Y         float64
	Type         string
	image        *ebiten.Image
	Interactable bool
	Health       int32
}

// Possible non-moving entities: rocks, trees/bushes
// Add moving entities separately

func (e *Entity) Interact(p *Player, isAttacking bool) {
	if isAttacking {
		// Damage interact
		e.Health -= p.Damage
		if e.IsDead() {
			fmt.Println("Entity dead")
		}
	} else {
		// Custom interact
		switch e.Type {
		case "Rock":

		case "Bush":
		}
	}
}

func (e *Entity) IsDead() bool {
	return e.Health <= 0
}
