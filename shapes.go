package GoSketch

import (
	"image/color"
	"math"
)

// Drawable é uma interface interna que define comportamentos para desenho
type Drawable interface {
	fill()
	stroke()
}

// BaseShape implementa funcionalidade comum para todas as formas
type BaseShape struct {
	// Não tem campos próprios, apenas métodos
}

// Draw implementa a interface Shape chamando fill e stroke na ordem correta
func (b *BaseShape) Draw(d Drawable) {
	if canvas == nil {
		return
	}
	d.fill()
	d.stroke()
}

// EllipseShape é uma implementação de Shape para elipses
type EllipseShape struct {
	BaseShape
	X, Y, Rx, Ry float64
}

// Fill define a cor de preenchimento para formas subsequentes
func Fill(c color.Color) { fillColor = c; fillEnabled = true }

// NoFill desabilita preenchimento
func NoFill() { fillEnabled = false }

// Stroke define a cor de contorno para formas subsequentes
func Stroke(c color.Color) { strokeColor = c; strokeEnabled = true }

// NoStroke desabilita contorno
func NoStroke() { strokeEnabled = false }

// StrokeWeight define a espessura do contorno para formas subsequentes
func StrokeWeight(w float64) {
	strokeWeight = w
}

// Preenche a área interna da elipse
func (e *EllipseShape) fill() {
	if !fillEnabled {
		return
	}
	for dx := -int(e.Rx); dx <= int(e.Rx); dx++ {
		for dy := -int(e.Ry); dy <= int(e.Ry); dy++ {
			if float64(dx*dx)/(e.Rx*e.Rx)+float64(dy*dy)/(e.Ry*e.Ry) <= 1 {
				canvas.img.Set(int(e.X)+dx, int(e.Y)+dy, fillColor)
			}
		}
	}
}

// Desenha o contorno da elipse conforme strokeWeight
func (e *EllipseShape) stroke() {
	if !strokeEnabled {
		return
	}
	steps := int(2 * math.Pi * math.Max(e.Rx, e.Ry))
	for i := 0; i < steps; i++ {
		theta := 2 * math.Pi * float64(i) / float64(steps)
		px := e.X + e.Rx*math.Cos(theta)
		py := e.Y + e.Ry*math.Sin(theta)
		for sw := -int(strokeWeight / 2); sw <= int(strokeWeight/2); sw++ {
			canvas.img.Set(int(px)+sw, int(py), strokeColor)
			canvas.img.Set(int(px), int(py)+sw, strokeColor)
		}
	}
}

// Draw executa fill e stroke na ordem correta
func (e *EllipseShape) Draw() {
	e.BaseShape.Draw(e)
}

// RectangleShape implementa Shape para retângulos preenchidos
type RectangleShape struct{ 
	BaseShape
	X, Y, W, H float64 
}

// Preenche a área interna do retângulo, se fillEnabled
func (r *RectangleShape) fill() {
	if !fillEnabled {
		return
	}
	for dx := 0; dx < int(r.W); dx++ {
		for dy := 0; dy < int(r.H); dy++ {
			canvas.img.Set(int(r.X)+dx, int(r.Y)+dy, fillColor)
		}
	}
}

// Desenha o contorno do retângulo conforme strokeWeight, se strokeEnabled
func (r *RectangleShape) stroke() {
	if !strokeEnabled {
		return
	}
	for sw := 0; sw < int(strokeWeight); sw++ {
		for dx := 0; dx < int(r.W); dx++ {
			canvas.img.Set(int(r.X)+dx, int(r.Y)+sw, strokeColor)
			canvas.img.Set(int(r.X)+dx, int(r.Y+r.H)-sw, strokeColor)
		}
		for dy := 0; dy < int(r.H); dy++ {
			canvas.img.Set(int(r.X)+sw, int(r.Y)+dy, strokeColor)
			canvas.img.Set(int(r.X+r.W)-sw, int(r.Y)+dy, strokeColor)
		}
	}
}

// Draw executa fill e stroke no retângulo
func (r *RectangleShape) Draw() {
	r.BaseShape.Draw(r)
}

// SquareShape implementa Shape para quadrados
type SquareShape struct{ 
	BaseShape
	X, Y, Size float64 
}

// fill implementa o preenchimento do quadrado
func (s *SquareShape) fill() {
	if !fillEnabled {
		return
	}
	for dx := 0; dx < int(s.Size); dx++ {
		for dy := 0; dy < int(s.Size); dy++ {
			canvas.img.Set(int(s.X)+dx, int(s.Y)+dy, fillColor)
		}
	}
}

// stroke implementa o contorno do quadrado
func (s *SquareShape) stroke() {
	if !strokeEnabled {
		return
	}
	for sw := 0; sw < int(strokeWeight); sw++ {
		for d := 0; d < int(s.Size); d++ {
			canvas.img.Set(int(s.X)+d, int(s.Y)+sw, strokeColor)
			canvas.img.Set(int(s.X)+d, int(s.Y+s.Size)-sw, strokeColor)
			canvas.img.Set(int(s.X)+sw, int(s.Y)+d, strokeColor)
			canvas.img.Set(int(s.X+s.Size)-sw, int(s.Y)+d, strokeColor)
		}
	}
}

// Draw executa fill e stroke no quadrado
func (s *SquareShape) Draw() {
	s.BaseShape.Draw(s)
}

// Forma adicional: linha (sem preenchimento, só contorno)
type LineShape struct {
	BaseShape
	X1, Y1, X2, Y2 float64
}

// fill para linha é um no-op (linhas não têm preenchimento)
func (l *LineShape) fill() {
	// Linhas não têm preenchimento
}

// stroke desenha a linha com a espessura definida
func (l *LineShape) stroke() {
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
	
	x, y := int(l.X1), int(l.Y1)
	
	for {
		// Desenha o ponto atual com a espessura definida
		for sw := -int(strokeWeight/2); sw <= int(strokeWeight/2); sw++ {
			for swY := -int(strokeWeight/2); swY <= int(strokeWeight/2); swY++ {
				canvas.img.Set(x+sw, y+swY, strokeColor)
			}
		}
		
		if x == int(l.X2) && y == int(l.Y2) {
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
	}
}

// Draw implementa a interface Shape para linhas
func (l *LineShape) Draw() {
	l.BaseShape.Draw(l)
}
