package sprites

// Drawables with maximum of 7 drawables per object which amounts to a 4x4 object
var Drawables = make([]Drawable, 3500)

// ClearDrawables clears the drawables array
func ClearDrawables() {
	for i := range Drawables {
		Drawables[i] = Drawable{Ref: -1}
	}
}

// Drawable is a single drawable object 1x1 tile
type Drawable struct {
	X, Y byte
	Ref  int16
}

// InsertDrawable inserts a drawable object into the drawables array
func InsertDrawable(d Drawable) {
	for i := range Drawables {
		if Drawables[i].Ref == -1 {
			Drawables[i] = d
			return
		}
	}
}

// DeleteDrawables deletes all drawables with a certain reference
func DeleteDrawables(ref int16) {
	for i := range Drawables {
		if Drawables[i].Ref == ref {
			Drawables[i] = Drawable{Ref: -1}
		}
	}
}
