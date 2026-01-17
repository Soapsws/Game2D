package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"errors"
	"math"
	"math/rand"
)

type Game struct {
	P *Player
	E []Entity
	T *Terrain
	R *Renderer
}

var WorldMap = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1, 0, 0, 2, 2, 1, 0, 0, 1, 0, 0, 0},
	{2, 2, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0},
	{2, 2, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{2, 2, 2, 2, 0, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 2, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 1, 2, 2, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0},
	{0, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 0, 0},
	{0, 0, 0, 2, 0, 0, 0, 0, 0, 1, 0, 0, 2, 2, 0, 0},
	{0, 0, 0, 2, 2, 2, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
}

// Universal Rendering Rule:
// Sprite-local transforms -> world placement -> camera transform

func (g *Game) DrawEntity(screen *ebiten.Image, ent Entity) {
	op := ebiten.DrawImageOptions{}
	bounds := float64(ent.image.Bounds().Dx())
	scale := float64(EntityWorldSize) / bounds
	op.GeoM.Scale(scale, scale)

	// Initialized in coordinate-space so translation isn't necessary

	// Stage 1
	op.GeoM.Translate(
		float64(ent.X),
		float64(ent.Y),
	)

	// Stage 2
	op.GeoM.Translate(
		-EntityWorldSize/2,
		-EntityWorldSize/2,
	)

	// Stage 3
	op.GeoM.Translate(
		-1*(g.P.X)+(HalfW),
		-1*(g.P.Y)+(HalfH),
	)

	screen.DrawImage(ent.image, &op)
}

func (g *Game) DrawTile(screen *ebiten.Image, tile *ebiten.Image, tileX, tileY int) {
	op := ebiten.DrawImageOptions{}
	bounds := float64(tile.Bounds().Dx())
	scale := float64(TileWorldSize) / bounds
	op.GeoM.Scale(scale, scale)

	// Initialized in tile-space so have to be manually translated
	// See terrain.go for tile-space info

	// 256 =

	// Stage 1
	op.GeoM.Translate(
		float64(tileX*TileSize),
		float64(tileY*TileSize),
	)

	// Stage 2
	op.GeoM.Translate(
		-1*float64(len(WorldMap))/2*TileWorldSize,
		-1*float64(len(WorldMap))/2*TileWorldSize,
	)

	// Stage 3
	op.GeoM.Translate(
		-1*(g.P.X)+(HalfW),
		-1*(g.P.Y)+(HalfH),
	)

	screen.DrawImage(tile, &op)
}

func (g *Game) DrawZone(screen *ebiten.Image, zone *ebiten.Image) {
	oz := ebiten.DrawImageOptions{}
	boundsz := float64(zone.Bounds().Dx())
	scalez := float64(ZoneWorldSize / boundsz)
	wz := zone.Bounds().Dx()
	hz := zone.Bounds().Dy()

	oz.GeoM.Translate(-float64(wz)/2, -float64(hz)/2)
	oz.GeoM.Scale(scalez, scalez)

	oz.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	screen.DrawImage(g.P.zoneImage, &oz)
}

func (g *Game) DrawPlayer(screen *ebiten.Image, player *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	bounds := float64(player.Bounds().Dx())
	scale := float64(PlayerWorldSize) / bounds
	w := player.Bounds().Dx()
	h := player.Bounds().Dy()

	// Place-specific pipeline: centering before scaling
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2) // Centers the sprite - super important!
	op.GeoM.Scale(scale, scale)

	op.GeoM.Rotate(g.P.Heading)
	op.GeoM.Translate(float64(ScreenWidth)/2, float64(ScreenHeight)/2)

	screen.DrawImage(g.P.image, &op)
}

func (g *Game) PlayerMovement() {
	var dx float64 = 0
	var dy float64 = 0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dy = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		dy = 1
	}
	// DOWN RIGHT POSITIVE
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dx = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		dx = 1
	}

	if dx != 0 && dy != 0 {
		dx /= math.Sqrt(2)
		dy /= math.Sqrt(2)
	}

	g.P.Move(dx, 0, g)
	g.P.Move(0, dy, g)
}

func (g *Game) PlayerHeading() error {
	mouseX, mouseY := ebiten.CursorPosition()
	playerScreenX := float64(ScreenWidth) / 2
	playerScreenY := float64(ScreenHeight) / 2

	dx := float64(mouseX) - playerScreenX
	dy := float64(mouseY) - playerScreenY
	angle := math.Atan2(dy, dx)

	g.P.Rotate(angle + math.Pi/2) // upward facing default - offset
	return nil
}

