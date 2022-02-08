package dtos

type ListarMateriasQuery struct {
	Titulo     string `query:"titulo"`
	Categorias []int  `query:"categorias"`
	Pagina     int    `form:"pagina"`
	Limite     int    `form:"limite"`
}

type ListarMateriasResponse struct {
	Data       []MateriaReduzida `json:"data"`
	Pagina     int               `json:"pagina"`
	Quantidade int               `json:"quantidade"`
	Total      int               `json:"total"`
}

type MateriaReduzida struct {
	Titulo        string `json:"titulo"`
	UnidadeId     int    `json:"unidade_id"`
	UrlAmigavel   string `json:"url_amigavel"`
	ImagemUrl     string `json:"imagem_url"`
	Profissionais []int  `json:"profissionais"`
	Texto         string `json:"texto"`
}
