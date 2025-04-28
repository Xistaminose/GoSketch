package shapes

import (
	"image/color"
)

// RectangleShape implementa Shape para retângulos preenchidos
type RectangleShape struct {
	BaseShape
	X, Y, W, H float64
}

// Fill preenche a área interna do retângulo, se fillEnabled
func (r *RectangleShape) Fill(canvas Canvas, fillColor color.Color, fillEnabled bool) {
	if !fillEnabled {
		return
	}
	for dx := 0; dx < int(r.W); dx++ {
		for dy := 0; dy < int(r.H); dy++ {
			canvas.Set(int(r.X)+dx, int(r.Y)+dy, fillColor)
		}
	}
}

// Stroke desenha o contorno do retângulo conforme strokeWeight, se strokeEnabled
func (r *RectangleShape) Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64) {
	if !strokeEnabled {
		return
	}
	for sw := 0; sw < int(strokeWeight); sw++ {
		for dx := 0; dx < int(r.W); dx++ {
			canvas.Set(int(r.X)+dx, int(r.Y)+sw, strokeColor)
			canvas.Set(int(r.X)+dx, int(r.Y+r.H)-sw-1, strokeColor)
		}
		for dy := 0; dy < int(r.H); dy++ {
			canvas.Set(int(r.X)+sw, int(r.Y)+dy, strokeColor)
			canvas.Set(int(r.X+r.W)-sw-1, int(r.Y)+dy, strokeColor)
		}
	}
}

// Draw executa fill e stroke no retângulo
func (r *RectangleShape) Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	r.BaseShape.Draw(canvas, r, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
}

// NewRectangle cria uma nova forma de retângulo
func NewRectangle(x, y, w, h float64) *RectangleShape {
	return &RectangleShape{X: x, Y: y, W: w, H: h}
}

// SquareShape é especializado como um caso particular de retângulo onde W = H
type SquareShape struct {
	RectangleShape
}

// NewSquare cria uma nova forma de quadrado
func NewSquare(x, y, size float64) *RectangleShape {
	return NewRectangle(x, y, size, size)
} 