package dtos

type BuscarListagemProfissionaisResponse struct {
	Data       []ProfissionalReduzido `json:"data"`
	Pagina     int                    `json:"pagina"`
	Quantidade int                    `json:"quantidade"`
}

type ProfissionalReduzido struct {
	Id              int      `json:"id"`
	Nome            string   `json:"nome"`
	UrlAmigavel     string   `json:"url_amigavel"`
	Tipo            string   `json:"tipo"`
	Especialidades  []string `json:"especialidades"`
	Recomendado     bool     `json:"recomendado"`
	ImagemPerfilUrl string   `json:"imagem_perfil_url"`
	UnidadeId       int      `json:"unidade_id"`
	Facebook        string   `json:"facebook"`
	WhatsApp        int64    `json:"whatsapp"`
	Instagram       string   `json:"instagram"`
	Email           string   `json:"email"`
	Site            string   `json:"site"`
}

type BuscarListagemProfissionaisQuery struct {
	Pagina             int    `form:"pagina"`
	Limite             int    `form:"limite"`
	UnidadeId          int    `form:"unidade_id"`
	Nome               string `form:"nome"`
	TipoProfissionalId int    `form:"tipo_profissional_id"`
	EspecialidadeId    int    `form:"especialidade_id"`
}
