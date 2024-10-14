package main

import "sync"

type Configuracao struct{}

var instancia *Configuracao
var once sync.Once

func GetInstancia() *Configuracao {
    once.Do(func() {
        instancia = &Configuracao{}
    })
    return instancia
}
