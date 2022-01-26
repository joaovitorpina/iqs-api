// Code generated by entc, DO NOT EDIT.

package estado

const (
	// Label holds the string label denoting the estado type in the database.
	Label = "estado"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNome holds the string denoting the nome field in the database.
	FieldNome = "nome"
	// EdgeCidades holds the string denoting the cidades edge name in mutations.
	EdgeCidades = "cidades"
	// Table holds the table name of the estado in the database.
	Table = "estados"
	// CidadesTable is the table that holds the cidades relation/edge.
	CidadesTable = "cidades"
	// CidadesInverseTable is the table name for the Cidade entity.
	// It exists in this package in order to avoid circular dependency with the "cidade" package.
	CidadesInverseTable = "cidades"
	// CidadesColumn is the table column denoting the cidades relation/edge.
	CidadesColumn = "estado_cidades"
)

// Columns holds all SQL columns for estado fields.
var Columns = []string{
	FieldID,
	FieldNome,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NomeValidator is a validator for the "nome" field. It is called by the builders before save.
	NomeValidator func(string) error
)
