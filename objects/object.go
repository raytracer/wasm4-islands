package objects

import (
	"cart/island"
	"cart/random"
	"cart/sprites"
	"cart/w4"
)

// GameObjects is a list of game objects
var GameObjects = make([]GameObject, 1000)

// ShipInfos is a list of ship infos (max 10 ships)
var ShipInfos = make([]ShipInfo, 10)

const (
	// GameObjectKindTree is a tree
	GameObjectKindTree = iota + 1
	// GameObjectKindChurch is a church
	GameObjectKindChurch
	// GameObjectKindLumberjack is a lumberjack
	GameObjectKindLumberjack
)

// GameObject is a game object with a position and a kind
type GameObject struct {
	X    byte
	Y    byte
	Kind byte
}

// ShipInfo saves the offset of the ship + cargo and destination
type ShipInfo struct {
	Ref          int
	OffsetX      float64
	OffsetY      float64
	DestinationX byte
	DestinationY byte
	State        byte
	Cargo        byte
	Amount       byte
}

// GetObjectSize gets the width and height of the object
func (g *GameObject) GetObjectSize() (byte, byte) {
	switch g.Kind {
	case GameObjectKindTree:
		return 1, 1
	case GameObjectKindChurch:
		return 4, 3
	case GameObjectKindLumberjack:
		return 4, 4
	default:
		panic("unknown object kind")
	}
}

// GetObjectCosts gets the costs of the object (gold, wood, stone)
func (g *GameObject) GetObjectCosts() (int, byte, byte) {
	switch g.Kind {
	case GameObjectKindTree:
		return 10, 0, 0
	case GameObjectKindChurch:
		return 100, 4, 0
	case GameObjectKindLumberjack:
		return 50, 2, 0
	default:
		panic("unknown object kind")
	}
}

// InsertDrawables inserts the drawables of the game object
func (g *GameObject) InsertDrawables(index int16) {
	width, height := g.GetObjectSize()
	for x := g.X - width + 1; x < g.X; x++ {
		sprites.InsertDrawable(sprites.Drawable{X: byte(x), Y: byte(g.Y), Ref: index})
	}
	for y := g.Y - height + 1; y < g.Y; y++ {
		sprites.InsertDrawable(sprites.Drawable{X: byte(g.X), Y: byte(y), Ref: index})
	}
	sprites.InsertDrawable(sprites.Drawable{X: byte(g.X), Y: byte(g.Y), Ref: index})
}

// DeleteDrawables deletes the drawables of the game object
func (g *GameObject) DeleteDrawables(index int16) {
	sprites.DeleteDrawables(index)
}

// CanBeBuilt checks if the object can be built
func (g *GameObject) CanBeBuilt() bool {
	isl := island.FindIsland(int(g.X), int(g.Y))

	if isl == nil {
		return false
	}

	gold, wood, stone := g.GetObjectCosts()

	if island.Gold < gold || isl.Wood < wood || isl.Stone < stone {
		return false
	}

	w, h := g.GetObjectSize()

	if !island.CheckLand(int(g.X), int(g.Y), int(w), int(h)) {
		return false
	}

	thisWidth, thisHeight := g.GetObjectSize()
	for _, obj := range GameObjects {
		if obj.Kind != 0 && obj.Kind != GameObjectKindTree {
			otherWidth, otherHeight := obj.GetObjectSize()
			if obj.X >= (g.X-thisWidth+1) && obj.X-otherWidth+1 <= g.X && obj.Y >= (g.Y-thisHeight+1) && obj.Y-otherHeight+1 <= g.Y {
				return false
			}
		}
	}

	return true
}

