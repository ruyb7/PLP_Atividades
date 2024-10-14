package main

import (
	"fmt"
)

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) somar(outro Produto) Produto {
	return Produto{nome: p.nome + " + " + outro.nome, preco: p.preco + outro.preco}
}

func (p Produto) String() string {
	return fmt.Sprintf("%s: R$%.2f", p.nome, p.preco)
}

func main() {
	produto1 := Produto{"Produto A", 10.50}
	produto2 := Produto{"Produto B", 15.75}
	produto3 := produto1.somar(produto2)
	fmt.Println(produto3)
}
