class Animal:
    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        return "Latido"

class Gato(Animal):
    def som(self):
        return "Miau"

animais = [Cachorro(), Gato()]

for animal in animais:
    print(animal.som())
