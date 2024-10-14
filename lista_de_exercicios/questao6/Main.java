package lista_de_exercicios.questao6;

class Motor {
    private int potencia;

    public Motor(int potencia) {
        this.potencia = potencia;
    }

    public int getPotencia() {
        return potencia;
    }
}

class Carro {
    private String modelo;
    private Motor motor;

    public Carro(String modelo, Motor motor) {
        this.modelo = modelo;
        this.motor = motor;
    }

    public String getModelo() {
        return modelo;
    }

    public Motor getMotor() {
        return motor;
    }
}

public class Main {
    public static void main(String[] args) {
        Motor motor = new Motor(200);
        Carro carro = new Carro("Sedan", motor);
        System.out.println(carro.getModelo() + " " + carro.getMotor().getPotencia());
    }
}
