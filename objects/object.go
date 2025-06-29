package objects

import (
	"cart/island"
	"cart/random"
	"cart/sprites"
)

// GameObjects is a list of game objects
var GameObjects = make([]GameObject, 1000)

// ShipInfos is a list of ship infos (max 10 ships)
var ShipInfos = make([]ShipInfo, 10)

const (
	// Destination North
	NorthEast = iota
	NorthWest
	SouthEast
	SouthWest
)

const (
	// GameObjectKindTree is a tree
	GameObjectKindTree = iota + 1
	// GameObjectKindTree is a tree
	GameObjectKindMountain
	// GameObjectKindChurch is a church
	GameObjectKindChurch
	// GameObjectKindLumberjack is a lumberjack
	GameObjectKindLumberjack
	// GameObjectKindPioneer is a pioneer house
	GameObjectKindPioneer
	// GameObjectKindShip is a ship
	GameObjectKindShip
	// GameObjectKindWarehouse is a warehouse
	GameObjectKindWarehouse
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
	Direction    byte
	Cargo        byte
	Amount       byte
}

// GetRadius gets the radius of the object in tiles
func (g *GameObject) GetRadius() int {
	switch g.Kind {
	case GameObjectKindChurch:
		return 5
	case GameObjectKindLumberjack:
		return 3
	default:
		return 0
	}
}

// GetObjectSize gets the width and height of the object
func (g *GameObject) GetObjectSize() (byte, byte) {
	switch g.Kind {
	case GameObjectKindTree:
		return 1, 1
	case GameObjectKindMountain:
		return 4, 4
	case GameObjectKindChurch:
		return 4, 3
	case GameObjectKindLumberjack:
		return 4, 4
	case GameObjectKindPioneer:
		return 4, 4
	case GameObjectKindShip:
		return 4, 4
	case GameObjectKindWarehouse:
		return 3, 5
	default:
		panic("unknown object kind")
	}
}

// GetObjectCosts gets the costs of the object (gold, wood, stone)
func (g *GameObject) GetObjectMaintenance() int {
	switch g.Kind {
	case GameObjectKindChurch:
		return 5
	case GameObjectKindLumberjack:
		return 5
	case GameObjectKindWarehouse:
		return 10
	case GameObjectKindPioneer:
		return -10
	default:
		return 0
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
	case GameObjectKindPioneer:
		return 25, 4, 0
	case GameObjectKindWarehouse:
		return 100, 6, 0
	default:
		return 0, 0, 0
	}
}

