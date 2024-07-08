package main

// Pt is a point in 2D space.
type Pt struct{ X, Y int }

// Rect is a rectangle.
type Rect struct{ Min, Max Pt }

// ContainsPt reports whether r contains p.
func (r Rect) ContainsPt(pt Pt) bool {
	return true &&
		pt.X >= r.Min.X && pt.X < r.Max.X &&
		pt.Y >= r.Min.Y && pt.Y < r.Max.Y
}

func main() {

}
