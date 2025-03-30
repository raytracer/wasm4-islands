package main

import (
	"cart/island"
	"cart/objects"
	"cart/sprites"
	"cart/w4"
	"math"
)

func drawSelectedRect() {
	selectedObject := getSelectedObject()
	if selectedObject != nil && ticks%60 < 30 {
		width, height := selectedObject.GetObjectSize()
		x, y := float64(selectedObject.X), float64(selectedObject.Y)

		if selectedObject.Kind == objects.GameObjectKindShip {
			si := objects.GetShipInfo(selectedObjectRef)

			if si != nil {
				x += si.OffsetX
				y += si.OffsetY
			}
		}

		sx1, sy1 := sprites.IsoTransform(x, y)
		sx2, sy2 := sprites.IsoTransform(x-float64(width), y-float64(height))

		tx1, ty1 := sprites.IsoTransform(x-float64(width), y)
		tx2, ty2 := sprites.IsoTransform(x, y-float64(height))

		*w4.DRAW_COLORS = 0x4
		w4.Line(tx1+sprites.CameraX, ty1+sprites.CameraY, sx1+sprites.CameraX, sy1+sprites.CameraY)
		w4.Line(sx1+sprites.CameraX, sy1+sprites.CameraY, tx2+sprites.CameraX, ty2+sprites.CameraY)
		w4.Line(tx1+sprites.CameraX, ty1+sprites.CameraY, sx2+sprites.CameraX, sy2+sprites.CameraY)
		w4.Line(sx2+sprites.CameraX, sy2+sprites.CameraY, tx2+sprites.CameraX, ty2+sprites.CameraY)
	}
}

func drawRadius(obj *objects.GameObject) {
	if obj != nil && ticks%60 < 30 {
		width, height := obj.GetObjectSize()
		radius := float64(obj.GetRadius())

		if radius > 0 {
			x, y := float64(obj.X), float64(obj.Y)

			sx1, sy1 := sprites.IsoTransform(x+radius, y+radius)
			sx2, sy2 := sprites.IsoTransform(x-float64(width)-radius, y-float64(height)-radius)

			tx1, ty1 := sprites.IsoTransform(x-float64(width)-radius, y+radius)
			tx2, ty2 := sprites.IsoTransform(x+radius, y-float64(height)-radius)

			*w4.DRAW_COLORS = 0x4
			w4.Line(tx1+sprites.CameraX, ty1+sprites.CameraY, sx1+sprites.CameraX, sy1+sprites.CameraY)
			w4.Line(sx1+sprites.CameraX, sy1+sprites.CameraY, tx2+sprites.CameraX, ty2+sprites.CameraY)
			w4.Line(tx1+sprites.CameraX, ty1+sprites.CameraY, sx2+sprites.CameraX, sy2+sprites.CameraY)
			w4.Line(sx2+sprites.CameraX, sy2+sprites.CameraY, tx2+sprites.CameraX, ty2+sprites.CameraY)
		}
	}
}

func drawUI() {
	if mode == gameMode || mode == placementMode {
		if getSelectedObject() != nil {
			drawSelectedObjectUI()
		}
		drawStatusLine()
	} else if mode == buildMode {
		drawBuildUI()
	} else if mode == minimapMode {
		drawMiniMap()
	}
}

func drawSelectedObjectUI() {
	width := 50
	*w4.DRAW_COLORS = 0x1
	w4.Rect(160-width, 8, uint(width), 160-8)
	*w4.DRAW_COLORS = 0x4
	w4.Line(159-width, 8, 159-width, 160)

	selectedObject := getSelectedObject()

	if selectedObject != nil && selectedObject.Kind != objects.GameObjectKindPioneer {
		costs := selectedObject.GetObjectMaintenance()
		w4.Text("Costs", 115, 10)
		sprites.DrawGold(125, 25)
		sprites.DrawNumber(costs, 130, 26, 2)
	}
}

