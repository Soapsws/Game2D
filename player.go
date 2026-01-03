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

	DisplayingZone bool

	image     *ebiten.Image
	zoneImage *ebiten.Image
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
	} else {
		// Circle collision smooth physics (copied)
		if ent, hit := p.TouchingEntity(g); hit {
			dx := p.X - ent.X
			dy := p.Y - ent.Y

			dist := math.Sqrt(dx*dx + dy*dy)
			if dist == 0 {
				return nil
			}

			overlap := (PlayerWorldSize/2 + EntityWorldSize/2) - dist
			if overlap > 0 {
				nx := dx / dist
				ny := dy / dist

				p.X += nx * overlap
				p.Y += ny * overlap
			}
		}
	}

	return nil
}

func (p *Player) CheckInteractableZone(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		p.DisplayingZone = true
	} else {
		p.DisplayingZone = false
	}

	for _, e := range g.E {
		if DistanceCalculator(p.X, e.X, p.Y, e.Y, 200) {
			e.Interactable = true
		}
	}
}

func (p *Player) ScanInteractable(g *Game) {
	// mouseX, mouseY := ebiten.CursorPosition()
}

func (p *Player) TouchingEntity(g *Game) (Entity, bool) {
	for _, e := range g.E {
		b := CollisionDetectorCircle(p.X, p.Y, e.X, e.Y,
			float64(PlayerWorldSize)/2, EntityWorldSize/2)
		if b {
			return e, true
		}
	}
	return Entity{-1, -1, "", nil, false}, false
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
