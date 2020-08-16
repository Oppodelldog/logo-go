package turtle

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"math"
)

type Turtler interface {
	C(r, g, b, a uint8)
	S(x, y float64)
	M(d float64)
	TR(a float64)
	TL(a float64)
	G() *draw2dimg.GraphicContext
}

func New(g *draw2dimg.GraphicContext) *Turtle {
	return &Turtle{
		g: g,
	}
}

type Turtle struct {
	// x pos
	x float64
	// y pos
	y float64
	// a - angle
	a float64
	// d - distance
	d float64
	// g - graphics context
	g *draw2dimg.GraphicContext
}

func (t *Turtle) C(r, g, b, a uint8) {
	c := color.RGBA{r, g, b, a}
	t.g.SetStrokeColor(c)
}

func (t *Turtle) S(x, y float64) {
	t.x = x
	t.y = y
	mustNumberCoords(t.x, t.x)
	t.g.MoveTo(t.x, t.y)
}

func (t *Turtle) M(d float64) {
	a := t.a * math.Pi / 180
	t.x = t.x + (d * math.Cos(a))
	t.y = t.y + (d * math.Sin(a))
	mustNumberCoords(t.x, t.x)
	t.g.LineTo(t.x, t.y)
}

func (t *Turtle) G() *draw2dimg.GraphicContext {
	return t.g
}
func mustNumberCoords(v1, v2 float64) {
	mustNumber(v1)
	mustNumber(v2)
}
func mustNumber(v float64) {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		panic(fmt.Sprintf("Must be a numbr, got : %v", v))
	}
}
func (t *Turtle) TR(a float64) {
	t.a = t.a + a
	mustNumber(t.a)
}

func (t *Turtle) TL(a float64) {
	t.a = t.a - a
	mustNumber(t.a)
}
