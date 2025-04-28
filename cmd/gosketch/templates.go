package main

// Templates para diferentes tipos de projetos GoSketch

// Template básico: uma elipse simples
const basicTemplate = `package main

import (
	"image/color"
	"github.com/Xistaminose/gosketch"  
)

func setup() {
	gosketch.CreateCanvas(400, 400)
	gosketch.Fill(color.RGBA{255, 100, 100, 255})
	gosketch.Stroke(color.RGBA{0, 0, 0, 255})
	gosketch.StrokeWeight(2)
}

func draw() {
	gosketch.Background(color.RGBA{220, 220, 220, 255})
	gosketch.NoStroke()
	gosketch.RenderShape(&gosketch.EllipseShape{X: 200, Y: 200, Rx: 80, Ry: 50})
}

func main() {
	gosketch.Setup(setup)
	gosketch.Draw(draw)
	gosketch.Run()
}
`

// Template com retângulos
const rectangleTemplate = `package main

import (
	"image/color"
	"github.com/Xistaminose/gosketch"  
)

func setup() {
	gosketch.CreateCanvas(400, 400)
	gosketch.Fill(color.RGBA{100, 150, 255, 255})
	gosketch.Stroke(color.RGBA{50, 50, 50, 255})
	gosketch.StrokeWeight(3)
}

func draw() {
	gosketch.Background(color.RGBA{240, 240, 240, 255})
	
	// Desenha um retângulo central
	gosketch.RenderShape(&gosketch.RectangleShape{X: 100, Y: 100, W: 200, H: 150})
	
	// Desenha um quadrado menor
	gosketch.Fill(color.RGBA{255, 200, 100, 255})
	gosketch.RenderShape(&gosketch.SquareShape{X: 150, Y: 150, Size: 100})
}

func main() {
	gosketch.Setup(setup)
	gosketch.Draw(draw)
	gosketch.Run()
}
`

// Template com linhas
const lineTemplate = `package main

import (
	"image/color"
	"github.com/Xistaminose/gosketch"  
)

func setup() {
	gosketch.CreateCanvas(400, 400)
	gosketch.Stroke(color.RGBA{0, 0, 0, 255})
	gosketch.StrokeWeight(2)
}

func draw() {
	gosketch.Background(color.RGBA{250, 250, 250, 255})
	
	// Desenha uma grade de linhas
	gosketch.Stroke(color.RGBA{200, 200, 200, 255})
	for i := 0; i < 400; i += 20 {
		// Linhas horizontais
		gosketch.RenderShape(&gosketch.LineShape{X1: 0, Y1: float64(i), X2: 400, Y2: float64(i)})
		// Linhas verticais
		gosketch.RenderShape(&gosketch.LineShape{X1: float64(i), Y1: 0, X2: float64(i), Y2: 400})
	}
	
	// Desenha algumas diagonais
	gosketch.Stroke(color.RGBA{255, 100, 100, 255})
	gosketch.StrokeWeight(3)
	gosketch.RenderShape(&gosketch.LineShape{X1: 0, Y1: 0, X2: 400, Y2: 400})
	gosketch.RenderShape(&gosketch.LineShape{X1: 0, Y1: 400, X2: 400, Y2: 0})
}

func main() {
	gosketch.Setup(setup)
	gosketch.Draw(draw)
	gosketch.Run()
}
`

// Mapeamento de templates disponíveis
var templates = map[string]string{
	"basic":     basicTemplate,
	"rectangle": rectangleTemplate,
	"line":      lineTemplate,
} 