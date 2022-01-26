package videos_mapper

import (
	"profissional/dtos"
	"profissional/ent"
)

func ToDomain(model []*ent.Video) []dtos.BuscarVideosResponse {
	var domain []dtos.BuscarVideosResponse
	for _, video := range model {
		domain = append(domain, dtos.BuscarVideosResponse{
			Titulo:       video.Titulo,
			Url:          video.URL,
			UrlThumbnail: video.URLThumbnail,
		})
	}

	return domain
}
