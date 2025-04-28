package shapes

import (
	"image/color"
	"math"
)

// TriangleShape implementa Shape para triângulos
type TriangleShape struct {
	BaseShape
	X1, Y1, X2, Y2, X3, Y3 float64 // Três vértices do triângulo
}

// Fill preenche a área interna do triângulo, se fillEnabled
func (t *TriangleShape) Fill(canvas Canvas, fillColor color.Color, fillEnabled bool) {
	if !fillEnabled {
		return
	}
	
	// Encontra os limites do triângulo
	minX := math.Min(t.X1, math.Min(t.X2, t.X3))
	maxX := math.Max(t.X1, math.Max(t.X2, t.X3))
	minY := math.Min(t.Y1, math.Min(t.Y2, t.Y3))
	maxY := math.Max(t.Y1, math.Max(t.Y2, t.Y3))
	
	// Verificação para triângulos degenerados (linha ou ponto)
	area := 0.5 * math.Abs((t.X1*(t.Y2-t.Y3) + t.X2*(t.Y3-t.Y1) + t.X3*(t.Y1-t.Y2)))
	if area < 0.01 { // Área muito pequena, triângulo degenerado
		return
	}
	
	// Limita os limites ao tamanho do canvas para evitar processamento desnecessário
	canvasWidth := canvas.GetWidth()
	canvasHeight := canvas.GetHeight()
	
	startX := int(math.Max(0, minX))
	endX := int(math.Min(float64(canvasWidth-1), maxX))
	startY := int(math.Max(0, minY))
	endY := int(math.Min(float64(canvasHeight-1), maxY))
	
	// Algoritmo de preenchimento por escaneamento
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			// Verifica se o ponto (x,y) está dentro do triângulo
			if isPointInTriangle(float64(x), float64(y), t.X1, t.Y1, t.X2, t.Y2, t.X3, t.Y3) {
				canvas.Set(x, y, fillColor)
			}
		}
	}
}

// Stroke desenha o contorno do triângulo conforme strokeWeight
func (t *TriangleShape) Stroke(canvas Canvas, strokeColor color.Color, strokeEnabled bool, strokeWeight float64) {
	if !strokeEnabled {
		return
	}
	
	// Verificação para triângulos degenerados (linha ou ponto)
	area := 0.5 * math.Abs((t.X1*(t.Y2-t.Y3) + t.X2*(t.Y3-t.Y1) + t.X3*(t.Y1-t.Y2)))
	if area < 0.01 { // Área muito pequena, triângulo degenerado
		// Se for um triângulo degenerado, mas os pontos são os mesmos, é um ponto
		if math.Abs(t.X1-t.X2) < 0.1 && math.Abs(t.X1-t.X3) < 0.1 && 
		   math.Abs(t.Y1-t.Y2) < 0.1 && math.Abs(t.Y1-t.Y3) < 0.1 {
			// Desenha um ponto
			for dx := -int(strokeWeight / 2); dx <= int(strokeWeight/2); dx++ {
				for dy := -int(strokeWeight / 2); dy <= int(strokeWeight/2); dy++ {
					canvas.Set(int(t.X1)+dx, int(t.Y1)+dy, strokeColor)
				}
			}
		} else {
			// Desenha uma ou mais linhas, dependendo de quais pontos são diferentes
			if math.Abs(t.X1-t.X2) > 0.1 || math.Abs(t.Y1-t.Y2) > 0.1 {
				line := LineShape{X1: t.X1, Y1: t.Y1, X2: t.X2, Y2: t.Y2}
				line.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
			}
			if math.Abs(t.X2-t.X3) > 0.1 || math.Abs(t.Y2-t.Y3) > 0.1 {
				line := LineShape{X1: t.X2, Y1: t.Y2, X2: t.X3, Y2: t.Y3}
				line.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
			}
			if math.Abs(t.X3-t.X1) > 0.1 || math.Abs(t.Y3-t.Y1) > 0.1 {
				line := LineShape{X1: t.X3, Y1: t.Y3, X2: t.X1, Y2: t.Y1}
				line.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
			}
		}
		return
	}
	
	// Desenha as três linhas que formam o triângulo
	line1 := LineShape{X1: t.X1, Y1: t.Y1, X2: t.X2, Y2: t.Y2}
	line2 := LineShape{X1: t.X2, Y1: t.Y2, X2: t.X3, Y2: t.Y3}
	line3 := LineShape{X1: t.X3, Y1: t.Y3, X2: t.X1, Y2: t.Y1}
	
	line1.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
	line2.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
	line3.Stroke(canvas, strokeColor, strokeEnabled, strokeWeight)
}

// Draw executa fill e stroke no triângulo
func (t *TriangleShape) Draw(canvas Canvas, fillColor, strokeColor color.Color, fillEnabled, strokeEnabled bool, strokeWeight float64) {
	t.BaseShape.Draw(canvas, t, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
}

// NewTriangle cria uma nova forma de triângulo
func NewTriangle(x1, y1, x2, y2, x3, y3 float64) *TriangleShape {
	return &TriangleShape{X1: x1, Y1: y1, X2: x2, Y2: y2, X3: x3, Y3: y3}
}

// isPointInTriangle verifica se um ponto está dentro de um triângulo usando coordenadas baricêntricas
func isPointInTriangle(px, py, x1, y1, x2, y2, x3, y3 float64) bool {
	// Calcular área total do triângulo usando fórmula da área
	area := 0.5 * math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)))
	
	// Se a área for muito pequena, consideramos que não é um triângulo válido
	if area < 0.01 {
		return false
	}
	
	// Calcular as três áreas parciais
	s1 := 0.5 * math.Abs((px*(y2-y3) + x2*(y3-py) + x3*(py-y2)))
	s2 := 0.5 * math.Abs((x1*(py-y3) + px*(y3-y1) + x3*(y1-py)))
	s3 := 0.5 * math.Abs((x1*(y2-py) + x2*(py-y1) + px*(y1-y2)))
	
	// Verificar se a soma das áreas parciais é igual à área total
	// Usamos uma pequena margem de erro para compensar imprecisões de ponto flutuante
	return math.Abs((s1+s2+s3)-area) < 0.1 * area
} 