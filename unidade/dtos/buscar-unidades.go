package dtos

type BuscarUnidadesResponse struct {
	Id          int     `json:"id"`
	Descricao   string  `json:"descricao"`
	UrlAmigavel string  `json:"url_amigavel"`
	EnderecoId  int     `json:"endereco_id"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Telefone    int64   `json:"telefone,omitempty"`
	Celular     int64   `json:"celular,omitempty"`
	Email       string  `json:"email"`
	Facebook    string  `json:"facebook,omitempty"`
	Instagram   string  `json:"instagram,omitempty"`
	Youtube     string  `json:"youtube,omitempty"`
	Ativo       bool    `json:"ativo"`
}
