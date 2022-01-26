package dtos

type BuscarVideosResponse struct {
	Titulo       string `json:"titulo"`
	Url          string `json:"url"`
	UrlThumbnail string `json:"url_thumbnail"`
}
