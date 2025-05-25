/*
Projeto: GoSketch - Funções de Imagem
Descrição: Implementação de funções para manipulação de imagens inspiradas em p5.js
Inclui: loadImage(), image(), getPixel(), setPixel(), text() e pixels[]
*/

package gosketch

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// Estrutura para representar uma imagem carregada
type SketchImage struct {
	img    *ebiten.Image
	width  int
	height int
}

// Variáveis globais para gerenciamento de imagens e texto
var (
	loadedImages map[string]*SketchImage = make(map[string]*SketchImage)
	currentFont  font.Face               = basicfont.Face7x13
	textSize     float64                 = 12
	textColor    color.Color             = color.Black
	pixels       [][]color.Color         // Array 2D para acesso aos pixels
)

// LoadImage carrega uma imagem do sistema de arquivos
func LoadImage(path string) *SketchImage {
	// Verifica se a imagem já foi carregada
	if img, exists := loadedImages[path]; exists {
		return img
	}

	// Abre o arquivo
	file, err := os.Open(path)
	if err != nil {
		reportError(fmt.Errorf("erro ao abrir imagem '%s': %v", path, err))
		return nil
	}
	defer file.Close()

	// Decodifica a imagem baseado na extensão
	var img image.Image
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".png":
		img, err = png.Decode(file)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	default:
		// Tenta decodificar automaticamente
		img, _, err = image.Decode(file)
	}

	if err != nil {
		reportError(fmt.Errorf("erro ao decodificar imagem '%s': %v", path, err))
		return nil
	}

	// Converte para ebiten.Image
	ebitenImg := ebiten.NewImageFromImage(img)
	bounds := img.Bounds()

	sketchImg := &SketchImage{
		img:    ebitenImg,
		width:  bounds.Dx(),
		height: bounds.Dy(),
	}

	// Armazena no cache
	loadedImages[path] = sketchImg

	return sketchImg
}

// Image desenha uma imagem no canvas
func Image(img *SketchImage, x, y float64, dimensions ...float64) {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de desenhar imagem sem canvas inicializado"))
		return
	}

	if img == nil {
		reportError(fmt.Errorf("tentativa de desenhar imagem nula"))
		return
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)

	// Se largura e altura foram especificadas, aplica escala
	if len(dimensions) >= 2 {
		scaleX := dimensions[0] / float64(img.width)
		scaleY := dimensions[1] / float64(img.height)
		opts.GeoM.Scale(scaleX, scaleY)
	} else if len(dimensions) == 1 {
		// Apenas largura especificada, mantém proporção
		scale := dimensions[0] / float64(img.width)
		opts.GeoM.Scale(scale, scale)
	}

	canvas.img.DrawImage(img.img, opts)
}

// GetPixel retorna a cor de um pixel específico do canvas
func GetPixel(x, y int) color.Color {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de obter pixel sem canvas inicializado"))
		return color.Black
	}

	if x < 0 || x >= canvas.Width || y < 0 || y >= canvas.Height {
		reportError(fmt.Errorf("coordenadas de pixel fora dos limites: (%d, %d)", x, y))
		return color.Black
	}

	return canvas.img.At(x, y)
}

// SetPixel define a cor de um pixel específico do canvas
func SetPixel(x, y int, c ColorValue) {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de definir pixel sem canvas inicializado"))
		return
	}

	if x < 0 || x >= canvas.Width || y < 0 || y >= canvas.Height {
		reportError(fmt.Errorf("coordenadas de pixel fora dos limites: (%d, %d)", x, y))
		return
	}

	color := ParseColorValue(c)
	canvas.img.Set(x, y, color)
}

// LoadPixels carrega os pixels do canvas para o array pixels[]
func LoadPixels() {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de carregar pixels sem canvas inicializado"))
		return
	}

	// Inicializa o array de pixels
	pixels = make([][]color.Color, canvas.Height)
	for y := 0; y < canvas.Height; y++ {
		pixels[y] = make([]color.Color, canvas.Width)
		for x := 0; x < canvas.Width; x++ {
			pixels[y][x] = canvas.img.At(x, y)
		}
	}
}

// UpdatePixels aplica as mudanças do array pixels[] de volta ao canvas
func UpdatePixels() {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de atualizar pixels sem canvas inicializado"))
		return
	}

	if pixels == nil {
		reportError(fmt.Errorf("tentativa de atualizar pixels sem carregar pixels primeiro"))
		return
	}

	for y := 0; y < len(pixels) && y < canvas.Height; y++ {
		for x := 0; x < len(pixels[y]) && x < canvas.Width; x++ {
			canvas.img.Set(x, y, pixels[y][x])
		}
	}
}

