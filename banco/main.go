package main

import (
	"fmt"
	"poo/banco/contas"
)

func main() {
	conta := contas.ContaCorrente{
		Titular: "Lucas",
		Agencia: 11,
		Conta:   1132,
		Saldo:   5368.65,
	}

	conta2 := contas.ContaCorrente{
		Titular: "Iris",
		Agencia: 11,
		Conta:   1562,
		Saldo:   3600,
	}

	sucess, error := conta.Transfere(555.75, &conta2)
	fmt.Println(sucess, error)
	fmt.Println(conta2.Saldo)
	fmt.Println(conta.Saldo)

}
