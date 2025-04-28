/*
Projeto: GoSketch
Descrição: Biblioteca de arte generativa 2D em Go inspirada em p5.js/Processing.
Oferece API simples e intuitiva para criação de sketches interativos com formas geométricas,
cores, animações e interatividade. Ideal para artistas e programadores explorarem arte
computacional e visualização criativa.
Adicionado: suporte a fill e interface Shape para formas extensíveis.
Dependência mínima: Ebiten para janela e desenho 2D.
*/

package gosketch

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Canvas interno: contexto de desenho
type Canvas struct {
	Width, Height int
	img           *ebiten.Image
}

// Estado global da API
var (
	canvas        *Canvas
	setupFn       func()
	drawFn        func()
	fillColor     color.Color = color.White
	strokeColor   color.Color = color.Black
	strokeWeight  float64     = 1
	fillEnabled   bool        = true
	strokeEnabled bool        = true
)

// Setup registra a função de inicialização (setup)
func Setup(f func()) {
	setupFn = f
}

// Draw registra a função de desenho (draw)
func Draw(f func()) {
	drawFn = f
}

// CreateCanvas define largura e altura do canvas
func CreateCanvas(w, h int) {
	canvas = &Canvas{Width: w, Height: h, img: ebiten.NewImage(w, h)}
}

// Background preenche todo o canvas com a cor especificada
func Background(c color.Color) {
	if canvas != nil {
		canvas.img.Fill(c)
	}
}

// Shape é a interface que todas as formas devem implementar
type Shape interface {
	Draw()
}

// RenderShape executa o método Draw de qualquer Shape
func RenderShape(s Shape) {
	if s == nil {
		log.Println("Shape não definido")
		return
	}
	s.Draw()
}

// Run inicia o loop principal da janela Ebiten
func Run() error {
	// Executa o setup antes de iniciar o loop
	if setupFn != nil {
		setupFn()
		setupFn = nil
	}
	if canvas == nil {
		return fmt.Errorf("Canvas não criado. Use CreateCanvas no setup ")
	}
	ebiten.SetWindowSize(canvas.Width, canvas.Height)
	ebiten.SetWindowTitle("Arte Generativa (Go + p5.js API)")
	return ebiten.RunGame(&internalGame{})
}

// internalGame implementa ebiten.Game chamando drawFn e exibindo o canvas

type internalGame struct{}

func (g *internalGame) Update() error {
	return nil
}

func (g *internalGame) Draw(screen *ebiten.Image) {
	if drawFn != nil {
		drawFn()
	}
	if canvas != nil {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(canvas.img, op)
	}
}

func (g *internalGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return canvas.Width, canvas.Height
}
