# Minicurso: Introdução ao Padrão SAGA com GoLang

## Sobre mim

- **Zoey Pessanha**, desenvolvedora de software.
- Experiência com Elixir e GoLang.
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

### História do SAGA

- Introduzido por Hector Garcia-Molina e Kenneth Salem em 1987.
- Criado para resolver transações distribuídas em **bancos de dados** e **sistemas distribuídos**.
- Popularizado com a adoção de **microserviços**.

### Referências

- [Original SAGA Paper (1987)](https://dl.acm.org/doi/10.1145/38713.38742)
- [Padrão SAGA - Wikipedia](https://en.wikipedia.org/wiki/Saga_pattern)
- [Microservices Patterns Book](https://microservices.io/patterns/data/saga.html)

---

## 2. Implementação de uma SAGA

### Introdução ao Padrão SAGA

- Ver [código no Gist](https://gist.github.com/zoedsoupe/c79f9a7c71e4a2c3e216dbc54d22851c#file-introduction_saga-go)
- Exemplo simples onde temos uma transação e uma ação compensatória.

#### Desafio 1:
- Adicione uma segunda operação compensatória para um serviço fictício, como "Reserva de excursão", e implemente sua lógica de compensação.

### Transação SAGA com múltiplas etapas

1. **Reserve voo**
2. **Reserve hotel**
3. **Reserve carro**
4. **Compensar em caso de falha**

### Exemplo básico

- Ver [código no Gist](https://gist.github.com/zoedsoupe/c79f9a7c71e4a2c3e216dbc54d22851c#file-saga_multi_step-go)
- Exemplo simples de transação com compensações.

#### Desafio 2:
- Insira um novo serviço na cadeia, como a "Reserva de passeio turístico". Implemente a lógica de sucesso e falha para esse serviço.

---

## 3. Gerenciador de SAGA

### Por que usar um Gerenciador?

- **Encapsular operações** e suas compensações.
- Centralizar o **tratamento de erros**.
- Facilitar o **reuso**.

### Estrutura Básica

- **Lista de passos**: ações da transação.
- **Lista de compensações**: desfaz ações em caso de erro.

### Exemplo de Gerenciador

- Ver [código no Gist](https://gist.github.com/zoedsoupe/c79f9a7c71e4a2c3e216dbc54d22851c#file-saga_manager-go)

#### Desafio 3:
- Modifique o gerenciador para incluir um contador de tentativas. Se uma etapa falhar, ela deve ser repetida até 3 vezes antes de disparar a compensação.

---

## 4. Aplicações do Padrão SAGA

### Onde usar?

- **E-commerce**: coordenação de pagamento, estoque, envio.
- **Sistemas financeiros**: transações distribuídas entre contas/bancos.
- **Microserviços**: garantir consistência sem transações ACID.

### Discussão

- Como você aplicaria o padrão SAGA no seu trabalho?
- Qual cenário parece mais relevante?

#### Desafio 4:
- Pense em um exemplo do seu ambiente de trabalho onde o padrão SAGA pode ser aplicado. Qual seria a sequência de ações e compensações?

---

## Links úteis

- [Padrão SAGA - Wikipedia](https://en.wikipedia.org/wiki/Saga_pattern)
- [Microservices.io - SAGA](https://microservices.io/patterns/data/saga.html)
- [Golang Documentation](https://golang.org/doc/)
- [Original SAGA Paper (1987)](https://dl.acm.org/doi/10.1145/38713.38742)

---

## Agradecimentos

Obrigada pela participação!  
[LinkedIn](https://linkedin.com/in/zoedsoupe) | [GitHub](https://github.com/zoedsoupe) | [BlueSky](https://bsky.app/profile/zoedsoupe.zeetech.io)
