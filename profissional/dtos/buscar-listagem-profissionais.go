package dtos

type BuscarListagemProfissionaisResponse struct {
	Data       []ProfissionalReduzido `json:"data"`
	Pagina     int                    `json:"pagina"`
	Quantidade int                    `json:"quantidade"`
	Total      int                    `json:"total"`
}

type ProfissionalReduzido struct {
	Nome            string   `json:"nome"`
	UrlAmigavel     string   `json:"url_amigavel"`
	Tipo            string   `json:"tipo"`
	Especialidades  []string `json:"especialidades"`
	Recomendado     bool     `json:"recomendado"`
	ImagemPerfilUrl string   `json:"imagem_perfil_url"`
	UnidadeId       int      `json:"unidade_id"`
	WhatsApp        int64    `json:"whatsapp"`
	Email           string   `json:"email,omitempty"`
	Site            string   `json:"site,omitempty"`
	Facebook        string   `json:"facebook,omitempty"`
	Instagram       string   `json:"instagram,omitempty"`
	Youtube         string   `json:"youtube,omitempty"`
	Linkedin        string   `json:"linkedin,omitempty"`
}

type BuscarListagemProfissionaisQuery struct {
	Pagina             int    `form:"pagina"`
	Limite             int    `form:"limite"`
	UnidadeId          int    `form:"unidade_id"`
	Nome               string `form:"nome"`
	TipoProfissionalId int    `form:"tipo_profissional_id"`
	EspecialidadeId    int    `form:"especialidade_id"`
}
