package main

type ContaBancaria struct {
    titular string
    saldo   float64
}

func NovaConta(titular string, saldo float64) *ContaBancaria {
    return &ContaBancaria{titular: titular, saldo: saldo}
}

func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.saldo += valor
    }
}

func (c *ContaBancaria) Sacar(valor float64) {
    if valor > 0 && valor <= c.saldo {
        c.saldo -= valor
    }
}

func (c *ContaBancaria) GetSaldo() float64 {
    return c.saldo
}
