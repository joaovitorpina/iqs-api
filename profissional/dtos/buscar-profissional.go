package dtos

type BuscarProfissionalPorIdResponse struct {
	Id                  int      `json:"id"`
	Nome                string   `json:"nome"`
	ImagemPerfilUrl     string   `json:"imagem_perfil_url"`
	EnderecoId          int      `json:"endereco_id"`
	Tipo                string   `json:"tipo"`
	Especialidades      []string `json:"especialidades"`
	Conselho            string   `json:"conselho"`
	NumeroIdentificacao string   `json:"numero_identificacao"`
	Telefone            int64    `json:"telefone"`
	Celular             int64    `json:"celular"`
	Facebook            string   `json:"facebook"`
	Instagram           string   `json:"instagram"`
	Email               string   `json:"email"`
	Site                string   `json:"site"`
	Sobre               string   `json:"sobre"`
}

type BuscarProfissionalPorUrlAmigavel struct {
	Nome                string   `json:"nome"`
	UrlAmigavel         string   `json:"url_amigavel"`
	Recomendado         bool     `json:"recomendado"`
	Sobre               string   `json:"sobre,omitempty"`
	Conselho            string   `json:"conselho,omitempty"`
	NumeroIdentificacao string   `json:"numero_identificacao,omitempty"`
	Telefone            int64    `json:"telefone"`
	Celular             int64    `json:"celular"`
	Email               string   `json:"email,omitempty"`
	Site                string   `json:"site,omitempty"`
	Facebook            string   `json:"facebook,omitempty"`
	Instagram           string   `json:"instagram,omitempty"`
	Youtube             string   `json:"youtube,omitempty"`
	Linkedin            string   `json:"linkedin,omitempty"`
	UnidadeId           int      `json:"unidade_id"`
	EnderecoId          int      `json:"endereco_id"`
	ImagemPerfilUrl     string   `json:"imagem_perfil_url"`
	Tipo                string   `json:"tipo"`
	Especialidades      []string `json:"especialidades"`
}
