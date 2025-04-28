package shapes

import (
	"image/color"
)

// PointShape implementa Shape para pontos
type PointShape struct {
	BaseShape
	X, Y float64
}

// Fill para ponto é um no-op (pontos não têm preenchimento)
func (p *PointShape) Fill(canvas Canvas, fillColor color.Color, fillEnabled bool) {
	// Pontos não têm preenchimento
}

// Stroke desenha o ponto com a espessura definida
func (p *PointShape) Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64) {
	if !strokeEnabled {
		return
	}
	
	// Um ponto é renderizado como um pequeno círculo de raio = strokeWeight / 2
	radius := int(strokeWeight / 2)
	if radius < 1 {
		radius = 1
	}
	
	// Desenha um círculo preenchido
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			if dx*dx+dy*dy <= radius*radius {
				canvas.Set(int(p.X)+dx, int(p.Y)+dy, strokeColor)
			}
		}
	}
}

// Draw executa stroke no ponto
func (p *PointShape) Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	p.BaseShape.Draw(canvas, p, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
}

// NewPoint cria uma nova forma de ponto
func NewPoint(x, y float64) *PointShape {
	return &PointShape{X: x, Y: y}
} 