package lista_de_exercicios.questao5;

abstract class Animal {
    abstract String som();
}

class Cachorro extends Animal {
    @Override
    String som() {
        return "Latido";
    }
}

class Gato extends Animal {
    @Override
    String som() {
        return "Miau";
    }
}

public class Main {
    public static void main(String[] args) {
        Animal[] animais = {new Cachorro(), new Gato()};

        for (Animal animal : animais) {
            System.out.println(animal.som());
        }
    }
}
