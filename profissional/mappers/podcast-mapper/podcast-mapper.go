package podcast_mapper

import (
	"profissional/dtos"
	"profissional/ent"
)

func ToDomain(model []*ent.Podcast) []dtos.BuscarPodcastsResponse {
	var domain []dtos.BuscarPodcastsResponse
	for _, podcast := range model {
		domain = append(domain, dtos.BuscarPodcastsResponse{
			Titulo: podcast.Titulo,
			Codigo: podcast.Codigo,
		})
	}

	return domain
}
