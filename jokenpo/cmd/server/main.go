package main

import "github.com/Kovokar/go-net-communication/jokenpo/internal/network"

func main() {
	server := network.NewTCPServer("8080")
	server.Start()
}
