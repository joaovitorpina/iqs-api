// Code generated by entc, DO NOT EDIT.

package materia

const (
	// Label holds the string label denoting the materia type in the database.
	Label = "materia"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUnidadeID holds the string denoting the unidade_id field in the database.
	FieldUnidadeID = "unidade_id"
	// FieldTitulo holds the string denoting the titulo field in the database.
	FieldTitulo = "titulo"
	// FieldURLAmigavel holds the string denoting the url_amigavel field in the database.
	FieldURLAmigavel = "url_amigavel"
	// FieldDataAgendamento holds the string denoting the data_agendamento field in the database.
	FieldDataAgendamento = "data_agendamento"
	// FieldFonte holds the string denoting the fonte field in the database.
	FieldFonte = "fonte"
	// FieldLinkFonte holds the string denoting the link_fonte field in the database.
	FieldLinkFonte = "link_fonte"
	// FieldTexto holds the string denoting the texto field in the database.
	FieldTexto = "texto"
	// FieldImagemURL holds the string denoting the imagem_url field in the database.
	FieldImagemURL = "imagem_url"
	// FieldPatrocinado holds the string denoting the patrocinado field in the database.
	FieldPatrocinado = "patrocinado"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeCategoria holds the string denoting the categoria edge name in mutations.
	EdgeCategoria = "categoria"
	// EdgeProfissionalMaterias holds the string denoting the profissional_materias edge name in mutations.
	EdgeProfissionalMaterias = "profissional_materias"
	// Table holds the table name of the materia in the database.
	Table = "materia"
	// CategoriaTable is the table that holds the categoria relation/edge.
	CategoriaTable = "materia"
	// CategoriaInverseTable is the table name for the Categoria entity.
	// It exists in this package in order to avoid circular dependency with the "categoria" package.
	CategoriaInverseTable = "categoria"
	// CategoriaColumn is the table column denoting the categoria relation/edge.
	CategoriaColumn = "categoria_materias"
	// ProfissionalMateriasTable is the table that holds the profissional_materias relation/edge.
	ProfissionalMateriasTable = "profissional_materias"
	// ProfissionalMateriasInverseTable is the table name for the ProfissionalMaterias entity.
	// It exists in this package in order to avoid circular dependency with the "profissionalmaterias" package.
	ProfissionalMateriasInverseTable = "profissional_materias"
	// ProfissionalMateriasColumn is the table column denoting the profissional_materias relation/edge.
	ProfissionalMateriasColumn = "materia_profissional_materias"
)

// Columns holds all SQL columns for materia fields.
var Columns = []string{
	FieldID,
	FieldUnidadeID,
	FieldTitulo,
	FieldURLAmigavel,
	FieldDataAgendamento,
	FieldFonte,
	FieldLinkFonte,
	FieldTexto,
	FieldImagemURL,
	FieldPatrocinado,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "materia"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"categoria_materias",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPatrocinado holds the default value on creation for the "patrocinado" field.
	DefaultPatrocinado bool
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
)
