package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kovokar/go-net-communication/jokenpo/internal/network"
)

func main() {
	client, err := network.NewTCPClient("localhost:8080")
	if err != nil {
		panic(err)
	}

	name := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite seu nome: ")
	name.Scan()
	client.Send(name.Text())

	defer client.Close()

	fmt.Println("Conectado ao servidor")

	// Responsabilidade de ouvir o servidor (Background)
	go client.Listen(func(res string) {
		fmt.Printf("\nResposta do servidor: %s\nDigite uma mensagem: ", res)
	})

	// Responsabilidade de ler o input do usuário
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Digite uma mensagem: ")
		if !scanner.Scan() {
			break
		}
		client.Send(scanner.Text())
	}
}
