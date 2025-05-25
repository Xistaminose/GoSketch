/*
Projeto: GoSketch
Descrição: Biblioteca de arte generativa 2D em Go inspirada em p5.js/Processing.
Oferece API simples e intuitiva para criação de sketches interativos com formas geométricas,
cores, animações e interatividade. Ideal para artistas e programadores explorarem arte
computacional e visualização criativa.
Adicionado: 
- suporte a fill e interface Shape para formas extensíveis.
- funções de controle de loop: NoLoop(), Loop() e Redraw(n) para controlar animações.
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
	isLooping     bool        = true  // Controla se o draw loop está ativo ou pausado
	redrawCount   int         = 0     // Contador para múltiplas execuções do draw quando solicitado
    targetFPS    int         = 60
    maxFPS       float64     = 240.0
    lastDrawTime time.Time   = time.Now()
    minFrameTime float64     = 0.0001
    sketchStartTime time.Time = time.Now() // Armazena o momento de início do sketch

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
	
	// Captura e reporta possíveis pânicos durante o desenho
	defer func() {
		if r := recover(); r != nil {
			reportError(fmt.Errorf("pânico durante renderização: %v", r))
		}
	}()
	
	s.Draw(canvas, fillColor, strokeColor, fillEnabled, strokeEnabled, strokeWeight)
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
	now := time.Now()
	elapsed := now.Sub(lastFrameTime).Seconds()
	lastFrameTime = now

	// Protege contra valores irreais
	if elapsed < minFrameTime {
		elapsed = minFrameTime
	}

	frameRate = 1.0 / elapsed

	// Protege contra estouros absurdos
	if frameRate > maxFPS {
		frameRate = maxFPS
	}

	return nil
}

func (g *internalGame) Draw(screen *ebiten.Image) {
	// Executa o draw se o loop estiver ativo ou se há contagem de redraws
	if drawFn != nil && (isLooping || redrawCount > 0) {
		// Captura e reporta possíveis pânicos durante o desenho
		defer func() {
			if r := recover(); r != nil {
				reportError(fmt.Errorf("pânico durante draw: %v", r))
			}
		}()
		
		// Quando isLooping é falso mas temos redrawCount positivo,
		// executamos a função draw uma vez por frame e decrementamos o contador
		drawFn()
		
		// Reseta redrawCount após uso
		if redrawCount > 0 {
			redrawCount--
		}
	}
	
	if canvas != nil {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(canvas.img, op)
	}
    delta := time.Since(lastDrawTime)
    wait := time.Second/time.Duration(targetFPS) - delta
    if wait > 0 {
        time.Sleep(wait)
    }
    lastDrawTime = time.Now()
}

func (g *internalGame) Layout(outsideWidth, outsideHeight int) (int, int) {
    if canvas == nil {
        reportError(fmt.Errorf("canvas não inicializado ao definir layout"))
		return 300, 300 // valor padrão em caso de erro
	}
	return canvas.Width, canvas.Height
}

// NoLoop para a execução contínua da função draw
// Similar à função noLoop() do p5.js/Processing
func NoLoop() {
    isLooping = false
}

// Loop reinicia a execução contínua da função draw após uma chamada a NoLoop()
// Similar à função loop() do p5.js/Processing
func Loop() {
    isLooping = true
}

// IsLooping retorna se o loop de desenho está ativo ou não
func IsLooping() bool {
    return isLooping
}

// Redraw força uma ou mais atualizações do canvas
// Se n > 1, executa a função draw n vezes consecutivas (uma por frame)
// Útil quando noLoop() está ativo, mas precisa-se atualizar o canvas
// Quando chamado sem argumentos, executa draw() uma única vez
// Quando loop está ativo, essa função não tem efeito
// Similar à função redraw(n) do p5.js/Processing
func Redraw(n ...int) {
    if !isLooping {
        count := 1 // Valor padrão
        if len(n) > 0 && n[0] > 0 {
            count = n[0]
        }
        redrawCount = count
    }
}

// RedrawOnce força uma única atualização do canvas
// Útil quando noLoop() está ativo, mas precisa-se atualizar o canvas uma vez
// Equivalente a Redraw(1)
// Deprecated: Use Redraw() instead
func RedrawOnce() {
    Redraw(1)
}

// FrameRate define o número de quadros por segundo desejado.
// Se fps <= 0, libera para rodar o mais rápido possível.
func FrameRate(fps int) {
    targetFPS = fps
    if fps > 0 {
        ebiten.SetVsyncEnabled(false)
        ebiten.SetTPS(fps)
    } else {
        ebiten.SetVsyncEnabled(true)
        ebiten.SetTPS(ebiten.SyncWithFPS)
    }
}

// Millis retorna o número de milissegundos desde que o sketch começou a ser executado
// Similar à função millis() do p5.js/Processing
func Millis() int {
	return int(time.Since(sketchStartTime).Milliseconds())
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

// GetWidth retorna a largura do canvas atual
func GetWidth() int {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de obter largura sem canvas inicializado"))
		return 0
	}
	return canvas.Width
}

// GetHeight retorna a altura do canvas atual
func GetHeight() int {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de obter altura sem canvas inicializado"))
		return 0
	}
	return canvas.Height
}
