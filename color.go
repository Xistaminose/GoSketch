package gosketch

import (
	"fmt"
	"image/color"
)

// ErrorReporter is a function type used to report errors
type ErrorReporter func(err error)

// Default implementation that simply prints to stdout
var errorReporter ErrorReporter = func(err error) {
	fmt.Printf("ERRO: %v\n", err)
}

// SetColorErrorReporter sets the error reporter function for color-related errors
func SetColorErrorReporter(reporter ErrorReporter) {
	if reporter != nil {
		errorReporter = reporter
	}
}

// ColorValue é um tipo que pode representar uma cor de várias formas:
// - um valor color.Color (ex: color.RGBA{255, 0, 0, 255})
// - um valor uint8/int (0-255) para escala de cinza (ex: 128 para cinza médio)
type ColorValue struct {
	value interface{}
}

// ColorFrom cria um ColorValue a partir de uma color.Color
func ColorFrom(c color.Color) ColorValue {
	return ColorValue{value: c}
}

// GrayFrom cria um ColorValue a partir de um valor uint8 (0-255) para escala de cinza
func GrayFrom(gray uint8) ColorValue {
	return ColorValue{value: gray}
}

// ParseColorValue converte um ColorValue para color.Color
func ParseColorValue(c ColorValue) color.Color {
	switch v := c.value.(type) {
	case color.Color:
		return v
	case uint8:
		return color.RGBA{R: v, G: v, B: v, A: 255}
	case int:
		// Converte int para uint8 se estiver no range correto
		if v < 0 || v > 255 {
			// Para erro em valores fora do range, retorna branco
			errorReporter(fmt.Errorf("valor int fora do range para cor: %d - deve estar entre 0-255", v))
			return color.White
		}
		return color.RGBA{R: uint8(v), G: uint8(v), B: uint8(v), A: 255}
	default:
		// Para tipos inválidos, retorna branco
		errorReporter(fmt.Errorf("tipo de cor inválido dentro de ColorValue: %T", c.value))
		return color.White
	}
}

// RGB creates a color.Color from red, green, and blue uint8 values (0-255)
// with alpha set to 255 (fully opaque).
// Example: RGB(255, 0, 0) // Bright red
func RGB(r, g, b uint8) ColorValue {
	return ColorValue{value: color.RGBA{R: r, G: g, B: b, A: 255}}
}

// RGBA creates a color.Color from red, green, blue, and alpha uint8 values (0-255).
// Alpha controls transparency (0 = fully transparent, 255 = fully opaque).
// Example: RGBA(255, 0, 0, 128) // Semi-transparent red
func RGBA(r, g, b, a uint8) ColorValue {
	return ColorValue{value: color.RGBA{R: r, G: g, B: b, A: a}}
}

// Color creates a ColorValue from a single uint8 value (grayscale)
// with alpha set to 255 (fully opaque).
// 0 = black, 255 = white, values in between = shades of gray.
// Example: Color(128) // Medium gray
func Color(gray uint8) ColorValue {
	return ColorValue{value: gray}
}

// ColorA creates a ColorValue from a grayscale and alpha uint8 value.
// First parameter controls brightness (0 = black, 255 = white).
// Second parameter controls transparency (0 = fully transparent, 255 = fully opaque).
// Example: ColorA(200, 150) // Light gray with some transparency
func ColorA(gray, a uint8) ColorValue {
	return ColorValue{value: color.RGBA{R: gray, G: gray, B: gray, A: a}}
} 