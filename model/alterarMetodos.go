package model

import (
	"time"
)

type PessoaAlterar struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

// Metodo
func (p *PessoaAlterar) CalculaIdade2() {

	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade += anoAtual - anoDeNascimento
}