// GetPixels retorna uma referência ao array de pixels (somente leitura)
func GetPixels() [][]color.Color {
	return pixels
}

// Text desenha texto no canvas
func Text(str string, x, y float64) {
	if canvas == nil {
		reportError(fmt.Errorf("tentativa de desenhar texto sem canvas inicializado"))
		return
	}

	text.Draw(canvas.img, str, currentFont, int(x), int(y), textColor)
}

// TextSize define o tamanho do texto (funcionalidade limitada com basicfont)
func TextSize(size float64) {
	if size <= 0 {
		reportError(fmt.Errorf("tamanho de texto inválido: %.2f - deve ser positivo", size))
		return
	}
	textSize = size
	// Nota: basicfont tem tamanho fixo, esta função é para compatibilidade
}

// TextColor define a cor do texto
func TextColor(c ColorValue) {
	textColor = ParseColorValue(c)
}

// TextWidth retorna a largura aproximada de um texto
func TextWidth(str string) int {
	if currentFont == nil {
		return len(str) * 7 // Estimativa para basicfont
	}

	width := 0
	for _, r := range str {
		advance, ok := currentFont.GlyphAdvance(r)
		if ok {
			width += int(advance >> 6) // Converte de fixed.Int26_6 para int
		} else {
			width += 7 // Largura padrão
		}
	}
	return width
}

// TextHeight retorna a altura do texto atual
func TextHeight() int {
	if currentFont == nil {
		return 13 // Altura padrão do basicfont
	}
	metrics := currentFont.Metrics()
	return int(metrics.Height >> 6) // Converte de fixed.Int26_6 para int
}

// Métodos para SketchImage

// Width retorna a largura da imagem
func (img *SketchImage) Width() int {
	return img.width
}

// Height retorna a altura da imagem
func (img *SketchImage) Height() int {
	return img.height
}

// GetPixel retorna a cor de um pixel específico da imagem
func (img *SketchImage) GetPixel(x, y int) color.Color {
	if x < 0 || x >= img.width || y < 0 || y >= img.height {
		return color.Black
	}
	return img.img.At(x, y)
}

// Copy cria uma cópia da imagem
func (img *SketchImage) Copy() *SketchImage {
	newImg := ebiten.NewImage(img.width, img.height)
	newImg.DrawImage(img.img, nil)

	return &SketchImage{
		img:    newImg,
		width:  img.width,
		height: img.height,
	}
}

// Resize redimensiona a imagem (cria uma nova imagem)
func (img *SketchImage) Resize(newWidth, newHeight int) *SketchImage {
	if newWidth <= 0 || newHeight <= 0 {
		reportError(fmt.Errorf("dimensões de redimensionamento inválidas: %dx%d", newWidth, newHeight))
		return img
	}

	newImg := ebiten.NewImage(newWidth, newHeight)

	opts := &ebiten.DrawImageOptions{}
	scaleX := float64(newWidth) / float64(img.width)
	scaleY := float64(newHeight) / float64(img.height)
	opts.GeoM.Scale(scaleX, scaleY)

	newImg.DrawImage(img.img, opts)

	return &SketchImage{
		img:    newImg,
		width:  newWidth,
		height: newHeight,
	}
}

// CreateImage cria uma nova imagem vazia
func CreateImage(width, height int) *SketchImage {
	if width <= 0 || height <= 0 {
		reportError(fmt.Errorf("dimensões de imagem inválidas: %dx%d", width, height))
		return nil
	}

	img := ebiten.NewImage(width, height)
	return &SketchImage{
		img:    img,
		width:  width,
		height: height,
	}
}

// SaveImage salva o canvas atual como imagem
func SaveImage(filename string) error {
	if canvas == nil {
		return fmt.Errorf("tentativa de salvar imagem sem canvas inicializado")
	}

	// Cria o arquivo
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo '%s': %v", filename, err)
	}
	defer file.Close()

	// Converte ebiten.Image para image.Image
	bounds := image.Rect(0, 0, canvas.Width, canvas.Height)
	img := image.NewRGBA(bounds)

	for y := 0; y < canvas.Height; y++ {
		for x := 0; x < canvas.Width; x++ {
			img.Set(x, y, canvas.img.At(x, y))
		}
	}

	// Salva baseado na extensão
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		return png.Encode(file, img)
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	default:
		return png.Encode(file, img) // Padrão para PNG
	}
}
