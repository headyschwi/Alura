package main

import (
	c "Alura/2_Orientacao_a_Objetos/imports/contas"
)

func main() {
	conta := c.ContaCorrente{}
	conta.Depositar(1000)

	PagarBoleto(&conta, 100)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64, bool)
}

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}
