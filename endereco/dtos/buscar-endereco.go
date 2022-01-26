package dtos

type BuscarEnderecoQuery struct {
	Id int `form:"id" binding:"required"`
}

type BuscarEnderecoResponse struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Numero     string `json:"numero"`
	Cep        int32  `json:"cep"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}
