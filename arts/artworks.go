package arts

import (
	"github.com/Oppodelldog/logo-go/turtle"
	"math"
)

type TurtleFunc func(t turtle.Turtler, w, h float64)

func MalaikaKelly(P float64) TurtleFunc {
	return func(t turtle.Turtler, w, h float64) {
		cx := w / 2
		cy := h / 2

		for i:=float64(0); i < 255;i++ {
			for x:=0; x< 10; x++ {
				t.C(255, uint8(x*10), uint8(i), 255)
				t.TR(36 * i * float64(x))
				t.S(cx, cy)
				t.M(100 + P)
				t.TR(80)
				t.M(100)
				t.TL(20)
				t.M(100 + P*4)
				t.TR(200 + P)
				t.M(100)
				t.TR(80 + P)
				t.M(100)
				t.TR(30 + P)
				t.M(50)
				t.TR(20 + P*i)
				t.M(30)
				t.TR(10 + P)
				t.M(20)
				t.TR(1 + P)
				t.M(5)
				t.G().Stroke()
				t.G().BeginPath()
			}
		}
	}
}

func UrsulaConnolly(P float64) TurtleFunc {
	return func(t turtle.Turtler, w, h float64) {
		cx := w / 2
		cy := h / 2
		t.S(cx+200, cy+200)
		t.C(255, 200, 0, 255)
		for i := float64(1); i < 4360; i++ {
			t.M(i / (P*math.Cos(P) + 0.01))
			t.TR(10 + P/i)
		}
		t.G().FillStroke()

		t.G().BeginPath()
		t.S(cx+200, cy+200)
		t.C(255, 0, 0, 255)
		for i := float64(1); i < 4360; i++ {
			t.M(i / (P*math.Cos(P) + 0.01))
			t.TR(10 + P/i)
		}
	}
}

func DanicaVelez(P float64) TurtleFunc {
	return func(t turtle.Turtler, w, h float64) {
		cx := w / 2
		cy := h / 2
		t.S(cx+200, cy+200)

		for j := float64(0); j < 180; j += 1 {
			t.C(0, 255, 0, uint8(j))
			for i := float64(0); i < 20; i += 1 {
				t.M(j - P)
				t.TR(i - j*P)
			}
		}
	}
}

func DieselMalone(P float64) TurtleFunc {
	return func(t turtle.Turtler, w, h float64) {
		cx := w / 2
		cy := h / 2
		t.S(cx+200, cy+200)

		for j := float64(0); j < 180; j += 1 {
			for i := float64(0); i < 20; i += 1 {
				t.M(j - P)
				t.TR(i)
			}
		}
	}
}

func ChelsyWagner(P float64) TurtleFunc {
	return func(t turtle.Turtler, w, h float64) {
		cx := w / 2
		cy := h / 2
		t.S(cx, cy)

		for i := float64(0); i < 110; i += 0.1 {
			t.M(i)
			t.TR(P)
		}
	}
}
