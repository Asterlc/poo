package contas

import (
	titular "poo/banco/entidades"
	handler "poo/banco/handlers"
)

type ContaCorrente struct {
	titular titular.Cliente
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Saque(valor float64) string {
	if valor < c.saldo {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Valor Insuficiente"
	}
}

func (c *ContaCorrente) Deposita(valor float64) (handler.RetornoTransacao, handler.ErrorTransacao) {
	retval := new(handler.RetornoTransacao)
	erroval := new(handler.ErrorTransacao)
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

func (c *ContaCorrente) Transfere(valor float64, destino *ContaCorrente) (handler.RetornoTransacao, handler.ErrorTransacao) {
	retval := new(handler.RetornoTransacao)
	erroval := new(handler.ErrorTransacao)

	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		destino.saldo += valor
		retval.Message = "Sucesso"
		retval.Tipo = "Transferencia"
		retval.SaldoAtual = c.saldo
		retval.Valor = valor
	} else if valor > c.saldo {
		erroval.Reason = "Valor insuficiente"
		erroval.Code = "COD_116_NOFUND"
	} else {
		erroval.Reason = "Valor inválido"
		erroval.Code = "COD_104_INV"
	}

	return *retval, *erroval
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) SetAgencia(numAgencia int) {
	c.agencia = numAgencia
}

func (c *ContaCorrente) GetAgencia() int {
	return c.agencia
}

func (c *ContaCorrente) SetConta(numConta int) {
	c.conta = numConta
}

func (c *ContaCorrente) GetConta() int {
	return c.conta
}

func (c *ContaCorrente) SetTitular(titular titular.Cliente) {
	c.titular = titular
}

func (c *ContaCorrente) GetTitular() titular.Cliente {
	return c.titular
}
