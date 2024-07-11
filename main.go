package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	var nomePizzaria string = "Pizzaria Go"
	var toscana Pizza
	instagram, telefone := "@pizzaria_go", 11951
	fmt.Println(nomePizzaria, instagram, telefone)
}
