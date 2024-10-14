package main

import "fmt"

type Motor struct {
    Potencia int
}

type Carro struct {
    Modelo string
    Motor  Motor
}

func main() {
    motor := Motor{Potencia: 200}
    carro := Carro{Modelo: "Sedan", Motor: motor}
    fmt.Println(carro.Modelo, carro.Motor.Potencia)
}
