package main

import (
	"fmt"
)

func main() {
	conta := ContaCorrente{
		titular: "Lucas",
		agencia: 11,
		conta:   1132,
		saldo:   5368.65,
	}

	conta2 := ContaCorrente{
		titular: "Iris",
		agencia: 11,
		conta:   1562,
		saldo:   3600,
	}

	sucess, error := conta.transfere(557.75, &conta2)
	fmt.Println(sucess, error)
	fmt.Println(conta2.saldo)
	fmt.Println(conta.saldo)

}

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

type RetornoTransacao struct {
	message string
	tipo    string
	valor   float64
}

type ErrorTransacao struct {
	reason string
	code   string
}

func (c *ContaCorrente) saque(valor float64) string {
	if valor < c.saldo {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Valor Insuficiente"
	}
}

func (c *ContaCorrente) deposita(valor float64) (RetornoTransacao, ErrorTransacao) {
	retval := new(RetornoTransacao)
	erroval := new(ErrorTransacao)
	if valor > 0 {
		c.saldo += valor
		retval.message = "Sucesso"
		retval.tipo = "Deposito"
		retval.valor = valor
	} else {
		erroval.reason = "Valor inválido"
		erroval.code = "COD_104_INV"
	}

	return *retval, *erroval
}

func (c *ContaCorrente) transfere(valor float64, conta *ContaCorrente) (RetornoTransacao, ErrorTransacao) {
	retval := new(RetornoTransacao)
	erroval := new(ErrorTransacao)

	if valor <= conta.saldo {
		c.saldo -= valor
		conta.saldo += valor
		retval.message = "Sucesso"
		retval.tipo = "Transferencia"
		retval.valor = valor
	} else {
		erroval.reason = "Valor insuficiente"
		erroval.code = "COD_116_NOFUND"
	}

	return *retval, *erroval
}
