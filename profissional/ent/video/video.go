// Code generated by entc, DO NOT EDIT.

package video

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitulo holds the string denoting the titulo field in the database.
	FieldTitulo = "titulo"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldURLThumbnail holds the string denoting the url_thumbnail field in the database.
	FieldURLThumbnail = "url_thumbnail"
	// EdgeProfissional holds the string denoting the profissional edge name in mutations.
	EdgeProfissional = "profissional"
	// Table holds the table name of the video in the database.
	Table = "videos"
	// ProfissionalTable is the table that holds the profissional relation/edge.
	ProfissionalTable = "videos"
	// ProfissionalInverseTable is the table name for the Profissional entity.
	// It exists in this package in order to avoid circular dependency with the "profissional" package.
	ProfissionalInverseTable = "profissionals"
	// ProfissionalColumn is the table column denoting the profissional relation/edge.
	ProfissionalColumn = "profissional_videos"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldTitulo,
	FieldURL,
	FieldURLThumbnail,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "videos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profissional_videos",
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
