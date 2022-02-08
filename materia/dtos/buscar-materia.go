package dtos

type BuscarMateriaResponse struct {
	Titulo        string `json:"titulo"`
	UnidadeId     int    `json:"unidade_id"`
	UrlAmigavel   string `json:"url_amigavel"`
	ImagemUrl     string `json:"imagem_url"`
	Profissionais []int  `json:"profissionais"`
	Texto         string `json:"texto"`
	Fonte         string `json:"fonte,omitempty"`
	LinkFonte     string `json:"link_fonte,omitempty"`
	Patrocinado   bool   `json:"patrocinado"`
}
