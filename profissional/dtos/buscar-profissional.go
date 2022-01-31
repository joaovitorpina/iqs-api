package dtos

type BuscarProfissionalResponse struct {
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
