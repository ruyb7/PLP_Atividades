package lista_de_exercicios.questao4;

abstract class Animal {
    abstract String som();
}

class Cachorro extends Animal {
    String som() {
        return "Latido";
    }
}

class Gato extends Animal {
    String som() {
        return "Miau";
    }
}

public class Main {
    public static void main(String[] args) {
        Animal cachorro = new Cachorro();
        Animal gato = new Gato();
        System.out.println(cachorro.som());  // Latido
        System.out.println(gato.som());      // Miau
    }
}
