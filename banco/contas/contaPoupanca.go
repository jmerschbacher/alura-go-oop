package contas

import "alura-go-oop/banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	isTransacaoValida := valor > 0 && valor <= c.saldo

	if isTransacaoValida {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) string {
	isTransacaoValida := valor > 0

	if isTransacaoValida {
		c.saldo += valor
		return "Depósito realizado com sucesso"
	} else {
		return "Erro ao efetuar depósito: valor deve ser maior que zero"
	}
}

func (c *ContaPoupanca) Transferir(valor float64, contaDestino *ContaCorrente) string {
	isTransacaoValida := valor > 0 && valor <= c.saldo && contaDestino != nil

	if isTransacaoValida {
		c.saldo -= valor
		contaDestino.saldo += valor
		return "Transferência realizada com sucesso"
	} else {
		return "Erro ao realizar transferência"
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
