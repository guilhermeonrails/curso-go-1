package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	// toscana, Marguerita, atum com queijo
	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 49.5},
		{ID: 2, nome: "Marguerita", preco: 79.5},
		{ID: 3, nome: "Atum com queijo", preco: 69.5},
	}
	fmt.Println(pizzas)
}
