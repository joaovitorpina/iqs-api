// Code generated by entc, DO NOT EDIT.

package convenio

const (
	// Label holds the string label denoting the convenio type in the database.
	Label = "convenio"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNome holds the string denoting the nome field in the database.
	FieldNome = "nome"
	// EdgeProfissional holds the string denoting the profissional edge name in mutations.
	EdgeProfissional = "profissional"
	// Table holds the table name of the convenio in the database.
	Table = "convenios"
	// ProfissionalTable is the table that holds the profissional relation/edge.
	ProfissionalTable = "convenios"
	// ProfissionalInverseTable is the table name for the Profissional entity.
	// It exists in this package in order to avoid circular dependency with the "profissional" package.
	ProfissionalInverseTable = "profissionals"
	// ProfissionalColumn is the table column denoting the profissional relation/edge.
	ProfissionalColumn = "profissional_convenios"
)

// Columns holds all SQL columns for convenio fields.
var Columns = []string{
	FieldID,
	FieldNome,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "convenios"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profissional_convenios",
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
