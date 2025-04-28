# GoSketch

Biblioteca de arte generativa 2D em Go, inspirada em p5.js/Processing, com API simples para criação de sketches interativos.

## 📦 Instalação

```bash
go get github.com/Xistaminose/GoSketch
```

Ou, via módulo:

```bash
go mod init github.com/seuusuario/seurepo
go get github.com/Xistaminose/GoSketch@latest
```

## 🚀 Começando

Crie um arquivo `main.go` com o seguinte conteúdo:

```go
package main

import (
  "image/color"
  "github.com/Xistaminose/GoSketch/sketch"  
)

func setup() {
  sketch.CreateCanvas(400, 400)
  sketch.Fill(color.RGBA{255, 100, 100, 255})
  sketch.Stroke(color.RGBA{0, 0, 0, 255})
  sketch.StrokeWeight(2)
}

func draw() {
  sketch.Background(color.RGBA{220, 220, 220, 255})
  sketch.NoStroke()
  sketch.RenderShape(&sketch.EllipseShape{X: 200, Y: 200, Rx: 80, Ry: 50})
}

func main() {
  sketch.Setup(setup)
  sketch.Draw(draw)
  sketch.Run()
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

### Versão 1 – Fundamentos

- point(), line(), triangle(), rect(), ellipse()
- fill(), stroke(), strokeWeight(), background()
- noFill(), noStroke(), width, height
- setup(), draw(), sin(), cos(), radians(), random()
- frameRate(), loop(), noLoop(), millis()

### Versão 2 – Intermediário

- beginShape(), vertex(), endShape(), curve(), bezier()
- translate(), rotate(), scale(), pushMatrix(), popMatrix()
- noise(), colorMode(), strokeJoin(), strokeCap()

### Versão 3 – Avançado

- text(), loadImage(), image(), getPixel(), setPixel()
- pixels[], exp(), pow(), sqrt(), smooth()
- eventos: mousePressed(), keyPressed()

## 🤝 Contribuição

1. Fork este repositório
2. Crie uma branch feature: `git checkout -b feature/nome-da-funcao`
3. Commit suas alterações: `git commit -m "Adiciona <funcionalidade>"`
4. Envie para seu fork: `git push origin feature/nome-da-funcao`
5. Abra um Pull Request
