# Minicurso: Introdução ao Padrão SAGA com GoLang

## Sobre mim

- **Zoey Pessanha**, desenvolvedora de software.
- Experiência com linguagens funcionais e esotéricas (Elixir 💜).
- Atualmente como Senior na [Cumbuca](https://cumbuca.com)
- Foco em sistemas financeiros, distribuídos e arquitetura resiliente.

---
  
## Agenda

1. Introdução ao Padrão SAGA
2. Implementação de uma Transação SAGA
3. Gerenciador de SAGA
4. Aplicações e Discussão

---
    
## 1. O que é o Padrão SAGA?

- Gerencia **transações distribuídas** sem usar transações ACID.
- Divide uma transação longa em **mini-transações** com operações de **compensação**.
- Mantém a consistência do sistema em caso de falha.
    
> **ACID** é um conjunto de propriedades que garante a confiabilidade das transações do banco de dados:
> **A**tomicidade
> **C**onsistência
> **I**solamento
> **D**urabilidade
    
### História do SAGA

- Introduzido por Hector Garcia-Molina e Kenneth Salem em 1987.
- Criado para resolver transações distribuídas em **bancos de dados** e **sistemas distribuídos**.
- Popularizado com a adoção de **microserviços**.
    
### Referências

- [Original SAGA Paper (1987)](https://dl.acm.org/doi/10.1145/38713.38742)
- [Padrão SAGA - Wikipedia](https://en.wikipedia.org/wiki/Saga_pattern)
- [Microservices Patterns Book](https://microservices.io/patterns/data/saga.html)
- [Understanding SAGA Pattern](https://www.baeldung.com/microservices-saga-pattern)
- [Go by Example](https://gobyexample.com/)

---
    
## 2. Implementação de uma SAGA

### Introdução ao Padrão SAGA

**Contexto Lúdico**: Imagine que você está organizando uma viagem de férias para um grupo de amigos. Cada etapa da reserva (voo, hotel, carro) precisa ser coordenada. Se uma falha ocorrer em qualquer etapa, você deve desfazer as reservas anteriores para manter tudo em ordem.
- [Playground](https://onecompiler.com/go/42smb65ht)
    
#### Exemplo Simples

```go
package main

import "fmt"

// Função para reservar voo
func reserveFlight() error {
    fmt.Println("Reserva de voo realizada. 💜")
    return nil
}

// Função para cancelar voo
func cancelFlight() {
    fmt.Println("Reserva de voo cancelada. 😭💔")
}

func main() {
    err := reserveFlight()

    if err != nil {
        cancelFlight() // Operação compensatória
    }
}
```

---
    
### Simulação de um Erro Durante a Transação

```go
package main

import (
    "fmt"
    "errors"
)

// Função para reservar voo com erro simulado
func reserveFlight() error {
    fmt.Println("Reserva de voo realizada. 💜")
    return errors.New("DEU RUIM")
}

// Função para cancelar voo
func cancelFlight() {
    fmt.Println("Reserva de voo cancelada. 😭💔")
}

func main() {
    err := reserveFlight()

    if err != nil {
        cancelFlight() // Operação compensatória
    }
}
```

---
    
### Desafio 1:

- **Contexto**: Continuando com a viagem de férias, você quer garantir que, se a reserva do voo falhar, o sistema notifique o usuário sobre a falha.
- **Tarefa**: 
  - Adicione uma mensagem de erro no `main` para informar o usuário quando a reserva do voo falhar.

---
    
### Transação SAGA com Múltiplas Etapas

1. **Reserve voo**
2. **Reserve hotel**
3. **Compensar em caso de falha**
    
### Exemplo Básico

- [Playgroud](https://onecompiler.com/go/42smbp49x)
    
```go
package main

import "fmt"

// Funções de reserva e cancelamento
func reserveFlight() error {
    fmt.Println("Reserva de voo realizada. 💜")
    return nil
}

func cancelFlight() {
    fmt.Println("Reserva de voo cancelada. 😭💔")
}

func reserveHotel() error {
    fmt.Println("Reserva de hotel realizada. 💜")
    return nil
}

func cancelHotel() {
    fmt.Println("Reserva de hotel cancelada. 😭💔")
}
```

---

```go
func main() {
    err := reserveFlight()
    if err != nil {
        cancelFlight()
        return
    }

    err = reserveHotel()
    if err != nil {
        cancelFlight()
        cancelHotel()
        return
    }

    fmt.Println("Todas as reservas realizadas com sucesso! 🎉")
}
```

---
    
### Desafio 2:

- **Contexto**: Além de voo e hotel, seu grupo também quer reservar um passeio turístico. Se a reserva do passeio falhar, todas as reservas anteriores devem ser canceladas.
- **Tarefa**:
  - Insira um novo serviço na cadeia, como a "Reserva de passeio turístico".
  - Implemente a lógica de sucesso e falha para esse serviço.

---
    
## 3. Gerenciador de SAGA

### Por que usar um Gerenciador?

- **Encapsular operações** e suas compensações.
- Centralizar o **tratamento de erros**.
- Facilitar o **reuso**.
    
### Estrutura Básica

- **Lista de passos**: ações da transação.
- **Lista de compensações**: desfaz ações em caso de erro.

---
    
### Exemplo de Gerenciador

- [Playground](https://onecompiler.com/go/42smbr28c)

```go
// Estrutura do SAGA
type Saga struct {
    steps         []func() error
    compensations []func()
}

// Adiciona um passo e sua compensação
func (s *Saga) AddStep(step func() error, compensation func()) {
    s.steps = append(s.steps, step)
    s.compensations = append(s.compensations, compensation)
}

// Executa o SAGA
func (s *Saga) Execute() error {
    for i, step := range s.steps {
        if err := step(); err != nil {
            // Executa as compensações em caso de erro
            for j := i - 1; j >= 0; j-- {
                s.compensations[j]()
            }
            return err
        }
    }
    return nil
}
```

---

```go
func main() {
    saga := &Saga{}

    saga.AddStep(reserveFlight, cancelFlight)
    saga.AddStep(reserveHotel, cancelHotel)

    if err := saga.Execute(); err != nil {
        fmt.Println("Transação falhou:", err)
    } else {
        fmt.Println("Todas as reservas realizadas com sucesso! 🎉")
    }
}
```

---
    
### Desafio 3:

- **Contexto**: O gerenciador de SAGA precisa ser mais funcional e facilitar a identificação de quais etapas foram realizadas.
- **Tarefa**:
  - Adicione mensagens de log que indiquem quando cada etapa começa e termina.
  - Garanta que, ao executar a SAGA, cada passo e sua compensação sejam claramente indicados no console.

---

## 4. Estratégias de Coordenação do Padrão SAGA

### Orquestração

- **Definição**: Um componente central (orquestrador) coordena todas as etapas da transação, enviando comandos para os serviços participantes e gerenciando as compensações em caso de falha.
  
- **Características**:
  - Centralizado.
  - Facilita o monitoramento e a gestão das transações.
  
- **Exemplo Lúdico**: O organizador da viagem que coordena cada reserva e decide quando cancelar as reservas caso algo dê errado.

---

### Coreografia

- **Definição**: Não há um componente central. Cada serviço participa observando eventos e reagindo de acordo, coordenando a transação através de eventos publicados e assinados.
  
- **Características**:
  - Descentralizado.
  - Menor acoplamento entre serviços.
  
- **Exemplo Lúdico**: Cada amigo responsável por uma reserva na viagem reage aos eventos dos outros amigos, sem um coordenador central.

---
    
## 5. Aplicações do Padrão SAGA

### Onde usar?

- **E-commerce**: Coordenação de pagamento, estoque, envio.
- **Sistemas financeiros**: Transações distribuídas entre contas/bancos.
- **Microserviços**: Garantir consistência sem transações ACID.
    
### Discussão

- **Como você aplicaria o padrão SAGA no seu trabalho?**
- **Qual cenário parece mais relevante?**
    
#### Desafio 4:

- **Contexto**: Pense em um cenário do seu ambiente de trabalho ou projeto pessoal onde o padrão SAGA pode ser aplicado.
- **Tarefa**:
  - Desenhe a sequência de ações e compensações para esse cenário.
  - Implemente uma versão simplificada do fluxo SAGA para esse caso.
    
---
    
## Links úteis

- [Padrão SAGA - Wikipedia](https://en.wikipedia.org/wiki/Saga_pattern)
- [Microservices.io - SAGA](https://microservices.io/patterns/data/saga.html)
- [Golang Documentation](https://golang.org/doc/)
- [Original SAGA Paper (1987)](https://dl.acm.org/doi/10.1145/38713.38742)
- [Understanding SAGA Pattern](https://www.baeldung.com/microservices-saga-pattern)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://play.golang.org/)

---
    
## Agradecimentos

Obrigada pela participação!  
- [LinkedIn](https://linkedin.com/in/zoedsoupe)
- [GitHub](https://github.com/zoedsoupe)
- [BlueSky](https://bsky.app/profile/zoedsoupe.zeetech.io)
