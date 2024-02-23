package model

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Contato   string `json:"contato"`
	Endereço  string `json:"endereço"`
	DataNasc  string `json:"data_nasc"`
	CPF       string `json:"cpf"`
}

type QueueBody struct {
	MessageId uuid.UUID
	ID        int64       `json:"id,omitempty"`
	Method    string      `json:"method"`
	Content   interface{} `json:"content,omitempty"`
	User      User        `json:"user,omitempty"`
}
