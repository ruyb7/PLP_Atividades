package lista_de_exercicios.questao11;

abstract class Funcionario {
    abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    private int horasTrabalhadas;
    private double valorPorHora;

    FuncionarioHorista(int horasTrabalhadas, double valorPorHora) {
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorPorHora = valorPorHora;
    }

    double calcularSalario() {
        return horasTrabalhadas * valorPorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    private double salarioMensal;

    FuncionarioAssalariado(double salarioMensal) {
        this.salarioMensal = salarioMensal;
    }

    double calcularSalario() {
        return salarioMensal;
    }
}
