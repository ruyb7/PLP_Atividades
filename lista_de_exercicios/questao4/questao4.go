package main

import "fmt"

type Animal interface {
    Som() string
}

type Cachorro struct{}

func (c Cachorro) Som() string {
    return "Latido"
}

type Gato struct{}

func (g Gato) Som() string {
    return "Miau"
}

func main() {
    var cachorro Animal = Cachorro{}
    var gato Animal = Gato{}
    fmt.Println(cachorro.Som())  // Latido
    fmt.Println(gato.Som())      // Miau
}
