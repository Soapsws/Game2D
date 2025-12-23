package main

import (
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

	// p.X = math.Max(HalfW, math.Min(p.X, float64(WorldWidth)-HalfW))
	// p.Y = math.Max(HalfH, math.Min(p.Y, float64(WorldHeight)-HalfH))

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
