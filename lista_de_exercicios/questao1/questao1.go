package main
import "fmt"

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
}

func (c Carro) ExibirDetalhes() string {
    return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d", c.Marca, c.Modelo, c.Ano)
}

func main() {
    carro1 := Carro{"Toyota", "Corolla", 2020}
    carro2 := Carro{"Ford", "Mustang", 2021}
    carro3 := Carro{"Honda", "Civic", 2022}

    fmt.Println(carro1.ExibirDetalhes())
    fmt.Println(carro2.ExibirDetalhes())
    fmt.Println(carro3.ExibirDetalhes())
}