func (g *Game) Update() error {
	g.PlayerMovement()
	g.PlayerHeading()
	g.P.UpdateMousePos()

	UpdateRenderer(g)

	g.P.CheckInteractableZone(g)
	if g.P.InteractCooldown() {
		g.P.ScanInteractable(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Terrain

	for y := 0; y < len(WorldMap); y++ {
		for x := 0; x < len(WorldMap[y]); x++ {
			TileID := WorldMap[y][x]
			switch TileID {
			case 0:
				g.DrawTile(screen, g.T.Grass, TileScale*x, TileScale*y)
			case 1:
				g.DrawTile(screen, g.T.Stone, TileScale*x, TileScale*y)
			case 2:
				g.DrawTile(screen, g.T.Water, TileScale*x, TileScale*y)
			}
		}
	}

	// Entities
	for i := 0; i < len(g.E); i++ {
		if g.E[i].IsDead(g.P) {
			continue
		}
		g.DrawEntity(screen, g.E[i])
	}

	// Player + Interactable Zone
	g.DrawPlayer(screen, g.P.image)

	debug := fmt.Sprintf(
		"Player X: %.1f | Player Y: %.1f \nTile[%d,%d]\nMouse X: %.1f | Mouse Y: %.1f",
		g.P.X,
		g.P.Y,
		TilePosition(g.P).IDX,
		TilePosition(g.P).IDY,
		g.P.MouseX,
		g.P.MouseY,
	)
	ebitenutil.DebugPrintAt(screen, debug, 10, 10)

	if g.R.DisplayingZone {
		g.DrawZone(screen, g.P.zoneImage)
		g.R.DisplayingZone = false
	} else if g.R.DisplayingInventory {
		s := g.P.CheckInventory(1)
		ebitenutil.DebugPrintAt(screen, s, ScreenWidth/2, ScreenHeight/2)
		g.R.DisplayingInventory = false
	} else if g.R.DisplayingCrafting {
		s := g.P.CheckCrafting(1)
		ebitenutil.DebugPrintAt(screen, s, ScreenWidth/2, ScreenHeight/2)
		g.R.DisplayingCrafting = false
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func Init() (*Player, *[]Entity, *Terrain, *Renderer, error) {
	img, _, err := ebitenutil.NewImageFromFile("images/Char_Game.png")
	zoneImg, _, err := ebitenutil.NewImageFromFile("images/InteractableZone.png")

	if err != nil {
		return nil, nil, nil, nil, errors.New("Bad image")
	}

	p := Player{
		X:              0,
		Y:              0,
		Health:         100,
		Speed:          4,
		Heading:        0,
		DisplayingZone: false,
		MouseX:         0,
		MouseY:         0,
		Damage:         10,
		Inventory:      make(map[string]int),
		image:          img,
		zoneImage:      zoneImg,
	}

	p.TimeInit()
	p.InventoryInit()
	p.CraftingInit()

	NumEnts := 100
	e := make([]Entity, NumEnts)
	pts := RandomCoordGenerator(NumEnts) // how many entities?
	re, _, err := ebitenutil.NewImageFromFile("images/RockEntityTransparent.png")
	be, _, err := ebitenutil.NewImageFromFile("images/BushEntityTransparent.png")
	for i := 0; i < NumEnts; i++ {
		randomPick := rand.Intn(100)
		if randomPick >= 50 {
			e[i] = Entity{pts[i].X, pts[i].Y, "Rock", re, false, 80, true}
		} else {
			e[i] = Entity{pts[i].X, pts[i].Y, "Bush", be, false, 40, true}
		}
	}

	gt, _, err := ebitenutil.NewImageFromFile("images/GrassTile.png")
	st, _, err := ebitenutil.NewImageFromFile("images/StoneTile.png")
	wt, _, err := ebitenutil.NewImageFromFile("images/WaterTile.png")

	t := Terrain{
		Grass: gt,
		Stone: st,
		Water: wt,

		MapX: 0,
		MapY: 0,
	}

	r := Renderer{
		DisplayingZone:      false,
		DisplayingCrafting:  false,
		DisplayingInventory: false,
	}

	return &p, &e, &t, &r, nil
}

func main() {
	ebiten.SetWindowSize(WindowX, WindowY)
	ebiten.SetWindowTitle("2D Game")
	player, entities, terrain, renderer, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player,
		*entities,
		terrain,
		renderer,
	}

	Tiles = InitTiles(WorldMap)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
