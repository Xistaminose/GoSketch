package shapes

import (
	"image/color"
	"math"
)

// LineShape representa uma linha entre dois pontos
type LineShape struct {
	BaseShape
	X1, Y1, X2, Y2 float64
}

// Fill para linha é um no-op (linhas não têm preenchimento)
func (l *LineShape) Fill(canvas Canvas, fillColor color.Color, fillEnabled bool) {
	// Linhas não têm preenchimento
}

// Stroke desenha a linha com a espessura definida
func (l *LineShape) Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64) {
	if !strokeEnabled {
		return
	}
	// Algoritmo de Bresenham para desenho de linhas
	dx := int(math.Abs(l.X2 - l.X1))
	dy := int(math.Abs(l.Y2 - l.Y1))
	sx := -1
	if l.X1 < l.X2 {
		sx = 1
	}
	sy := -1
	if l.Y1 < l.Y2 {
		sy = 1
	}
	err := dx - dy

	// Converte para inteiros apenas uma vez
	x, y := int(l.X1), int(l.Y1)
	destX, destY := int(l.X2), int(l.Y2)
	
	// Verifica se a linha é apenas um ponto
	if x == destX && y == destY {
		drawThickPoint(canvas, x, y, strokeColor, strokeWeight)
		return
	}

	// Limites do canvas para evitar desenho fora dos limites
	canvasWidth := canvas.GetWidth()
	canvasHeight := canvas.GetHeight()

	// Loop principal de desenho da linha
	for {
		// Verifica se o ponto está dentro dos limites do canvas
		if x >= 0 && x < canvasWidth && y >= 0 && y < canvasHeight {
			drawThickPoint(canvas, x, y, strokeColor, strokeWeight)
		}

		// Verifica se chegamos ao destino (usando proximidade ao invés de igualdade exata)
		if x == destX && y == destY {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
		
		// Mecanismo de segurança para evitar loops infinitos
		// Se percorremos mais que o dobro da distância entre os pontos, algo está errado
		if math.Abs(float64(x-int(l.X1))) > 2*math.Abs(l.X2-l.X1) || 
		   math.Abs(float64(y-int(l.Y1))) > 2*math.Abs(l.Y2-l.Y1) {
			break
		}
	}
}

// drawThickPoint é uma função auxiliar para desenhar um ponto com espessura
func drawThickPoint(canvas Canvas, x, y int, color color.Color, thickness float64) {
	radius := int(thickness / 2)
	if radius < 1 {
		radius = 1
	}
	
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			if dx*dx+dy*dy <= radius*radius {
				canvas.Set(x+dx, y+dy, color)
			}
		}
	}
}

// Draw implementa a interface Shape para linhas
func (l *LineShape) Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	l.BaseShape.Draw(canvas, l, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
}

// NewLine cria uma nova forma de linha
func NewLine(x1, y1, x2, y2 float64) *LineShape {
	return &LineShape{X1: x1, Y1: y1, X2: x2, Y2: y2}
} 