func drawStatusLine() {
	*w4.DRAW_COLORS = 0x1
	w4.Rect(0, 0, 160, 7)

	sprites.DrawNumber(island.Gold, 1, 1, 5)

	sprites.DrawGold(26, 0)
	tx, ty := sprites.ScreenToIsoTransform(int(*w4.MOUSE_X)-sprites.CameraX, int(*w4.MOUSE_Y)-sprites.CameraY+sprites.TileHeight)

	island := island.FindIsland(tx, ty)

	if island != nil {
		*w4.DRAW_COLORS = 0x4
		sprites.DrawNumber(int(island.Wood), 35, 1, 3)
		sprites.DrawNumber(int(island.Stone), 70, 1, 3)

		sprites.DrawWood(50, 0)
		sprites.DrawStone(85, 0)
	}

	sprites.DrawNumber(int(year), 145, 1, 3)

	*w4.DRAW_COLORS = 0x4
	w4.Line(0, 7, 160, 7)
}

func drawStartScreen() {
	*w4.DRAW_COLORS = 0x4
	w4.Text("Islands", 55, 30)
	w4.Text("Start with \x80", 30, 90)
	w4.Text("Show Help with \x81", 15, 120)
}

func drawBuildUI() {
	sprites.DrawChurch(40, 62)
	*w4.DRAW_COLORS = 0x30
	w4.Rect(4, 4, 64, 64)
	*w4.DRAW_COLORS = 0x3
	w4.Text("Church", 12, 72)

	sprites.DrawLumberjack(120, 62)
	*w4.DRAW_COLORS = 0x30
	w4.Rect(88, 4, 64, 64)
	*w4.DRAW_COLORS = 0x3
	w4.Text("Lumberjack", 80, 72)

	sprites.DrawPioneer(36, 142)
	*w4.DRAW_COLORS = 0x30
	w4.Rect(4, 83, 64, 64)
	*w4.DRAW_COLORS = 0x3
	w4.Text("House", 16, 150)
}

func drawMiniMap() {
	sx1, sy1 := sprites.ScreenToIsoTransform(-sprites.CameraX, -sprites.CameraY)
	sx2, sy2 := sprites.ScreenToIsoTransform(-sprites.CameraX+160, -sprites.CameraY)
	sx3, sy3 := sprites.ScreenToIsoTransform(-sprites.CameraX+160, -sprites.CameraY+160)
	sx4, sy4 := sprites.ScreenToIsoTransform(-sprites.CameraX, -sprites.CameraY+160)

	for x := range 160 {
		for y := range 160 {
			convertedX := int(math.Floor(float64(x) / 160.0 * float64(island.Size)))
			convertedY := int(math.Floor(float64(y) / 160.0 * float64(island.Size)))
			if island.CheckIsLand(convertedX, convertedY) {
				*w4.DRAW_COLORS = 0x2
				w4.Rect(x, y, 1, 1)
			}
		}
	}

	*w4.DRAW_COLORS = 0x4
	cx1, cy1 := int(math.Floor(float64(sx1)/island.Size*float64(160.0))), int(math.Floor(float64(sy1)/island.Size*float64(160.0)))
	cx2, cy2 := int(math.Floor(float64(sx2)/island.Size*float64(160.0))), int(math.Floor(float64(sy2)/island.Size*float64(160.0)))
	cx3, cy3 := int(math.Floor(float64(sx3)/island.Size*float64(160.0))), int(math.Floor(float64(sy3)/island.Size*float64(160.0)))
	cx4, cy4 := int(math.Floor(float64(sx4)/island.Size*float64(160.0))), int(math.Floor(float64(sy4)/island.Size*float64(160.0)))
	w4.Line(cx1, cy1, cx2, cy2)
	w4.Line(cx2, cy2, cx3, cy3)
	w4.Line(cx3, cy3, cx4, cy4)
	w4.Line(cx4, cy4, cx1, cy1)
}
