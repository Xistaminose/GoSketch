# GoSketch CLI

Uma ferramenta de linha de comando para criação de projetos [GoSketch](https://github.com/Xistaminose/gosketch) de arte generativa.

## Instalação

```bash
go install github.com/Xistaminose/gosketch/cmd/gosketch@latest
```

Certifique-se de que `$GOPATH/bin` está em seu PATH.

## Uso

### Criar um novo projeto

```bash
gosketch new meu-projeto
```

Este comando cria um novo diretório `meu-projeto` contendo:
- `main.go` - Exemplo básico de arte generativa
- `go.mod` - Arquivo de módulo Go configurado

### Executar o projeto

Após criar o projeto:

```bash
cd meu-projeto
go mod tidy
go run main.go
```

### Ajuda

Para ver todas as opções disponíveis:

```bash
gosketch help
```

## Exemplos

### Projeto Básico

```bash
gosketch new circulo-pulsante
cd circulo-pulsante
go mod tidy
go run main.go
```

## Contribuições

Veja as diretrizes de contribuição no [repositório principal](https://github.com/Xistaminose/gosketch). 