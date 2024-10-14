package main

import "fmt"

type Funcionario interface {
	calcularSalario() float64
}

type FuncionarioHorista struct {
	horasTrabalhadas int
	valorPorHora     float64
}

func (f FuncionarioHorista) calcularSalario() float64 {
	return float64(f.horasTrabalhadas) * f.valorPorHora
}

type FuncionarioAssalariado struct {
	salarioMensal float64
}

func (f FuncionarioAssalariado) calcularSalario() float64 {
	return f.salarioMensal
}

func main() {
	fh := FuncionarioHorista{horasTrabalhadas: 40, valorPorHora: 15.0}
	fa := FuncionarioAssalariado{salarioMensal: 3000.0}
	fmt.Println(fh.calcularSalario())
	fmt.Println(fa.calcularSalario())
}
