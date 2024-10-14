package main

import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct{}

func (r Relatorio) Imprimir() {
    fmt.Println("Imprimindo Relatório")
}

type Contrato struct{}

func (c Contrato) Imprimir() {
    fmt.Println("Imprimindo Contrato")
}
