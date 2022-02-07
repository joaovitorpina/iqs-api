package dtos

type BuscarEnderecoResponse struct {
	Id         int    `json:"id"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Numero     string `json:"numero"`
	Cep        int32  `json:"cep"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}
