package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"errors"
	"math"
)

type Game struct {
	P *Player

	T *Terrain
}

var WorldMap = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1, 0, 0, 2, 2, 1, 1, 0, 1, 0, 0, 0},
	{2, 2, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0},
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

func (g *Game) DrawTile(screen *ebiten.Image, tile *ebiten.Image, tileX, tileY int) {
	op := ebiten.DrawImageOptions{}
	scale := float64(TileSize) / float64(tile.Bounds().Dx())
	op.GeoM.Scale(scale, scale)
	op.GeoM.Scale(TileScale, TileScale) // Rescale
	op.GeoM.Translate(
		float64(tileX*TileSize),
		float64(tileY*TileSize),
	)
	op.GeoM.Translate(
		-1*float64(len(WorldMap))/2*(TileScale*TileSize),
		-1*float64(len(WorldMap))/2*(TileScale*TileSize),
	)

	// Map-based movement
	op.GeoM.Translate(
		-1*(g.P.X)+(HalfW),
		-1*(g.P.Y)+(HalfH),
	)

	screen.DrawImage(tile, &op)
}

func (g *Game) DrawPlayer(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}   // Essentially creating an image transformation pipeline
	size := g.P.image.Bounds().Size() // New non-deprecated version
	w := size.X
	h := size.Y

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2) // Centers the sprite - super important!
	op.GeoM.Scale(PlayerScale, PlayerScale)         // Size img down - NOTE ORDER OF TRANSFORMS
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

	g.P.Move(dx, dy)
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

	// Player
	g.DrawPlayer(screen)

	debug := fmt.Sprintf(
		"Player X: %.1f\nPlayer Y: %.1f",
		g.P.X,
		g.P.Y,
	)
	ebitenutil.DebugPrintAt(screen, debug, 10, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func Init() (*Player, *Terrain, error) {
	img, _, err := ebitenutil.NewImageFromFile("Char_Game.png")

	if err != nil {
		return nil, nil, errors.New("Bad image")
	}

	p := Player{
		X:       0,
		Y:       0,
		Health:  100,
		Speed:   2,
		Heading: 0,
		image:   img,
	}

	gt, _, err := ebitenutil.NewImageFromFile("GrassTile.png")
	st, _, err := ebitenutil.NewImageFromFile("StoneTile.png")
	wt, _, err := ebitenutil.NewImageFromFile("WaterTile.png")

	t := Terrain{
		Grass: gt,
		Stone: st,
		Water: wt,

		MapX: 0,
		MapY: 0,
	}

	return &p, &t, nil
}

func main() {
	ebiten.SetWindowSize(WindowX, WindowY)
	ebiten.SetWindowTitle("2D Game")
	player, terrain, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player,
		terrain,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
