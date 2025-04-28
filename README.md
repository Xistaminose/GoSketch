# GoSketch

Biblioteca de arte generativa 2D em Go, inspirada em p5.js/Processing, com API simples para criação de sketches interativos.

## 📦 Instalação

```bash
go get github.com/Xistaminose/gosketch
```

Ou, via módulo:

```bash
go mod init github.com/seuusuario/seurepo
go get github.com/Xistaminose/gosketch@latest
```

## 🚀 Começando

### Usando a CLI (Recomendado)

A maneira mais fácil de começar é usar nossa ferramenta de linha de comando:

```bash
# Instale a CLI
go install github.com/Xistaminose/gosketch/cmd/gosketch@latest

# Crie um novo projeto
gosketch new meu-primeiro-sketch

# Entre no diretório e execute
cd meu-primeiro-sketch
go mod tidy
go run main.go
```

### Manualmente

Crie um arquivo `main.go` com o seguinte conteúdo:

```go
package main

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

```

Execute:

```bash
go run main.go
```

## 📖 API Básica

- **Setup(func())**: registra função de inicialização
- **Draw(func())**: registra função de desenho
- **CreateCanvas(w, h int)**: define tamanho do canvas
- **Background(c color.Color)**: cor de fundo
- **Fill(c color.Color)** / **NoFill()**: cor de preenchimento ou desabilita
- **Stroke(c color.Color)** / **NoStroke()**: cor de contorno ou desabilita
- **StrokeWeight(w float64)**: espessura do traço
- **RenderShape(s Shape)**: desenha qualquer `Shape`
- **Run()**: inicia loop principal e exibe janela

## 📋 Roadmap

### Versão 1 – Fundamentos

- point(), line(), triangle(), rect(), ellipse()
- fill(), stroke(), strokeWeight(), background()
- noFill(), noStroke(), width, height
- setup(), draw(), sin(), cos(), radians(), random()
- frameRate(), loop(), noLoop(), millis()

### Versão 2 – Intermediário

- beginShape(), vertex(), endShape(), curve(), bezier()
- translate(), rotate(), scale(), pushMatrix(), popMatrix()
- noise(), colorMode(), strokeJoin(), strokeCap()

### Versão 3 – Avançado

- text(), loadImage(), image(), getPixel(), setPixel()
- pixels[], exp(), pow(), sqrt(), smooth()
- eventos: mousePressed(), keyPressed()

## 🤝 Contribuição

1. Fork este repositório
2. Crie uma branch feature: `git checkout -b feature/nome-da-funcao`
3. Commit suas alterações: `git commit -m "Adiciona <funcionalidade>"`
4. Envie para seu fork: `git push origin feature/nome-da-funcao`
5. Abra um Pull Request
