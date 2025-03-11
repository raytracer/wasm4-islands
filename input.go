package main

import (
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
		}
	}

	if pressedThisFrame == w4.MOUSE_LEFT {
		if mode == placementMode {
			tx, ty := getTileAtMousePosition()
			church := objects.GameObject{X: byte(tx), Y: byte(ty), Kind: objects.GameObjectKindChurch}

			if church.CanBeBuilt() {
				church.Place()
			}
		} else if mode == buildMode {
			placementType = objects.GameObjectKindChurch
			mode = placementMode
			selectedObject = nil
		} else if mode == gameMode {
			tx, ty := getTileAtMousePosition()
			obj := objects.GetObjectAt(tx, ty)

			if obj != nil && obj.Kind != objects.GameObjectKindTree {
				selectedObject = obj
			} else {
				selectedObject = nil
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
