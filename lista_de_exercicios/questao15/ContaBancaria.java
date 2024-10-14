package lista_de_exercicios.questao15;

class questao15{
    public class SaldoInsuficienteException extends Exception {
        public SaldoInsuficienteException(String message) {
            super(message);
        }
    }
    
    public class ContaBancaria {
        private double saldo;
    
        public ContaBancaria(double saldo) {
            this.saldo = saldo;
        }
    
        public void sacar(double valor) throws SaldoInsuficienteException {
            if (valor > saldo) {
                throw new SaldoInsuficienteException("Saldo insuficiente para o saque.");
            }
            saldo -= valor;
        }
    }
}

