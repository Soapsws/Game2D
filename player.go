package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y    float64
	Health  int32
	Speed   float64
	Heading float64

	image *ebiten.Image
}

func (p *Player) Move(xDir, yDir float64) error {
	p.X += p.Speed * xDir
	p.Y += p.Speed * yDir

	// new clamping so player never sees outside of the mAAp

	p.X = math.Max(-1*WorldWidth/2, math.Min(WorldWidth/2, p.X))
	p.Y = math.Max(-1*WorldHeight/2, math.Min(p.Y, WorldHeight/2))

	return nil
}

func (p *Player) Rotate(angle float64) {
	p.Heading = angle
}

func (p *Player) TakeDamage(dmg int32) {
	p.Health -= dmg
}

func (p *Player) IsDead() bool {
	return (p.Health <= 0)
}
