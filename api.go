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
	"runtime/debug"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Xistaminose/gosketch/shapes"
)

// Canvas interno: contexto de desenho
type Canvas struct {
	Width, Height int
	img           *ebiten.Image
}

// Set implements the shapes.Canvas interface
func (c *Canvas) Set(x, y int, clr color.Color) {
	// Verifica se está dentro dos limites do canvas
	if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
		c.img.Set(x, y, clr)
	}
}

// GetWidth returns the canvas width - implements shapes.Canvas interface
func (c *Canvas) GetWidth() int {
	return c.Width
}

// GetHeight returns the canvas height - implements shapes.Canvas interface
func (c *Canvas) GetHeight() int {
	return c.Height
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
	lastFrameTime time.Time   = time.Now()
	frameRate     float64     = 0
	errorHandler  func(error) = defaultErrorHandler
)

// defaultErrorHandler é o tratador de erros padrão que registra o erro e o stack trace
func defaultErrorHandler(err error) {
	log.Printf("ERRO NA GOSKETCH: %v\n", err)
	log.Printf("Stack trace:\n%s\n", debug.Stack())
}

// SetErrorHandler permite definir um tratador de erros personalizado
func SetErrorHandler(handler func(error)) {
	if handler != nil {
		errorHandler = handler
	} else {
		errorHandler = defaultErrorHandler
	}
	
	// Também configura o error handler para cores
	SetColorErrorReporter(errorHandler)
}

// reportError reporta um erro usando o errorHandler atual
func reportError(err error) {
	if errorHandler != nil {
		errorHandler(err)
	}
}

// init inicializa a biblioteca
func init() {
	// Configura o error reporter para cores
	SetColorErrorReporter(reportError)
}

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
	if w <= 0 || h <= 0 {
		reportError(fmt.Errorf("dimensões de canvas inválidas: %dx%d - as dimensões devem ser positivas", w, h))
		w = 100
		h = 100
	}
	canvas = &Canvas{Width: w, Height: h, img: ebiten.NewImage(w, h)}
}

// Background preenche todo o canvas com a cor especificada
func Background(c ColorValue) {
	if canvas != nil {
		color := ParseColorValue(c)
		canvas.img.Fill(color)
	} else {
		reportError(fmt.Errorf("tentativa de definir background sem canvas inicializado"))
	}
}

// Background com valor de cinza (0-255)
func BackgroundGray(gray uint8) {
	Background(Color(gray))
}

// BackgroundColor permite usar color.Color diretamente para manter compatibilidade
func BackgroundColor(c color.Color) {
	Background(ColorFrom(c))
}

// Fill define a cor de preenchimento para formas subsequentes
func Fill(c ColorValue) { 
	fillColor = ParseColorValue(c)
	fillEnabled = true 
}

// FillGray define uma cor de preenchimento em escala de cinza (0-255)
func FillGray(gray uint8) {
	Fill(Color(gray))
}

// FillColor permite usar color.Color diretamente para manter compatibilidade
func FillColor(c color.Color) {
	Fill(ColorFrom(c))
}

// NoFill desabilita preenchimento
func NoFill() { fillEnabled = false }

// Stroke define a cor de contorno para formas subsequentes
func Stroke(c ColorValue) { 
	strokeColor = ParseColorValue(c) 
	strokeEnabled = true 
}

// StrokeGray define uma cor de contorno em escala de cinza (0-255)
func StrokeGray(gray uint8) {
	Stroke(Color(gray))
}

// StrokeColor permite usar color.Color diretamente para manter compatibilidade
func StrokeColor(c color.Color) {
	Stroke(ColorFrom(c))
}

// NoStroke desabilita contorno
func NoStroke() { strokeEnabled = false }

// StrokeWeight define a espessura do contorno para formas subsequentes
func StrokeWeight(w float64) {
	if w < 0 {
		reportError(fmt.Errorf("espessura de contorno inválida: %.2f - deve ser não-negativa", w))
		w = 1
	}
	strokeWeight = w
}

// GetFrameRate retorna o frame rate atual (frames por segundo)
func GetFrameRate() float64 {
	return frameRate
}

