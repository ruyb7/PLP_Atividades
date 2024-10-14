package lista_de_exercicios.questao9;

interface Imprimivel {
    void imprimir();
}

class Relatorio implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo Relatório");
    }
}

class Contrato implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo Contrato");
    }
}
