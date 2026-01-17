package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Entity struct {
	X, Y         float64
	Type         string
	image        *ebiten.Image
	Interactable bool
	Health       int32
	Alive        bool
}

func InitEntities(t *Terrain, numEnts int) []Entity {
	e := make([]Entity, numEnts)
	pts := RandomCoordGenerator(t, numEnts, true) // how many entities?
	re, _, _ := ebitenutil.NewImageFromFile("images/RockEntityTransparent.png")
	be, _, _ := ebitenutil.NewImageFromFile("images/BushEntityTransparent.png")
	for i := 0; i < numEnts; i++ {
		randomPick := rand.Intn(100)
		if randomPick >= 50 {
			e[i] = Entity{pts[i].X, pts[i].Y, "Rock", re, false, 80, true}
		} else {
			e[i] = Entity{pts[i].X, pts[i].Y, "Bush", be, false, 40, true}
		}
	}
	return e
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
