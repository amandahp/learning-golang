package main //pacote principal

import (
	"fmt"
)

func main() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando) // primeiro modificador, endereço do comando (&)
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)
}
