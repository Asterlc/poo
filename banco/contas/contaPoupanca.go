package contas

import (
	titular "poo/banco/entidades"
	"poo/banco/handlers"
)

type ContaPoupanca struct {
	titular titular.Cliente
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaPoupanca) Saque(valor float64) string {
	if valor < c.saldo {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Valor Insuficiente"
	}
}

func (c *ContaPoupanca) Deposita(valor float64) (handlers.RetornoTransacao, handlers.ErrorTransacao) {
	retval := new(handlers.RetornoTransacao)
	erroval := new(handlers.ErrorTransacao)
	if valor > 0 {
		c.saldo += valor
		retval.Message = "Sucesso"
		retval.Tipo = "Deposito"
		retval.SaldoAtual = c.saldo
		retval.Valor = valor
	} else {
		erroval.Reason = "Valor inválido"
		erroval.Code = "COD_104_INV"
	}

	return *retval, *erroval
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) SetAgencia(numAgencia int) {
	c.agencia = numAgencia
}

func (c *ContaPoupanca) GetAgencia() int {
	return c.agencia
}

func (c *ContaPoupanca) SetConta(numConta int) {
	c.conta = numConta
}

func (c *ContaPoupanca) GetConta() int {
	return c.conta
}

func (c *ContaPoupanca) SetTitular(titular titular.Cliente) {
	c.titular = titular
}

func (c *ContaPoupanca) GetTitular() titular.Cliente {
	return c.titular
}
