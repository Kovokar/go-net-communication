# net-jokenpo

Um projeto de estudo em **Go** que implementa o jogo **Jokenpô (Pedra, Papel e Tesoura)** utilizando **comunicação bidirecional assíncrona** através do pacote `net`.

O objetivo deste repositório é explorar conceitos fundamentais de **networking em Go**, como conexões TCP, comunicação simultânea entre cliente e servidor e manipulação concorrente de mensagens.

---

## 🎯 Objetivo

O projeto **net-jokenpo** foi criado para estudar:

* Comunicação **bidirecional** entre cliente e servidor
* Troca de mensagens **assíncrona**
* Uso do pacote `net` da linguagem Go
* Concorrência com **goroutines**
* Sincronização de estado entre múltiplos clientes

A ideia é simular partidas de **Jokenpô em rede**, onde dois jogadores conectados enviam suas jogadas e recebem o resultado da rodada.

---

## 🧠 Conceitos estudados

Este projeto explora os seguintes conceitos:

* TCP sockets com `net.Listen` e `net.Conn`
* Comunicação **duplex** (cliente e servidor enviam mensagens simultaneamente)
* Uso de **goroutines** para leitura e escrita concorrente
* Sincronização de partidas
* Estruturação de protocolos simples de aplicação

---

## 🏗️ Estrutura do projeto (exemplo)

```
net-jokenpo/
│
├── cmd/
│   ├── server/
│   │   └── main.go
│   │
│   └── client/
│       └── main.go
│
├── internal/
│   ├── game/
│   │   └── jokenpo.go
│   │
│   └── protocol/
│       └── message.go
│
├── go.mod
└── README.md
```

Descrição:

* **server** → responsável por aceitar conexões e gerenciar partidas
* **client** → interface simples para enviar jogadas
* **game** → lógica do Jokenpô
* **protocol** → definição das mensagens trocadas na rede

---

## 🔄 Como funciona a comunicação

1. O **servidor TCP** inicia e aguarda conexões.
2. Dois clientes se conectam ao servidor.
3. Cada cliente envia sua jogada (`pedra`, `papel` ou `tesoura`).
4. O servidor processa o resultado.
5. O resultado é enviado de volta para ambos os clientes.

A comunicação ocorre de forma **assíncrona**, permitindo:

* leitura contínua da conexão
* envio de mensagens a qualquer momento
* múltiplas goroutines manipulando eventos

Fluxo simplificado:

```
Cliente A  ----\
                ---> Servidor ----> Resultado
Cliente B  ----/
```

---

## ▶️ Executando o projeto

### 1️⃣ Iniciar o servidor

```bash
go run ./cmd/server
```

O servidor iniciará escutando em uma porta TCP (exemplo: `:8080`).

---

### 2️⃣ Iniciar os clientes

Abra dois terminais:

```bash
go run ./cmd/client
```

Cada cliente poderá escolher:

```
1 - Pedra
2 - Papel
3 - Tesoura
```

Após ambos enviarem a jogada, o servidor calcula o resultado e retorna:

```
Você jogou: Pedra
Oponente jogou: Tesoura
Resultado: Você venceu!
```

---

## 📚 Aprendizados esperados

Este projeto ajuda a praticar:

* criação de **servidores TCP**
* manipulação de **conexões persistentes**
* comunicação **full-duplex**
* concorrência com **goroutines**
* construção de **protocolos simples de aplicação**

---

## 🚀 Possíveis melhorias

Algumas ideias para evoluir o projeto:

* suporte a **várias partidas simultâneas**
* sistema de **salas**
* placar entre jogadores
* comunicação usando **JSON**
* timeout de jogadas
* interface web ou CLI mais avançada

---

## 📖 Referências

* Documentação oficial do pacote `net`
* Conceitos de **TCP sockets**
* Concorrência em Go com **goroutines e channels**

---

## 🧑‍💻 Propósito

Este repositório faz parte de um estudo prático sobre **networking e concorrência em Go**, usando um projeto simples para explorar padrões de comunicação utilizados em aplicações distribuídas.
