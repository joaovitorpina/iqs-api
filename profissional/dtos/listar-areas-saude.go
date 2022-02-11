package dtos

type ListarAreasSaudeResponse struct {
	Id              int              `json:"id"`
	Descricao       string           `json:"descricao"`
	Especializacoes []Especializacao `json:"especializacoes"`
}

type Especializacao struct {
	Id        int    `json:"id"`
	Descricao string `json:"descricao"`
}
