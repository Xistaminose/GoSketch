package main

// Templates para diferentes tipos de projetos GoSketch

// Template básico: uma elipse simples
const basicTemplate = `package main

import (
	"github.com/Xistaminose/gosketch"
)

func setup() {
	gosketch.CreateCanvas(400, 400)
}

func draw() {
	gosketch.Background(gosketch.Color(255))
	gosketch.Point(200, 200)
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
	"github.com/Xistaminose/gosketch"
)

func setup() {
	gosketch.CreateCanvas(400, 400)
}

func draw() {
	gosketch.Background(gosketch.RGBA(240, 240, 240, 255))

	// Desenha um retângulo central
	gosketch.Fill(gosketch.RGB(200, 200, 100))
	gosketch.Rectangle(100, 100, 200, 150)

	// Desenha um quadrado menor
	gosketch.Fill(gosketch.RGB(100, 200, 200))
	gosketch.Square(150, 150, 100)
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
	"github.com/Xistaminose/gosketch"
)

func setup() {
	gosketch.CreateCanvas(400, 400)
}

func draw() {
	gosketch.Background(gosketch.Color(255))

	// Desenha uma grade de linhas
	gosketch.Stroke(gosketch.Color(200))
	for i := 0; i < 400; i += 20 {
		// Linhas horizontais
		gosketch.Line(0, float64(i), 400, float64(i))
		// Linhas verticais
		gosketch.Line(float64(i), 0, float64(i), 400)
	}

	// Desenha algumas diagonais
	gosketch.Stroke(gosketch.RGB(255, 100, 100))
	gosketch.StrokeWeight(3)

	gosketch.Line(0, 0, 400, 400)
	gosketch.Line(0, 400, 400, 0)
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