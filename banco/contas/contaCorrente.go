package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

type RetornoTransacao struct {
	message    string
	tipo       string
	saldoAtual float64
	valor      float64
}

type ErrorTransacao struct {
	reason string
	code   string
}

func (c *ContaCorrente) Saque(valor float64) string {
	if valor < c.Saldo {
		c.Saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Valor Insuficiente"
	}
}

func (c *ContaCorrente) Deposita(valor float64) (RetornoTransacao, ErrorTransacao) {
	retval := new(RetornoTransacao)
	erroval := new(ErrorTransacao)
	if valor > 0 && valor <= c.Saldo {
		c.Saldo += valor
		retval.message = "Sucesso"
		retval.tipo = "Deposito"
		retval.saldoAtual = c.Saldo
		retval.valor = valor
	} else {
		erroval.reason = "Valor inválido"
		erroval.code = "COD_104_INV"
	}

	return *retval, *erroval
}

func (c *ContaCorrente) Transfere(valor float64, destino *ContaCorrente) (RetornoTransacao, ErrorTransacao) {
	retval := new(RetornoTransacao)
	erroval := new(ErrorTransacao)

	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		destino.Saldo += valor
		retval.message = "Sucesso"
		retval.tipo = "Transferencia"
		retval.saldoAtual = c.Saldo
		retval.valor = valor
	} else if valor > c.Saldo {
		erroval.reason = "Valor insuficiente"
		erroval.code = "COD_116_NOFUND"
	} else {
		erroval.reason = "Valor inválido"
		erroval.code = "COD_104_INV"
	}

	return *retval, *erroval
}