// InsertDrawables inserts the drawables of the game object
func (g *GameObject) InsertDrawables(index int16) {
	if g.Kind == GameObjectKindShip {
		sprites.InsertDrawable(sprites.Drawable{X: byte(g.X), Y: byte(g.Y), Ref: index})
		return
	}

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

	if g.Kind == GameObjectKindWarehouse {
		if island.CheckLand(int(g.X)+1, int(g.Y), 1, int(h)) {
			return false
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

const tickRate = 60 * 30

// Simulate simulates the object for the current tick count
func (g *GameObject) Simulate(tick int64, ref int) {
	cost := g.GetObjectMaintenance()
	if tick%(tickRate) == 0 && island.Gold >= cost {
		island.Gold -= cost
	}
	switch g.Kind {
	case GameObjectKindChurch:
		if tick%(tickRate) == 0 {
			g.forEachObjOfKind(func(isl *island.Island) {
				island.Gold += 5
			}, GameObjectKindPioneer)
		}
	case GameObjectKindLumberjack:
		if tick%(tickRate) == 0 {
			g.forEachObjOfKind(func(isl *island.Island) {
				if isl.Wood < 255 {
					isl.Wood++
				}
			}, GameObjectKindTree)
		}
	case GameObjectKindShip:
		s := GetShipInfo(ref)

		const speed = 0.03

		if s != nil {
			if s.DestinationX > g.X {
				if s.OffsetX < 1 {
					s.OffsetX += speed
				} else {
					s.OffsetX = 0
					g.X += 1
				}
				s.Direction = SouthEast
			} else if s.DestinationX < g.X {
				if s.OffsetX > -1 {
					s.OffsetX -= speed
				} else {
					s.OffsetX = 0
					g.X -= 1
				}
				s.Direction = NorthWest
			} else if s.DestinationY > g.Y {
				if s.OffsetY < 1 {
					s.OffsetY += speed
				} else {
					s.OffsetY = 0
					g.Y += 1
				}
				s.Direction = SouthWest
			} else if s.DestinationY < g.Y {
				if s.OffsetY > -1 {
					s.OffsetY -= speed
				} else {
					s.OffsetY = 0
					g.Y -= 1
				}
				s.Direction = NorthEast
			}
		}
	}
}

func (g *GameObject) forEachObjOfKind(cb func(*island.Island), kind byte) {
	radius := g.GetRadius()
	width, height := g.GetObjectSize()
	for x := int(g.X) + radius; x >= int(g.X)-radius-int(width); x-- {
		for y := int(g.Y) + radius; y >= int(g.Y)-radius-int(height); y-- {
			obj, _ := GetObjectAt(x, y)

			if obj != nil && obj.Kind == kind {
				isl := island.FindIsland(x, y)
				cb(isl)
			}
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
			case GameObjectKindMountain:
				sprites.DrawMountainTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindChurch:
				sprites.DrawChurchTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindLumberjack:
				sprites.DrawLumberjackTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindPioneer:
				sprites.DrawPioneerTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindWarehouse:
				sprites.DrawWarehouseTile(int(drawable.X), int(drawable.Y), gameObject.X, gameObject.Y)
				break
			case GameObjectKindShip:
				si := GetShipInfo(int(drawable.Ref))

				if si != nil {
					flip := false

					if si.Direction == SouthEast {
						flip = true
					}

					sprites.DrawShipFrontTile(float64(gameObject.X)+si.OffsetX, float64(gameObject.Y)+si.OffsetY, flip)
				}
				break
			}
		}
	}
}

func GetShipInfo(ref int) *ShipInfo {
	for i := range ShipInfos {
		if ShipInfos[i].Ref == ref {
			return &ShipInfos[i]
		}
	}
	return nil
}

// GenerateGameObjects generates game objects
func GenerateGameObjects() {
	counter := 0
	for x := 0; x < island.Size; x++ {
		for y := 0; y < island.Size; y++ {
			if random.R.Float64() < 0.005 && counter < len(GameObjects) {
				object := GameObject{X: byte(x), Y: byte(y), Kind: GameObjectKindMountain}
				if object.CanBeBuilt() {
					GameObjects[counter] = object
					GameObjects[counter].InsertDrawables(int16(counter))
					counter++
				}
			} else if random.R.Float64() < 0.04 && counter < len(GameObjects) {
				object := GameObject{X: byte(x), Y: byte(y), Kind: GameObjectKindTree}
				if object.CanBeBuilt() {
					GameObjects[counter] = object
					GameObjects[counter].InsertDrawables(int16(counter))
					counter++
				}
			}
		}
	}

	ship := GameObject{X: byte(10), Y: byte(10), Kind: GameObjectKindShip}
	GameObjects[counter] = ship
	GameObjects[counter].InsertDrawables(int16(counter))
	ShipInfos[0] = ShipInfo{Ref: counter, OffsetX: 0, OffsetY: 0, DestinationX: 10, DestinationY: 10, Direction: 0, Cargo: 0, Amount: 0}
}

// SimulateObjects simulates all objects
func SimulateObjects(tick int64) {
	for i := range GameObjects {
		if GameObjects[i].Kind != 0 {
			GameObjects[i].Simulate(tick, i)
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
func GetObjectAt(x, y int) (*GameObject, int) {
	for i := range GameObjects {
		if GameObjects[i].Kind != 0 {
			width, height := GameObjects[i].GetObjectSize()
			if x >= int(GameObjects[i].X)-int(width)+1 && x <= int(GameObjects[i].X) && y >= int(GameObjects[i].Y)-int(height)+1 && y <= int(GameObjects[i].Y) {
				return &GameObjects[i], i
			}
		}
	}
	return nil, -1
}

// ClearGameObjects clears the game objects
func ClearGameObjects() {
	for i := range GameObjects {
		GameObjects[i] = GameObject{}
	}
}
