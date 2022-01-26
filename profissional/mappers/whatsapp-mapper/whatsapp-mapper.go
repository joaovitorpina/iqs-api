package whatsapp_mapper

import (
	"profissional/dtos"
	"profissional/ent"
)

func ToDomain(model []*ent.WhatsApp) []dtos.BuscarWhatsappsResponse {
	var domain []dtos.BuscarWhatsappsResponse
	for _, whatsApp := range model {
		domain = append(domain, dtos.BuscarWhatsappsResponse{
			Numero:    whatsApp.Numero,
			Principal: whatsApp.Principal,
		})
	}

	return domain
}
