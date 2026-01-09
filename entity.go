package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	X, Y         float64
	Type         string
	image        *ebiten.Image
	Interactable bool
	Health       int32
	Alive        bool
}

// Possible non-moving entities: rocks, trees/bushes
// Add moving entities separately

func (e *Entity) Interact(p *Player, isAttacking bool) {
	if isAttacking {
		// Damage interact
		e.Health -= p.Damage
	} else {
		// Custom interact
		switch e.Type {
		case "Rock":

		case "Bush":
		}
	}
}

func (e *Entity) DropLoot(p *Player) {
	switch e.Type {
	case "Rock":
		// stone := Item{"Stone"}
		p.Inventory["Stone"] += 1
	case "Bush":
		// stick := Item{"Stick"}
		p.Inventory["Stick"] += 1
	}
}

func (e *Entity) IsDead(p *Player) bool {
	if e.Health <= 0 && e.Alive {
		e.DropLoot(p)
		e.Alive = false
	}
	return e.Health <= 0
}
