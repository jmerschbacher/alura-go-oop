package main

import (
	"alura-go-oop/banco/clientes"
	"alura-go-oop/banco/contas"
	"fmt"
)

type contaInterface interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta contaInterface, valor float64) {
	conta.Sacar(valor)
}

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		Cpf:       "463.022.268-19",
		Profissao: "Engenheiro de Software"}

	contaBruno := contas.ContaPoupanca{
		Titular:       clienteBruno,
		NumeroAgencia: 9373,
		NumeroConta:   132623}

	contaBruno.Depositar(100)
	PagarBoleto(&contaBruno, 40)
	fmt.Println(contaBruno.ObterSaldo())

	clienteAna := clientes.Titular{
		Nome:      "Ana",
		Cpf:       "342.313.422-09",
		Profissao: "Administradora"}

	contaAna := contas.ContaCorrente{
		Titular:       clienteAna,
		NumeroAgencia: 9373,
		NumeroConta:   132623}

	contaAna.Depositar(350)
	PagarBoleto(&contaAna, 40)
	fmt.Println(contaAna.ObterSaldo())

}
