package main //pacote principal

import (
	"fmt"
	"reflect"
)

//tipo

func main() {
	cidade := "Cascavel" //inferência de tipo outra maneira de declarar var
	var nome string = "Douglas"
	var idade int
	var versao float32 = 1.1
	fmt.Println("Olá, sr.", nome, idade, ". Mora na cidade", cidade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel é", reflect.TypeOf(versao))
}
