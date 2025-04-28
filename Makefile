.PHONY: build install test clean

BINARY_NAME=gosketch
CLI_PATH=./cmd/gosketch

build:
	@echo "Building $(BINARY_NAME) CLI..."
	go build -o bin/$(BINARY_NAME) $(CLI_PATH)
	@echo "Built at bin/$(BINARY_NAME)"

install:
	@echo "Installing $(BINARY_NAME) CLI..."
	go install $(CLI_PATH)
	@echo "Installed! Make sure your GOPATH/bin is in your PATH."

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/
	
# Create a test project to verify the CLI works
test-cli: build
	@echo "Testing CLI..."
	mkdir -p tmp
	bin/$(BINARY_NAME) new tmp/test-project
	@echo "CLI test completed. Check tmp/test-project"

help:
	@echo "GoSketch - Arte Generativa em Go"
	@echo ""
	@echo "Makefile para desenvolvimento e instalação"
	@echo ""
	@echo "Comandos:"
	@echo "  make build      : Compila o programa para './bin/$(BINARY_NAME)'"
	@echo "  make install    : Instala o CLI no seu GOPATH/bin"
	@echo "  make test       : Executa os testes"
	@echo "  make test-cli   : Testa o CLI criando um projeto de exemplo"
	@echo "  make clean      : Remove arquivos de compilação"
	@echo "  make help       : Mostra esta ajuda"

# Default target
all: build 