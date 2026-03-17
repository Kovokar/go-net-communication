# Go `net` Study

Este repositório tem como objetivo estudar a biblioteca **`net`** da linguagem **Go**, explorando conceitos de **comunicação de rede**, **conexões TCP** e **comunicação assíncrona bidirecional**.

O foco principal é entender como aplicações podem trocar mensagens simultaneamente através de conexões persistentes, utilizando os recursos de concorrência da linguagem como **goroutines**.

---

## 🎯 Objetivo do estudo

Este projeto serve como um laboratório para experimentar:

* Uso da biblioteca `net`
* Criação de **servidores TCP**
* Comunicação entre **cliente e servidor**
* Comunicação **assíncrona**
* Comunicação **bidirecional (full-duplex)**
* Uso de **goroutines** para leitura e escrita simultânea
* Estruturação de protocolos simples de rede

---

## 📡 Comunicação assíncrona bidirecional

Na comunicação bidirecional, **cliente e servidor podem enviar e receber mensagens ao mesmo tempo**.

Diferente de modelos síncronos tradicionais (requisição → resposta), neste modelo:

* ambos os lados mantêm a conexão aberta
* mensagens podem ser enviadas a qualquer momento
* múltiplas goroutines lidam com leitura e escrita da conexão

Exemplo de fluxo:

```
Cliente  <------->  Servidor
   ↑                   ↑
 envia                envia
 mensagens            mensagens
 simultaneamente      simultaneamente
```

Esse modelo é muito utilizado em:

* jogos multiplayer
* sistemas de chat
* streaming de dados
* sistemas distribuídos

---

## 🧠 Conceitos estudados

Durante o desenvolvimento deste repositório são explorados:

* `net.Listen`
* `net.Dial`
* `net.Conn`
* leitura e escrita em sockets
* concorrência com goroutines
* controle de conexões
* protocolos simples de aplicação

---

## 🧪 Projeto de exemplo

Para aplicar os conceitos estudados, foi desenvolvido o projeto:

### **net-jokenpo**

Um jogo de **Jokenpô (Pedra, Papel e Tesoura)** jogado através de uma conexão de rede.

Nesse projeto:

* dois clientes conectam ao servidor
* cada jogador envia sua jogada
* o servidor processa o resultado
* o resultado é enviado de volta para ambos os jogadores

Esse exemplo demonstra na prática:

* comunicação **assíncrona**
* troca de mensagens entre múltiplos clientes
* uso de **goroutines para manipular conexões**
* manutenção de estado da aplicação no servidor

---

## 📚 O que este repositório não é

Este repositório **não tem foco em produção**, mas sim em:

* experimentação
* aprendizado
* testes de arquitetura
* entendimento do funcionamento interno da comunicação de rede em Go

---

## 🚀 Possíveis extensões de estudo

Alguns tópicos que podem ser explorados a partir deste estudo:

* protocolos baseados em **JSON**
* multiplexação de conexões
* controle de sessões
* criação de **servidores concorrentes**
* uso de **channels** para coordenação interna
* comparação entre **TCP e UDP**

---

## 📖 Referências

* Documentação oficial do Go
* Pacote `net`
* Conceitos de sockets TCP/IP
* Modelos de comunicação assíncrona

---

## 🧑‍💻 Motivação

A melhor forma de aprender networking é **construindo aplicações reais**.
Este repositório reúne pequenos experimentos e projetos simples para entender, na prática, como funciona a comunicação de rede em Go.
