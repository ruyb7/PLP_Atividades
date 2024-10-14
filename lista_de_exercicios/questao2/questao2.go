package main
import "fmt"

type Carro struct {
    Marca     string
    Modelo    string
    Ano       int
    Velocidade int
}

func (c *Carro) Acelerar(incremento int) {
    c.Velocidade += incremento
}

func (c *Carro) Frear(decremento int) {
    if c.Velocidade >= decremento {
        c.Velocidade -= decremento
    } else {
        c.Velocidade = 0
    }
}

func (c Carro) ExibirDetalhes() string {
    return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d", c.Marca, c.Modelo, c.Ano)
}

func (c Carro) ExibirVelocidade() string {
    return fmt.Sprintf("Velocidade atual: %d km/h", c.Velocidade)
}

func main() {
    carro1 := Carro{"Toyota", "Corolla", 2020, 0}
    carro1.Acelerar(50)
    fmt.Println(carro1.ExibirVelocidade())
    carro1.Frear(20)
    fmt.Println(carro1.ExibirVelocidade())
}
