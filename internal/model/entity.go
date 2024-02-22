package model

type User struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Contato   string `json:"contato"`
	Endereço  string `json:"endereço"`
	DataNasc  string `json:"data_nasc"`
	CPF       string `json:"cpf"`
}
