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

func (p *Player) Move(xDir, yDir float64, g *Game) error {
	var speed float64
	if TilePosition(p).Type == "Water" {
		speed = float64(p.Speed) * (0.7)
	} else {
		speed = p.Speed
	}

	p.X += speed * xDir
	p.Y += speed * yDir

	// new clamping so player never sees outside of the map

	p.X = math.Max(-1*WorldWidth/2, math.Min(WorldWidth/2-1, p.X))
	p.Y = math.Max(-1*WorldHeight/2, math.Min(p.Y, WorldHeight/2-1))

	if TilePosition(p).Type == "Stone" {
		p.X -= speed * xDir
		p.Y -= speed * yDir
	} else if p.TouchingEntity(g) {
		p.X -= speed * xDir
		p.Y -= speed * yDir
	}

	return nil
}

func (p *Player) TouchingEntity(g *Game) bool {
	for _, e := range g.E {
		b := CollisionDetectorCircle(p.X, p.Y, e.X, e.Y,
			float64(PlayerWorldSize)/2, EntityWorldSize/2)
		if b {
			return true
		}
	}
	return false
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
