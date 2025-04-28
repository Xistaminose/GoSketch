package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const goModTemplate = `module %s

go 1.22.0

require github.com/Xistaminose/gosketch v0.0.0-latest
`

func printUsage() {
	fmt.Println("GoSketch - Arte Generativa em Go")
	fmt.Println()
	fmt.Println("Uso:")
	fmt.Println("  gosketch new <project-name> [template]    Cria um novo projeto GoSketch")
	fmt.Println("  gosketch list-templates                  Lista templates disponíveis")
	fmt.Println("  gosketch help                           Exibe esta ajuda")
	fmt.Println()
	fmt.Println("Templates disponíveis:")
	fmt.Println("  basic     - Template básico com uma elipse simples (padrão)")
	fmt.Println("  rectangle - Template com retângulos e quadrados")
	fmt.Println("  line      - Template com grade de linhas")
}

func listTemplates() {
	fmt.Println("Templates disponíveis para o GoSketch:")
	fmt.Println()
	fmt.Println("  basic     - Template básico com uma elipse simples (padrão)")
	fmt.Println("  rectangle - Template com retângulos e quadrados")
	fmt.Println("  line      - Template com grade de linhas")
}

func createNewProject(projectName string, templateName string) error {
	// Verificar se o nome do projeto é válido
	if projectName == "" {
		return fmt.Errorf("Nome de projeto não especificado")
	}

	// Verificar template (usar 'basic' se não especificado)
	if templateName == "" {
		templateName = "basic"
	}

	// Verificar se o template existe
	template, ok := templates[templateName]
	if !ok {
		return fmt.Errorf("Template '%s' não encontrado. Use 'gosketch list-templates' para ver opções disponíveis", templateName)
	}

	// Criar diretório para o projeto
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("Erro ao criar diretório do projeto: %v", err)
	}

	// Criar arquivo main.go
	mainPath := filepath.Join(projectName, "main.go")
	err = os.WriteFile(mainPath, []byte(template), 0644)
	if err != nil {
		return fmt.Errorf("Erro ao criar arquivo main.go: %v", err)
	}

	// Criar go.mod
	moduleImportPath := fmt.Sprintf("github.com/%s", projectName)
	if !strings.Contains(projectName, "/") {
		// Se não tiver um caminho completo, usar apenas o nome do projeto
		moduleImportPath = projectName
	}
	goModContent := fmt.Sprintf(goModTemplate, moduleImportPath)
	goModPath := filepath.Join(projectName, "go.mod")
	err = os.WriteFile(goModPath, []byte(goModContent), 0644)
	if err != nil {
		return fmt.Errorf("Erro ao criar arquivo go.mod: %v", err)
	}

	fmt.Printf("Projeto '%s' criado com sucesso com o template '%s'!\n", projectName, templateName)
	fmt.Println()
	fmt.Println("Para executar seu projeto:")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Println("  go mod tidy")
	fmt.Println("  go run main.go")

	return nil
}

func main() {
	// Verificar argumentos
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Erro: Nome do projeto não especificado")
			fmt.Println()
			printUsage()
			os.Exit(1)
		}
		projectName := os.Args[2]
		templateName := ""
		if len(os.Args) >= 4 {
			templateName = os.Args[3]
		}
		err := createNewProject(projectName, templateName)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			os.Exit(1)
		}
	case "list-templates":
		listTemplates()
	case "help":
		printUsage()
	default:
		fmt.Printf("Comando desconhecido: %s\n", command)
		fmt.Println()
		printUsage()
		os.Exit(1)
	}
} 