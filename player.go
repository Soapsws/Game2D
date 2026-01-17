package main

import (
	"math"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y    float64
	Health  int32
	Speed   float64
	Heading float64

	DisplayingZone bool
	MouseX         float64
	MouseY         float64

	Damage int32

	Inventory           map[string]int
	InventoryVisualizer []string
	Crafting            CraftingBook

	PlayerTimestamp time.Time

	image     *ebiten.Image
	zoneImage *ebiten.Image
}

func (p *Player) InventoryInit() {
	p.InventoryVisualizer = make([]string, 0)
	p.InventoryVisualizer = append(p.InventoryVisualizer, "Stone")
	p.InventoryVisualizer = append(p.InventoryVisualizer, "Stick")
}

func (p *Player) CraftingInit() {
	p.Crafting = CraftingBook{make([]Recipe, 0)}
	p.Crafting.InitCB()
}

func (p *Player) CheckCrafting(page int) string {
	s := ""
	for i := page - 1; i < page+3; i++ {
		if i >= len(p.Crafting.Recipes) {
			break
		}
		s += ">" + p.Crafting.Recipes[i].Name + ":\n"
		for name, cnt := range p.Crafting.Recipes[i].Materials {
			s += name + " | " + strconv.Itoa(cnt) + "\n"
		}
	}
	return s
}

func (p *Player) CheckInventory(page int) string {
	s := ""
	for i := page - 1; i < page+7; i++ {
		if i >= len(p.InventoryVisualizer) {
			break
		}
		s += p.InventoryVisualizer[i] + " " + strconv.Itoa(p.Inventory[p.InventoryVisualizer[i]]) + "\n"
	}
	return s
}

func (p *Player) TimeInit() {
	p.PlayerTimestamp = time.Now()
}

func (p *Player) InteractCooldown() bool {
	return time.Since(p.PlayerTimestamp) >= PlayerInteractCooldown
}

func (p *Player) Move(xDir, yDir float64, g *Game) error {
	var speed float64
	if TileRelativeToPosition(g.T, p.X, p.Y).Type == "Water" {
		speed = float64(p.Speed) * (0.7)
	} else {
		speed = p.Speed
	}

	p.X += speed * xDir
	p.Y += speed * yDir

	// new clamping so player never sees outside of the map

	p.X = math.Max(-1*WorldWidth/2, math.Min(WorldWidth/2-1, p.X))
	p.Y = math.Max(-1*WorldHeight/2, math.Min(p.Y, WorldHeight/2-1))

	if TileRelativeToPosition(g.T, p.X, p.Y).Type == "Stone" {
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
	// Using slice references to access pointer
	for i := range g.E {
		g.E[i].Interactable = false
		if DistanceCalculator(p.X, p.Y, g.E[i].X, g.E[i].Y, 200) {
			g.E[i].Interactable = true
			// fmt.Println("True")
		}
	}
}

func (p *Player) UpdateMousePos() {
	mouseX, mouseY := ebiten.CursorPosition()
	// Adjusting for coordinate space
	p.MouseX = float64(mouseX) - HalfW + p.X
	p.MouseY = float64(mouseY) - HalfH + p.Y
}

func (p *Player) ScanInteractable(g *Game) {
	interacted := false
	for i := range g.E {
		if g.E[i].Interactable && ContainsPointCircle(g.E[i].X, g.E[i].Y, EntityWorldSize/2, p.MouseX, p.MouseY) {
			if ebiten.IsKeyPressed(ebiten.KeyE) {
				g.E[i].Interact(p, true)
				interacted = true
				break
			} else if ebiten.IsKeyPressed(ebiten.KeyF) {
				g.E[i].Interact(p, false)
				interacted = true
				break
			}
		}
	}
	if interacted {
		now := time.Now()
		p.PlayerTimestamp = now
	}
}

func (p *Player) TouchingEntity(g *Game) (Entity, bool) {
	for _, e := range g.E {
		if !e.Alive {
			continue
		}
		b := CollisionDetectorCircle(p.X, p.Y, e.X, e.Y,
			float64(PlayerWorldSize)/2, EntityWorldSize/2)
		if b {
			return e, true
		}
	}
	return Entity{-1, -1, "", nil, false, -1, false}, false
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
