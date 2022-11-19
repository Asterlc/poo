package handlers

type RetornoTransacao struct {
	Message    string
	Tipo       string
	SaldoAtual float64
	Valor      float64
}

type ErrorTransacao struct {
	Reason string
	Code   string
}
