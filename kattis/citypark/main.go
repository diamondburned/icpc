package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// StoneSurfaces is a collection of stone surfaces.
type StoneSurfaces []StoneSurface

// Add adds the given rectangle into the stone surfaces. If no existing stone
// surfaces exist that touches the given rectangle, then a new stone surface is
// created.
func (ss *StoneSurfaces) Add(rect Rectangle) {
	// A rectangle can be placed in the middle of two current stone surfaceIxs,
	// joining them together. We need to keep track of the stone surfaceIxs that
	// this rectangle touches and combine them after.
	var surfaceIxs []int
	for i := range *ss {
		if (*ss)[i].Add(rect) {
			surfaceIxs = append(surfaceIxs, i)
		}
	}

	// log.Printf("adding rect %v to %d surfaces", rect, len(surfaceIxs))

	// If we found only 1 surface to add this rectangle to, then we're done.
	if len(surfaceIxs) == 1 {
		return
	}

	s := StoneSurface{
		Bounds: rect,
		Rects:  []Rectangle{rect},
		Area:   rect.Area(),
	}

	if len(surfaceIxs) == 0 {
		*ss = append(*ss, s)
		return
	}

	if len(surfaceIxs) == 1 {
		(*ss)[surfaceIxs[0]] = s
		return
	}

	// If we found at least 2 surfaces, then we have to join them together.
	if len(surfaceIxs) > 1 {
		for _, otherIx := range surfaceIxs {
			other := (*ss)[otherIx]

			s.Rects = append(s.Rects, other.Rects...)
			s.Area += other.Area

			// Update the bounds depending on the edge we're touching. Only expand
			// the bound in the direction of the edge.
			switch s.Bounds.Touches(rect) {
			case EdgeTop:
				s.Bounds.Width += rect.Width
			case EdgeBottom:
				s.Bounds.Width += rect.Width
				s.Bounds.BottomLeft.X -= rect.Width
			case EdgeLeft:
				s.Bounds.Height += rect.Height
			case EdgeRight:
				s.Bounds.Height += rect.Height
				s.Bounds.BottomLeft.Y -= rect.Height
			default:
				panic("unreachable")
			}
		}
	}

	// Remove the old stone surfaces in undefined order. This is faster.
	for i := len(surfaceIxs) - 1; i >= 0; i-- {
		ix := surfaceIxs[i]
		(*ss)[ix] = (*ss)[len(*ss)-1]
		(*ss)[len(*ss)-1] = StoneSurface{}
		*ss = (*ss)[:len(*ss)-1]
	}

	// Add the new stone surface.
	*ss = append(*ss, s)
}

// StoneSurface represents a stone surface, which consists of a bunch of
// rectangles.
type StoneSurface struct {
	Rects  []Rectangle
	Bounds Rectangle
	Area   int
}

// Add adds the given rectangle into the stone surface.
func (s *StoneSurface) Add(rect Rectangle) bool {
	// Fast path: check if rect overlaps with our big rectangle.
	if !s.Bounds.Overlaps(rect) {
		return false
	}

	// Our rectangle doesn't touch the big rectangle. Check the smaller
	// rectangles to ensure exactly.
	var touchingEdge Edge
	for _, r := range s.Rects {
		if e := r.Touches(rect); e != NoEdge {
			touchingEdge = e
			break
		}
	}

	if touchingEdge == NoEdge {
		return false
	}

	s.Rects = append(s.Rects, rect)
	s.Area += rect.Area()

	// Update the bounds depending on the edge we're touching. Only expand the
	// bound in the direction of the edge.
	switch touchingEdge {
	case EdgeTop:
		s.Bounds.Width += rect.Width
	case EdgeBottom:
		s.Bounds.Width += rect.Width
		s.Bounds.BottomLeft.X -= rect.Width
	case EdgeLeft:
		s.Bounds.Height += rect.Height
	case EdgeRight:
		s.Bounds.Height += rect.Height
		s.Bounds.BottomLeft.Y -= rect.Height
	}

	return true
}

// Edge represents an edge of a rectangle.
type Edge int8

const (
	NoEdge Edge = iota
	EdgeLeft
	EdgeRight
	EdgeTop
	EdgeBottom
)

// Rectangle represents a rectangle.
type Rectangle struct {
	BottomLeft Pt
	Width      int32
	Height     int32
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() int {
	area := int(r.Width * r.Height)
	if area < 0 {
		area = -area
	}
	return area
}

// Overlaps returns true if r overlaps or touches other.
func (r Rectangle) Overlaps(other Rectangle) bool {
	r1 := r.BottomLeft
	r2 := r.BottomLeft.Add(Pt{r.Width, r.Height})

	o1 := other.BottomLeft
	o2 := other.BottomLeft.Add(Pt{other.Width, other.Height})

	return !(false ||
		r1.X > o2.X || r2.X < o1.X ||
		r1.Y > o2.Y || r2.Y < o1.Y)
}

// Touches returns true if the rectangle touches the other rectangle.
func (r Rectangle) Touches(other Rectangle) Edge {
	switch {
	case r.BottomLeft.X == other.BottomLeft.X+other.Width:
		return EdgeBottom
	case r.BottomLeft.X+r.Width == other.BottomLeft.X:
		return EdgeTop
	case r.BottomLeft.Y == other.BottomLeft.Y+other.Height:
		return EdgeLeft
	case r.BottomLeft.Y+r.Height == other.BottomLeft.Y:
		return EdgeRight
	default:
		return NoEdge
	}
}

// Pt is a point in 2D space.
type Pt struct{ X, Y int32 }

// Add does pt + other.
func (pt Pt) Add(other Pt) Pt {
	return Pt{pt.X + other.X, pt.Y + other.Y}
}

func main() {
	var n int
	mustSscanf(mustReadUntil('\n'), "%d", &n)

	var surfaces StoneSurfaces
	for i := 0; i < n; i++ {
		var r Rectangle
		mustSscanf(mustReadUntil('\n'), "%d %d %d %d", &r.BottomLeft.X, &r.BottomLeft.Y, &r.Width, &r.Height)
		surfaces.Add(r)
	}

	sort.Slice(surfaces, func(i, j int) bool {
		return surfaces[i].Area < surfaces[j].Area
	})

	fmt.Println(surfaces[len(surfaces)-1].Area)
}

var stdio = bufio.NewReader(os.Stdin)

func mustReadUntil(delim byte) string {
	s, err := stdio.ReadString(delim)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return s
}

func mustSscanf(s, f string, v ...interface{}) {
	_, err := fmt.Sscanf(s, f, v...)
	if err != nil {
		panic(err)
	}
}
