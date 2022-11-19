package main

import (
	"fmt"
	"poo/banco/contas"
	"poo/banco/contas/shared"
	"poo/banco/entidades"
)

func main() {
	contaCC := new(contas.ContaCorrente)
	contaPP := new(contas.ContaPoupanca)
	cliente := entidades.Cliente{
		Name:       "Lucas",
		Document:   "1597538565",
		Profession: "Programador",
		Phone:      119955887766,
	}

	cliente2 := entidades.Cliente{
		Name:       "Bruna",
		Document:   "1597538565",
		Profession: "Programador",
		Phone:      119955887766,
	}

	contaCC.SetTitular(cliente)
	contaCC.SetAgencia(11)
	contaCC.SetConta(11369)
	contaCC.Deposita(3500)
	shared.PagarBoleto(contaCC, 50)

	fmt.Println(contaCC.GetSaldo())

	contaPP.SetTitular(cliente2)
	contaPP.SetAgencia(11)
	contaPP.SetConta(11369)
	contaPP.Deposita(3500)
	shared.PagarBoleto(contaPP, 500)

	fmt.Println(contaPP.GetSaldo())
}
