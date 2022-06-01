package contas

import c "Alura/2_Orientacao_a_Objetos/imports/clientes"

type ContaPoupanca struct {
	Titular                              c.Titular
	NumeroConta, NumeroAgencia, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) (string, float64, bool) {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque efetuado com sucesso!", c.saldo, true
	} else if valor <= 0 {
		return "O valor da operação precisa ser superior a R$0,00.", c.saldo, false
	} else {
		return "saldo insuficiente.", c.saldo, false
	}
}
func (c *ContaPoupanca) Depositar(valor float64) (string, float64, bool) {
	if valor <= 0 {
		return "O valor da operação precisa ser superior a R$00,00.", c.saldo, false
	}
	c.saldo += valor
	return "Depósito realizado com sucesso!", c.saldo, true
}
func (c *ContaPoupanca) Transferir(valor float64, destino *ContaCorrente) (string, bool) {
	mensagem, _, status := c.Sacar(valor)
	if status {
		destino.Depositar(valor)
		return "Transação realizada com sucesso!", true
	}
	return mensagem, false
}
func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
