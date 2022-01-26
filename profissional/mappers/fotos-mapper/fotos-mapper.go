package fotos_mapper

import (
	"profissional/dtos"
	"profissional/ent"
)

func ToDomain(model []*ent.Foto) []dtos.BuscarFotosResponse {
	var domain []dtos.BuscarFotosResponse
	for _, foto := range model {
		domain = append(domain, dtos.BuscarFotosResponse{
			Titulo: foto.Titulo,
			Url:    foto.URL,
		})
	}

	return domain
}
