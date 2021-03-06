// Code generated by entc, DO NOT EDIT.

package cidade

const (
	// Label holds the string label denoting the cidade type in the database.
	Label = "cidade"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNome holds the string denoting the nome field in the database.
	FieldNome = "nome"
	// EdgeEstado holds the string denoting the estado edge name in mutations.
	EdgeEstado = "estado"
	// EdgeCeps holds the string denoting the ceps edge name in mutations.
	EdgeCeps = "ceps"
	// Table holds the table name of the cidade in the database.
	Table = "cidades"
	// EstadoTable is the table that holds the estado relation/edge.
	EstadoTable = "cidades"
	// EstadoInverseTable is the table name for the Estado entity.
	// It exists in this package in order to avoid circular dependency with the "estado" package.
	EstadoInverseTable = "estados"
	// EstadoColumn is the table column denoting the estado relation/edge.
	EstadoColumn = "estado_cidades"
	// CepsTable is the table that holds the ceps relation/edge.
	CepsTable = "ceps"
	// CepsInverseTable is the table name for the Cep entity.
	// It exists in this package in order to avoid circular dependency with the "cep" package.
	CepsInverseTable = "ceps"
	// CepsColumn is the table column denoting the ceps relation/edge.
	CepsColumn = "cidade_ceps"
)

// Columns holds all SQL columns for cidade fields.
var Columns = []string{
	FieldID,
	FieldNome,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cidades"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"estado_cidades",
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
	// NomeValidator is a validator for the "nome" field. It is called by the builders before save.
	NomeValidator func(string) error
)
