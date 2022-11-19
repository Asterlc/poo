package main

import (
	"fmt"
	"strconv"
)

func main() {
	conta := ContaCorrente{
		titular: "Lucas",
		agencia: 11,
		conta:   1132,
		saldo:   5368.65,
	}
	conta.saque(2500)
	fmt.Println(strconv.FormatFloat(conta.saldo, 'f', 4, 32))

	conta.deposita(500)
	fmt.Println(float32(conta.saldo))

}

type ContaPoupanca struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) saque(valor float64) string {
	if valor < c.saldo {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Valor Insuficiente"
	}
}

func (c *ContaCorrente) deposita(valor float64) {
	c.saldo += valor
}
