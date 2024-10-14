package main

import "fmt"

type Animal interface {
    Som() string
}

type Cachorro struct{}

func (C Cachorro) Som() string {
    return "Latido"
}

type Gato struct{}

func (G Gato) Som() string {
    return "Miau"
}

func main() {
    animais := []Animal{Cachorro{}, Gato{}}

    for _, animal := range animais {
        fmt.Println(animal.Som())
    }
}