// RenderShape executa o método Draw de qualquer Shape da nova API
func RenderShape(s shapes.Shape) {
	if s == nil {
		reportError(fmt.Errorf("tentativa de renderizar uma shape nula"))
		return
	}
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de renderizar sem canvas inicializado"))
		return
	}
	
	// Identifica o tipo de shape para relatórios de desempenho
	shapeName := fmt.Sprintf("%T", s)
	
	// Define um timeout para detecção de shapes lentas
	done := make(chan bool, 1)
	var renderErr error
	
	// Executa o desenho em uma goroutine
	go func() {
		// Captura e reporta possíveis pânicos durante o desenho
		defer func() {
			if r := recover(); r != nil {
				renderErr = fmt.Errorf("pânico durante renderização de %s: %v", shapeName, r)
			}
			done <- true
		}()
		
		s.Draw(canvas, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
	}()
	
	// Espera o desenho completar ou atingir timeout
	select {
	case <-done:
		// Renderização completou normalmente
		if renderErr != nil {
			reportError(renderErr)
		}
	case <-time.After(100 * time.Millisecond): // Ajuste este timeout conforme necessário
		// Timeout - a renderização está demorando demais
		reportError(fmt.Errorf("timeout durante renderização de %s - possível bloqueio", shapeName))
		// Continuamos a execução, não podemos cancelar a goroutine diretamente
		// mas podemos reportar o problema
		
		// Opcional: pegar alguma informação sobre a shape
		shapeInfo := fmt.Sprintf("%+v", s)
		reportError(fmt.Errorf("detalhes da shape problemática: %s", shapeInfo))
	}
}

// Run inicia o loop principal da janela Ebiten
func Run() error {
	// Executa o setup antes de iniciar o loop
	if setupFn != nil {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("pânico durante setup: %v", r)
				reportError(err)
				return
			}
		}()
		
		setupFn()
		setupFn = nil
	}
	
	if canvas == nil {
		err := fmt.Errorf("canvas não criado. Use CreateCanvas no setup")
		reportError(err)
		return err
	}
	
	ebiten.SetWindowSize(canvas.Width, canvas.Height)
	ebiten.SetWindowTitle("Arte Generativa (Go + p5.js API)")
	return ebiten.RunGame(&internalGame{})
}

// internalGame implementa ebiten.Game chamando drawFn e exibindo o canvas
type internalGame struct{}

func (g *internalGame) Update() error {
	// Calcula o frame rate
	now := time.Now()
	elapsed := now.Sub(lastFrameTime).Seconds()
	lastFrameTime = now
	if elapsed > 0 {
		frameRate = 1.0 / elapsed
	}
	
	return nil
}

func (g *internalGame) Draw(screen *ebiten.Image) {
	if drawFn != nil {
		// Captura e reporta possíveis pânicos durante o desenho
		defer func() {
			if r := recover(); r != nil {
				reportError(fmt.Errorf("pânico durante draw: %v", r))
			}
		}()
		
		drawFn()
	}
	
	if canvas != nil {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(canvas.img, op)
	}
}

func (g *internalGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	if canvas == nil {
		reportError(fmt.Errorf("canvas não inicializado ao definir layout"))
		return 300, 300 // valor padrão em caso de erro
	}
	return canvas.Width, canvas.Height
}

// ======= FUNÇÕES DE CONVENIÊNCIA =======

// Ellipse cria e renderiza uma elipse em um único passo
func Ellipse(x, y, rx, ry float64) {
	if rx <= 0 || ry <= 0 {
		reportError(fmt.Errorf("raio inválido para elipse: rx=%.2f, ry=%.2f - os raios devem ser positivos", rx, ry))
		return
	}
	shape := shapes.CreateEllipse(x, y, rx, ry)
	RenderShape(shape)
}

// Circle cria e renderiza um círculo em um único passo (caso especial de Ellipse)
func Circle(x, y, radius float64) {
	if radius <= 0 {
		reportError(fmt.Errorf("raio inválido para círculo: %.2f - o raio deve ser positivo", radius))
		return
	}
	shape := shapes.CreateCircle(x, y, radius)
	RenderShape(shape)
}

// Rectangle cria e renderiza um retângulo em um único passo
func Rectangle(x, y, w, h float64) {
	if w <= 0 || h <= 0 {
		reportError(fmt.Errorf("dimensões inválidas para retângulo: w=%.2f, h=%.2f - as dimensões devem ser positivas", w, h))
		return
	}
	shape := shapes.CreateRectangle(x, y, w, h)
	RenderShape(shape)
}

// Square cria e renderiza um quadrado em um único passo
func Square(x, y, size float64) {
	if size <= 0 {
		reportError(fmt.Errorf("tamanho inválido para quadrado: %.2f - o tamanho deve ser positivo", size))
		return
	}
	shape := shapes.CreateSquare(x, y, size)
	RenderShape(shape)
}

// Line cria e renderiza uma linha em um único passo
func Line(x1, y1, x2, y2 float64) {
	shape := shapes.CreateLine(x1, y1, x2, y2)
	RenderShape(shape)
}

// Point cria e renderiza um ponto em um único passo
func Point(x, y float64) {
	shape := shapes.CreatePoint(x, y)
	RenderShape(shape)
}

// Triangle cria e renderiza um triângulo em um único passo
func Triangle(x1, y1, x2, y2, x3, y3 float64) {
	shape := shapes.CreateTriangle(x1, y1, x2, y2, x3, y3)
	RenderShape(shape)
}
