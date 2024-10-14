package main

import "errors"

type SaldoInsuficienteException struct {
    message string
}

func (e *SaldoInsuficienteException) Error() string {
    return e.message
}

type ContaBancaria struct {
    saldo float64
}

func NewContaBancaria(saldo float64) *ContaBancaria {
    return &ContaBancaria{saldo: saldo}
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return &SaldoInsuficienteException{message: "Saldo insuficiente para o saque."}
    }
    c.saldo -= valor
    return nil
}
