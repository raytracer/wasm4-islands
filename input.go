package main

import (
	"cart/island"
	"cart/objects"
	"cart/sprites"
	"cart/w4"
)

func handleInput() {
	mouseButtons := *w4.MOUSE_BUTTONS
	pressedThisFrame := mouseButtons & (mouseButtons ^ previousMouseButtons)
	previousMouseButtons = mouseButtons

	gamepad := *w4.GAMEPAD1
	pressedThisFrameGamepad := gamepad & (gamepad ^ previousButtons)
	previousButtons = gamepad

	if pressedThisFrameGamepad == w4.BUTTON_1 {
		if mode == gameMode {
			mode = buildMode
		} else if mode == startMode {
			mode = gameMode
			initGame()
		}
	}

	if pressedThisFrameGamepad == w4.BUTTON_2 {
		if mode == minimapMode {
			mode = gameMode
		} else if mode == gameMode {
			mode = minimapMode
		}
	}

	if pressedThisFrame == w4.MOUSE_RIGHT {
		if mode == placementMode {
			mode = gameMode
		} else if mode == gameMode {
			obj := getSelectedObject()
			if obj != nil && obj.Kind == objects.GameObjectKindShip {
				si := objects.GetShipInfo(selectedObjectRef)
				if si != nil {
					tx, ty := getTileAtMousePosition()

					if !island.CheckIsLand(tx, ty) {
						si.DestinationX = byte(tx)
						si.DestinationY = byte(ty)
						w4.Tone(220|(400<<16), 1, 25, w4.TONE_TRIANGLE)
					}
				}
			}
		}
	}

	if pressedThisFrame == w4.MOUSE_LEFT {
		if mode == placementMode {
			tx, ty := getTileAtMousePosition()
			building := objects.GameObject{X: byte(tx), Y: byte(ty), Kind: byte(placementType)}

			if building.CanBeBuilt() {
				building.Place()
				w4.Tone(25, 5|(10<<8), 25, w4.TONE_TRIANGLE)
			}
		} else if mode == buildMode {
			if *w4.MOUSE_Y < 80 {
				if *w4.MOUSE_X < 80 {
					placementType = objects.GameObjectKindChurch
				} else {
					placementType = objects.GameObjectKindLumberjack
				}
			} else if *w4.MOUSE_Y >= 80 {
				if *w4.MOUSE_X < 80 {
					placementType = objects.GameObjectKindPioneer
				} else {
					placementType = objects.GameObjectKindWarehouse
				}
			}

			mode = placementMode
			selectedObjectRef = -1
		} else if mode == gameMode {
			tx, ty := getTileAtMousePosition()
			obj, ref := objects.GetObjectAt(tx, ty)

			if obj != nil && obj.Kind != objects.GameObjectKindTree && obj.Kind != objects.GameObjectKindMountain {
				selectedObjectRef = ref
			} else {
				selectedObjectRef = -1
			}
		}
	}

	if mode == gameMode || mode == placementMode {
		handleScreenBorderInput()
	}
}

func handleScreenBorderInput() {
	if *w4.MOUSE_X < 5 {
		sprites.CameraX++
	}

	if *w4.MOUSE_X > 155 {
		sprites.CameraX--
	}

	if *w4.MOUSE_Y < 5 {
		sprites.CameraY++
	}

	if *w4.MOUSE_Y > 155 {
		sprites.CameraY--
	}
}
