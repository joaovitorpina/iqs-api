package mappers

import "materia/ent"

func ToDomain(model []*ent.ProfissionalMaterias) []int {
	var domain []int
	for _, profissionalMaterias := range model {
		domain = append(domain, profissionalMaterias.ProfissionalID)
	}

	return domain
}