// RemoveTrees removes trees that are in the way of the object
func (g *GameObject) RemoveTrees() bool {
	thisWidth, thisHeight := g.GetObjectSize()
	for i, obj := range GameObjects {
		if obj.Kind == GameObjectKindTree {
			otherWidth, otherHeight := obj.GetObjectSize()
			if obj.X >= (g.X-thisWidth+1) && obj.X-otherWidth+1 <= g.X && obj.Y >= (g.Y-thisHeight+1) && obj.Y-otherHeight+1 <= g.Y {
				GameObjects[i].DeleteDrawables(int16(i))
				GameObjects[i] = GameObject{}
			}
		}
	}

	return true
}

// Simulate simulates the object for the current tick count
func (g *GameObject) Simulate(tick int64) {
	switch g.Kind {
	case GameObjectKindChurch:
		if tick%(60*30) == 0 {
			island.Gold -= 5
		}
	}
}

// Place places the object in the array
func (g *GameObject) Place() {
	g.RemoveTrees()
	for i, obj := range GameObjects {
		if obj.Kind == 0 {
			GameObjects[i] = *g
			g.InsertDrawables(int16(i))
			sprites.Drawables = quickSortStart(sprites.Drawables)
			gold, wood, stone := g.GetObjectCosts()
			island.Gold -= gold
			isl := island.FindIsland(int(g.X), int(g.Y))

			if isl != nil {
				isl.Wood -= wood
				isl.Stone -= stone
			}

			w4.Tone(25, 5|(10<<8), 25, w4.TONE_TRIANGLE)
			break
		}
	}
}

func val(gameObject sprites.Drawable) int {
	return int(gameObject.X) + int(gameObject.Y)
}

func partition(arr []sprites.Drawable, low, high int) ([]sprites.Drawable, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if val(arr[j]) < val(pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []sprites.Drawable, low, high int) []sprites.Drawable {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []sprites.Drawable) []sprites.Drawable {
	return quickSort(arr, 0, len(arr)-1)
}

// DrawGameObjects draws the game objects via the drawables
func DrawGameObjects() {
	for _, drawable := range sprites.Drawables {
		if drawable.Ref != -1 {
			gameObject := GameObjects[drawable.Ref]

			switch gameObject.Kind {
			case GameObjectKindTree:
				sprites.DrawTreeTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindChurch:
				sprites.DrawChurchTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindLumberjack:
				sprites.DrawLumberjackTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			}
		}
	}
}

// GenerateGameObjects generates game objects
func GenerateGameObjects() {
	counter := 0
	for x := 0; x < island.Size; x++ {
		for y := 0; y < island.Size; y++ {
			object := GameObject{X: byte(x), Y: byte(y), Kind: GameObjectKindTree}
			w, h := object.GetObjectSize()
			if island.CheckLand(x, y, int(w), int(h)) {
				if random.R.Float64() < 0.04 && counter < len(GameObjects) {
					GameObjects[counter] = object
					GameObjects[counter].InsertDrawables(int16(counter))
					counter++
				}
			}
		}
	}
}

// SimulateObjects simulates all objects
func SimulateObjects(tick int64) {
	for i := range GameObjects {
		if GameObjects[i].Kind != 0 {
			GameObjects[i].Simulate(tick)
		}
	}
}

// ClearShipInfos clears the ship infos to reference no ships
func ClearShipInfos() {
	for i := range ShipInfos {
		ShipInfos[i] = ShipInfo{Ref: -1}
	}
}

// GetObjectAt gets the object at the given position
func GetObjectAt(x, y int) *GameObject {
	for i := range GameObjects {
		if GameObjects[i].Kind != 0 {
			width, height := GameObjects[i].GetObjectSize()
			if x >= int(GameObjects[i].X)-int(width)+1 && x <= int(GameObjects[i].X) && y >= int(GameObjects[i].Y)-int(height)+1 && y <= int(GameObjects[i].Y) {
				return &GameObjects[i]
			}
		}
	}
	return nil
}

// ClearGameObjects clears the game objects
func ClearGameObjects() {
	for i := range GameObjects {
		GameObjects[i] = GameObject{}
	}
}
