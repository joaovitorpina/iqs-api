// Code generated by entc, DO NOT EDIT.

package profissional

import (
	"time"
)

const (
	// Label holds the string label denoting the profissional type in the database.
	Label = "profissional"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldNome holds the string denoting the nome field in the database.
	FieldNome = "nome"
	// FieldURLAmigavel holds the string denoting the url_amigavel field in the database.
	FieldURLAmigavel = "url_amigavel"
	// FieldRecomendado holds the string denoting the recomendado field in the database.
	FieldRecomendado = "recomendado"
	// FieldAtivo holds the string denoting the ativo field in the database.
	FieldAtivo = "ativo"
	// FieldSobre holds the string denoting the sobre field in the database.
	FieldSobre = "sobre"
	// FieldConselho holds the string denoting the conselho field in the database.
	FieldConselho = "conselho"
	// FieldNumeroIdentificacao holds the string denoting the numero_identificacao field in the database.
	FieldNumeroIdentificacao = "numero_identificacao"
	// FieldTelefone holds the string denoting the telefone field in the database.
	FieldTelefone = "telefone"
	// FieldCelular holds the string denoting the celular field in the database.
	FieldCelular = "celular"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldSite holds the string denoting the site field in the database.
	FieldSite = "site"
	// FieldFacebook holds the string denoting the facebook field in the database.
	FieldFacebook = "facebook"
	// FieldInstagram holds the string denoting the instagram field in the database.
	FieldInstagram = "instagram"
	// FieldYoutube holds the string denoting the youtube field in the database.
	FieldYoutube = "youtube"
	// FieldLinkedin holds the string denoting the linkedin field in the database.
	FieldLinkedin = "linkedin"
	// FieldUnidadeID holds the string denoting the unidade_id field in the database.
	FieldUnidadeID = "unidade_id"
	// FieldEnderecoID holds the string denoting the endereco_id field in the database.
	FieldEnderecoID = "endereco_id"
	// FieldImagemPerfilURL holds the string denoting the imagem_perfil_url field in the database.
	FieldImagemPerfilURL = "imagem_perfil_url"
	// EdgeWhatsapps holds the string denoting the whatsapps edge name in mutations.
	EdgeWhatsapps = "whatsapps"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// EdgeTratamentos holds the string denoting the tratamentos edge name in mutations.
	EdgeTratamentos = "tratamentos"
	// EdgePodcasts holds the string denoting the podcasts edge name in mutations.
	EdgePodcasts = "podcasts"
	// EdgeFotos holds the string denoting the fotos edge name in mutations.
	EdgeFotos = "fotos"
	// EdgeConvenios holds the string denoting the convenios edge name in mutations.
	EdgeConvenios = "convenios"
	// EdgeEspecializacoes holds the string denoting the especializacoes edge name in mutations.
	EdgeEspecializacoes = "especializacoes"
	// Table holds the table name of the profissional in the database.
	Table = "profissionals"
	// WhatsappsTable is the table that holds the whatsapps relation/edge.
	WhatsappsTable = "whats_apps"
	// WhatsappsInverseTable is the table name for the WhatsApp entity.
	// It exists in this package in order to avoid circular dependency with the "whatsapp" package.
	WhatsappsInverseTable = "whats_apps"
	// WhatsappsColumn is the table column denoting the whatsapps relation/edge.
	WhatsappsColumn = "profissional_whatsapps"
	// VideosTable is the table that holds the videos relation/edge.
	VideosTable = "videos"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "profissional_videos"
	// TratamentosTable is the table that holds the tratamentos relation/edge.
	TratamentosTable = "tratamentos"
	// TratamentosInverseTable is the table name for the Tratamento entity.
	// It exists in this package in order to avoid circular dependency with the "tratamento" package.
	TratamentosInverseTable = "tratamentos"
	// TratamentosColumn is the table column denoting the tratamentos relation/edge.
	TratamentosColumn = "profissional_tratamentos"
	// PodcastsTable is the table that holds the podcasts relation/edge.
	PodcastsTable = "podcasts"
	// PodcastsInverseTable is the table name for the Podcast entity.
	// It exists in this package in order to avoid circular dependency with the "podcast" package.
	PodcastsInverseTable = "podcasts"
	// PodcastsColumn is the table column denoting the podcasts relation/edge.
	PodcastsColumn = "profissional_podcasts"
	// FotosTable is the table that holds the fotos relation/edge.
	FotosTable = "fotos"
	// FotosInverseTable is the table name for the Foto entity.
	// It exists in this package in order to avoid circular dependency with the "foto" package.
	FotosInverseTable = "fotos"
	// FotosColumn is the table column denoting the fotos relation/edge.
	FotosColumn = "profissional_fotos"
	// ConveniosTable is the table that holds the convenios relation/edge.
	ConveniosTable = "convenios"
	// ConveniosInverseTable is the table name for the Convenio entity.
	// It exists in this package in order to avoid circular dependency with the "convenio" package.
	ConveniosInverseTable = "convenios"
	// ConveniosColumn is the table column denoting the convenios relation/edge.
	ConveniosColumn = "profissional_convenios"
	// EspecializacoesTable is the table that holds the especializacoes relation/edge. The primary key declared below.
	EspecializacoesTable = "especializacao_profissionais"
	// EspecializacoesInverseTable is the table name for the Especializacao entity.
	// It exists in this package in order to avoid circular dependency with the "especializacao" package.
	EspecializacoesInverseTable = "especializacaos"
)

// Columns holds all SQL columns for profissional fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldNome,
	FieldURLAmigavel,
	FieldRecomendado,
	FieldAtivo,
	FieldSobre,
	FieldConselho,
	FieldNumeroIdentificacao,
	FieldTelefone,
	FieldCelular,
	FieldEmail,
	FieldSite,
	FieldFacebook,
	FieldInstagram,
	FieldYoutube,
	FieldLinkedin,
	FieldUnidadeID,
	FieldEnderecoID,
	FieldImagemPerfilURL,
}

var (
	// EspecializacoesPrimaryKey and EspecializacoesColumn2 are the table columns denoting the
	// primary key for the especializacoes relation (M2M).
	EspecializacoesPrimaryKey = []string{"especializacao_id", "profissional_id"}
)

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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NomeValidator is a validator for the "nome" field. It is called by the builders before save.
	NomeValidator func(string) error
	// URLAmigavelValidator is a validator for the "url_amigavel" field. It is called by the builders before save.
	URLAmigavelValidator func(string) error
	// DefaultRecomendado holds the default value on creation for the "recomendado" field.
	DefaultRecomendado bool
	// DefaultAtivo holds the default value on creation for the "ativo" field.
	DefaultAtivo bool
	// UnidadeIDValidator is a validator for the "unidade_id" field. It is called by the builders before save.
	UnidadeIDValidator func(int) error
	// EnderecoIDValidator is a validator for the "endereco_id" field. It is called by the builders before save.
	EnderecoIDValidator func(int) error
)
