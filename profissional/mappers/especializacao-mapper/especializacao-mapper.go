package especializacao_mapper

import "profissional/ent"

func ToDomain(model []*ent.Especializacao) []string {
	var domain []string
	for _, especializacao := range model {
		domain = append(domain, especializacao.Descricao)
	}

	return domain
}
