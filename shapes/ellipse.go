package shapes

import (
	"image/color"
	"math"
)

// EllipseShape é uma implementação de Shape para elipses
type EllipseShape struct {
	BaseShape
	X, Y, Rx, Ry float64
}

// Fill preenche a área interna da elipse
func (e *EllipseShape) Fill(canvas Canvas, fillColor color.Color, fillEnabled bool) {
	if !fillEnabled {
		return
	}
	for dx := -int(e.Rx); dx <= int(e.Rx); dx++ {
		for dy := -int(e.Ry); dy <= int(e.Ry); dy++ {
			if float64(dx*dx)/(e.Rx*e.Rx)+float64(dy*dy)/(e.Ry*e.Ry) <= 1 {
				canvas.Set(int(e.X)+dx, int(e.Y)+dy, fillColor)
			}
		}
	}
}

// Stroke desenha o contorno da elipse conforme strokeWeight
func (e *EllipseShape) Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64) {
	if !strokeEnabled {
		return
	}
	steps := int(2 * math.Pi * math.Max(e.Rx, e.Ry))
	for i := 0; i < steps; i++ {
		theta := 2 * math.Pi * float64(i) / float64(steps)
		px := e.X + e.Rx*math.Cos(theta)
		py := e.Y + e.Ry*math.Sin(theta)
		for sw := -int(strokeWeight / 2); sw <= int(strokeWeight/2); sw++ {
			canvas.Set(int(px)+sw, int(py), strokeColor)
			canvas.Set(int(px), int(py)+sw, strokeColor)
		}
	}
}

// Draw executa fill e stroke na ordem correta
func (e *EllipseShape) Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	e.BaseShape.Draw(canvas, e, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
}

// NewEllipse cria uma nova forma de elipse
func NewEllipse(x, y, rx, ry float64) *EllipseShape {
	return &EllipseShape{X: x, Y: y, Rx: rx, Ry: ry}
}

// NewCircle cria uma nova forma de círculo (caso especial de elipse)
func NewCircle(x, y, radius float64) *EllipseShape {
	return NewEllipse(x, y, radius, radius)
} 