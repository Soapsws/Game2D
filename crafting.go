package main

type Recipe struct {
	Name      string
	Materials map[string]int
}

type CraftingBook struct {
	Recipes []Recipe
}

func (c *CraftingBook) Craft(recipe string, p *Player) bool {
	for i := range c.Recipes {
		if c.Recipes[i].Name == recipe {
			if !c.IsCraftable(recipe, p) {
				return false
			}
			for item, cnt := range c.Recipes[i].Materials {
				p.Inventory[item] -= cnt
			}
			p.Inventory[recipe]++
			return true
		}
	}
	return false
}

func (c *CraftingBook) IsCraftable(recipe string, p *Player) bool {
	for i := range c.Recipes {
		if c.Recipes[i].Name == recipe {
			for item, cnt := range c.Recipes[i].Materials {
				if p.Inventory[item] < cnt {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (c *CraftingBook) InitCB() {
	// Objects

	WoodBlock := Recipe{
		Name: "WoodBlock",
		Materials: map[string]int{
			"Stick": 3,
		},
	}

	StoneBlock := Recipe{
		Name: "StoneBlock",
		Materials: map[string]int{
			"Stone": 3,
		},
	}

	Wall := Recipe{
		Name: "Wall",
		Materials: map[string]int{
			"Stick": 5,
		},
	}

	StoneWall := Recipe{
		Name: "StoneWall",
		Materials: map[string]int{
			"Stick": 5,
			"Stone": 5,
		},
	}

	// Tools/Weapons

	StoneAxe := Recipe{
		Name: "StoneAxe",
		Materials: map[string]int{
			"Stick": 3,
			"Stone": 1,
		},
	}

	StonePickaxe := Recipe{
		Name: "StonePickaxe",
		Materials: map[string]int{
			"Stick": 2,
			"Stone": 2,
		},
	}

	StoneSword := Recipe{
		Name: "StoneSword",
		Materials: map[string]int{
			"Stick": 1,
			"Stone": 3,
		},
	}

	// Initializing recipe book

	c.Recipes = make([]Recipe, 0)
	c.Recipes = append(c.Recipes, WoodBlock)
	c.Recipes = append(c.Recipes, StoneBlock)
	c.Recipes = append(c.Recipes, Wall)
	c.Recipes = append(c.Recipes, StoneWall)
	c.Recipes = append(c.Recipes, StoneAxe)
	c.Recipes = append(c.Recipes, StonePickaxe)
	c.Recipes = append(c.Recipes, StoneSword)
}
