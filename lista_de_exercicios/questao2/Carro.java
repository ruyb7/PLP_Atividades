package lista_de_exercicios.questao2;

public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    public void acelerar(int incremento) {
        this.velocidade += incremento;
    }

    public void frear(int decremento) {
        this.velocidade = Math.max(0, this.velocidade - decremento);
    }

    public String exibirDetalhes() {
        return "Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano;
    }

    public String exibirVelocidade() {
        return "Velocidade atual: " + velocidade + " km/h";
    }

    public static void main(String[] args) {
        Carro carro1 = new Carro("Toyota", "Corolla", 2020);
        carro1.acelerar(50);
        System.out.println(carro1.exibirVelocidade());
        carro1.frear(20);
        System.out.println(carro1.exibirVelocidade());
    }
}
