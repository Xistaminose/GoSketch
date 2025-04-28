package shapes

import (
	"image/color"
)

// Canvas interface defines what our shapes need to know about the drawing surface
type Canvas interface {
	Set(x, y int, clr color.Color)
	GetWidth() int
	GetHeight() int
}

// Drawable é uma interface interna que define comportamentos para desenho
type Drawable interface {
	Fill(canvas Canvas, fillColor color.Color, fillEnabled bool)
	Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64)
}

// Shape é a interface que todas as formas devem implementar
type Shape interface {
	Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64)
}

// BaseShape implementa funcionalidade comum para todas as formas
type BaseShape struct{}

// Draw implementa a interface Shape chamando fill e stroke na ordem correta
func (b *BaseShape) Draw(canvas Canvas, d Drawable, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	if canvas == nil {
		return
	}
	d.Fill(canvas, fillColor, fillEnabled)
	d.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
} 