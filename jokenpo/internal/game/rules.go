package game

import "fmt"

type Option int

const (
	Pedra Option = iota
	Papel
	Tesoura
)

func showOptions() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Pedra")
	fmt.Println("2 - Papel")
	fmt.Println("3 - Tesoura")
}
