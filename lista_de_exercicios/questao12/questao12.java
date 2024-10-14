package lista_de_exercicios.questao12;

class Produto {
    private String nome;
    private double preco;

    Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    Produto somar(Produto outro) {
        return new Produto(this.nome + " + " + outro.nome, this.preco + outro.preco);
    }

    @Override
    public String toString() {
        return nome + ": R$" + String.format("%.2f", preco);
    }

    public static void main(String[] args) {
        Produto produto1 = new Produto("Produto A", 10.50);
        Produto produto2 = new Produto("Produto B", 15.75);
        Produto produto3 = produto1.somar(produto2);
        System.out.println(produto3);
    }
}
