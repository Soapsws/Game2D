package main

const (
	WindowX         = 960
	WindowY         = 720
	PlayerScale     = 1.7
	PlayerWorldSize = TileSize * PlayerScale
	ZoneWorldSize   = 400

	ScreenWidth  = 720
	ScreenHeight = 540

	TileSize  = 32
	TileGrass = 0
	TileStone = 1
	TileWater = 2

	TileScale     = 8
	TileWorldSize = TileSize * TileScale // 256 px per tile
)

const (
	EntityScale     = 2
	EntityWorldSize = TileSize * EntityScale
)

const (
	MapWidthTiles  = 16
	MapHeightTiles = 16

	WorldWidth  = MapWidthTiles * TileSize * TileScale
	WorldHeight = MapHeightTiles * TileSize * TileScale

	HalfW = float64(ScreenWidth) / 2
	HalfH = float64(ScreenHeight) / 2
)
