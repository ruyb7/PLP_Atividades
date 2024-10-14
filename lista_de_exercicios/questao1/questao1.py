class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def exibir_detalhes(self):
        return f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}"

carro1 = Carro("Toyota", "Corolla", 2020)
carro2 = Carro("Ford", "Mustang", 2021)
carro3 = Carro("Honda", "Civic", 2022)

print(carro1.exibir_detalhes())
print(carro2.exibir_detalhes())
print(carro3.exibir_detalhes())
