# GoSketch CLI - Guia de Uso

Este documento explica como usar a ferramenta de linha de comando GoSketch para criar e gerenciar projetos de arte generativa.

## Instalação

Existem duas maneiras de instalar a CLI GoSketch:

### 1. Usando `go install`

```bash
go install github.com/Xistaminose/gosketch/cmd/gosketch@latest
```

Certifique-se de que `$GOPATH/bin` está em seu PATH.

### 2. Compilando manualmente

Clone o repositório e compile:

```bash
git clone https://github.com/Xistaminose/gosketch.git
cd gosketch
make build
```

O binário será gerado em `./bin/gosketch`. Você pode movê-lo para um diretório em seu PATH ou usar o comando `make install`.

## Comandos Básicos

### Listar templates disponíveis

```bash
gosketch list-templates
```

### Criar um novo projeto

Syntax básica:
```bash
gosketch new <nome-do-projeto> [template]
```

Exemplos:

```bash
# Criar projeto com template padrão (basic)
gosketch new meu-sketch

# Criar projeto com template específico
gosketch new meu-sketch-retangulos rectangle

# Criar projeto com template de linhas
gosketch new grade-de-linhas line
```

### Ajuda

```bash
gosketch help
```

## Templates Disponíveis

1. **basic** - Template básico com uma elipse simples (padrão)
2. **rectangle** - Template com retângulos e quadrados
3. **line** - Template com grade de linhas

## Executando seu Projeto

Após criar um projeto, navegue até o diretório e execute:

```bash
cd nome-do-projeto
go mod tidy  # Resolve dependências
go run main.go
```

Isso abrirá uma janela mostrando seu sketch.

## Fluxo de Trabalho Típico

1. Crie um novo projeto com um template: `gosketch new meu-projeto`
2. Navegue para o diretório: `cd meu-projeto`
3. Instale dependências: `go mod tidy`
4. Edite o arquivo `main.go` com seu código personalizado
5. Execute o sketch: `go run main.go`

## Dicas

- Você pode usar subdiretórios ao criar projetos: `gosketch new projetos/sketch-1`
- Modifique as funções `setup()` e `draw()` para personalizar seu sketch
- Consulte a documentação da biblioteca GoSketch para mais informações sobre as formas disponíveis e funções de desenho 