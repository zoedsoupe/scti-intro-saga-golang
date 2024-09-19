# Minicurso: Introdução ao Padrão SAGA com GoLang

Este repositório contém os materiais e o código-fonte do minicurso **Introdução ao Padrão SAGA com GoLang**. O curso aborda conceitos básicos sobre o padrão SAGA, suas diferentes estratégias de coordenação (Orquestrador e Coreografia) e exemplos práticos implementados em Go.

## Estrutura do Repositório

```bash
.
├── assets                   # Recursos gráficos para a apresentação
│   ├── coreograph-diagram.png    # Diagrama de Coreografia em Mermaid
│   └── orchestrator-diagram.png  # Diagrama de Orquestrador em Mermaid
├── flake.nix                # Configuração do ambiente com Nix Flakes
├── presentation.md          # Apresentação do minicurso em formato Markdown
└── source                   # Código-fonte do projeto em Go
    ├── go.mod               # Arquivo de configuração de módulos do Go
    ├── introduction.go      # Exemplo introdutório de uma transação SAGA
    ├── main.go              # Código principal que coordena os exemplos
    ├── saga.go              # Implementação do gerenciador SAGA
    └── steps.go             # Etapas de uma transação SAGA com funções de reserva e compensação
```

## Conteúdo

### 1. Apresentação
A apresentação em formato Markdown (`presentation.md`) inclui uma introdução ao padrão SAGA, exemplos práticos em Go, e desafios para fixação do conteúdo. A apresentação abrange:

- Conceitos básicos do padrão SAGA.
- Implementação de transações SAGA.
- Estratégias de coordenação: Orquestrador e Coreografia.
- Aplicações práticas e discussão.

### 2. Código-fonte
A pasta `source` contém o código-fonte utilizado durante o minicurso. Inclui a implementação básica de transações SAGA e um gerenciador de SAGA em Go.

### 3. Diagramas
Os diagramas das estratégias de coordenação (`orchestrator-diagram.png` e `coreograph-diagram.png`) estão na pasta `assets` e ilustram visualmente as duas abordagens principais de implementação do padrão SAGA.

## Como Executar

1. **Instalação de Dependências**
   Se você estiver usando o Nix, pode rodar o projeto com o seguinte comando:
   ```bash
   nix develop
   ```
   Caso contrário, certifique-se de ter o Go instalado e execute:
   ```bash
   go mod tidy
   ```

2. **Executando o Projeto**
   Para executar o código principal e visualizar os exemplos de transação SAGA:
   ```bash
   go run source/main.go
   ```

3. **Explorando os Exemplos**
   Edite os arquivos na pasta `source` para explorar diferentes cenários de transações SAGA e suas respectivas compensações.

## Referências

- [Original SAGA Paper (1987)](https://dl.acm.org/doi/10.1145/38713.38742)
- [Padrão SAGA - Wikipedia](https://en.wikipedia.org/wiki/Saga_pattern)
- [Microservices Patterns Book](https://microservices.io/patterns/data/saga.html)
- [Go by Example](https://gobyexample.com/)

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir um issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
