package main

import "fmt"

func main() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var command int
	fmt.Scan(&command)

	switch command {
	case 0:
		fmt.Println("Saindo...")
		break
	case 1:
		fmt.Println("Monitorando...")
		break
	case 2:
		fmt.Println("Exibindo logs...")
		break
	default:
		fmt.Println("Nāo conheço esse comando")
	}

}
