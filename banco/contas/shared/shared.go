package shared

func PagarBoleto(conta VerificaConta, valor float64) {
	conta.Saque(valor)
}

type VerificaConta interface {
	Saque(valor float64) string
}
