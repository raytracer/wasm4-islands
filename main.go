package main

import (
	"cart/island"
	"cart/objects"
	"cart/random"
	"cart/sprites"
	"cart/w4"
)

const (
	gameMode      = iota
	buildMode     = iota
	placementMode = iota
	minimapMode   = iota
	startMode     = iota
	helpMode      = iota
)

var mode = startMode

var ticks int64 = 0
var year uint = 1400

var selectedObject *objects.GameObject = nil
var placementType = objects.GameObjectKindChurch

//go:export start
func start() {
}

func initGame() {
	random.InitRand(ticks)
	objects.ClearShipInfos()
	objects.ClearGameObjects()
	sprites.ClearDrawables()
	island.GenerateIslands()
	objects.GenerateGameObjects()
	ticks = 0
}

//go:export update
func update() {
	handleInput()
	if mode != startMode && mode != helpMode {
		draw()
		objects.SimulateObjects(ticks)

		if (ticks % 60) == 0 {
			year++
		}
	} else if mode == startMode {
		drawStartScreen()
	}

	ticks++
}

func draw() {
	drawScene()
	drawUI()
}

func drawScene() {
	if mode == gameMode || mode == placementMode {
		island.DrawIsland()
		objects.DrawGameObjects()
		drawSelectedRect()
		drawRadius(selectedObject)
	}

	if mode == placementMode {
		tx, ty := sprites.ScreenToIsoTransform(int(*w4.MOUSE_X)-sprites.CameraX, int(*w4.MOUSE_Y)-sprites.CameraY+sprites.TileHeight)
		x1, y1 := sprites.IsoTransform(tx, ty)

		obj := objects.GameObject{X: byte(tx), Y: byte(ty), Kind: byte(placementType)}
		drawRadius(&obj)

		switch placementType {
		case objects.GameObjectKindChurch:
			sprites.DrawChurch(x1+sprites.CameraX, y1+sprites.CameraY)
			break
		case objects.GameObjectKindLumberjack:
			sprites.DrawLumberjack(x1+sprites.CameraX, y1+sprites.CameraY)
			break
		case objects.GameObjectKindPioneer:
			sprites.DrawPioneer(x1+sprites.CameraX, y1+sprites.CameraY)
			break
		}
	}
}

var previousMouseButtons byte
var previousButtons byte

func getTileAtMousePosition() (int, int) {
	return sprites.ScreenToIsoTransform(int(*w4.MOUSE_X)-sprites.CameraX, int(*w4.MOUSE_Y)-sprites.CameraY+sprites.TileHeight)
}
