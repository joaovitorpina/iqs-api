// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AreaSaudesColumns holds the columns for the "area_saudes" table.
	AreaSaudesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "descricao", Type: field.TypeString},
	}
	// AreaSaudesTable holds the schema information for the "area_saudes" table.
	AreaSaudesTable = &schema.Table{
		Name:       "area_saudes",
		Columns:    AreaSaudesColumns,
		PrimaryKey: []*schema.Column{AreaSaudesColumns[0]},
	}
	// ConveniosColumns holds the columns for the "convenios" table.
	ConveniosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nome", Type: field.TypeString, Unique: true},
		{Name: "profissional_convenios", Type: field.TypeInt, Nullable: true},
	}
	// ConveniosTable holds the schema information for the "convenios" table.
	ConveniosTable = &schema.Table{
		Name:       "convenios",
		Columns:    ConveniosColumns,
		PrimaryKey: []*schema.Column{ConveniosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "convenios_profissionals_convenios",
				Columns:    []*schema.Column{ConveniosColumns[2]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EspecializacaosColumns holds the columns for the "especializacaos" table.
	EspecializacaosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "descricao", Type: field.TypeString},
		{Name: "area_saude_especializacoes", Type: field.TypeInt, Nullable: true},
	}
	// EspecializacaosTable holds the schema information for the "especializacaos" table.
	EspecializacaosTable = &schema.Table{
		Name:       "especializacaos",
		Columns:    EspecializacaosColumns,
		PrimaryKey: []*schema.Column{EspecializacaosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "especializacaos_area_saudes_especializacoes",
				Columns:    []*schema.Column{EspecializacaosColumns[2]},
				RefColumns: []*schema.Column{AreaSaudesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FotosColumns holds the columns for the "fotos" table.
	FotosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "titulo", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "profissional_fotos", Type: field.TypeInt, Nullable: true},
	}
	// FotosTable holds the schema information for the "fotos" table.
	FotosTable = &schema.Table{
		Name:       "fotos",
		Columns:    FotosColumns,
		PrimaryKey: []*schema.Column{FotosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fotos_profissionals_fotos",
				Columns:    []*schema.Column{FotosColumns[3]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PodcastsColumns holds the columns for the "podcasts" table.
	PodcastsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "titulo", Type: field.TypeString},
		{Name: "codigo", Type: field.TypeString},
		{Name: "profissional_podcasts", Type: field.TypeInt, Nullable: true},
	}
	// PodcastsTable holds the schema information for the "podcasts" table.
	PodcastsTable = &schema.Table{
		Name:       "podcasts",
		Columns:    PodcastsColumns,
		PrimaryKey: []*schema.Column{PodcastsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "podcasts_profissionals_podcasts",
				Columns:    []*schema.Column{PodcastsColumns[3]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProfissionalsColumns holds the columns for the "profissionals" table.
	ProfissionalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "nome", Type: field.TypeString},
		{Name: "url_amigavel", Type: field.TypeString},
		{Name: "recomendado", Type: field.TypeBool, Default: false},
		{Name: "ativo", Type: field.TypeBool, Default: true},
		{Name: "sobre", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "conselho", Type: field.TypeString, Nullable: true},
		{Name: "numero_identificacao", Type: field.TypeString, Nullable: true},
		{Name: "telefone", Type: field.TypeInt64},
		{Name: "celular", Type: field.TypeInt64},
		{Name: "email", Type: field.TypeString},
		{Name: "site", Type: field.TypeString},
		{Name: "facebook", Type: field.TypeString, Nullable: true},
		{Name: "instagram", Type: field.TypeString, Nullable: true},
		{Name: "youtube", Type: field.TypeString, Nullable: true},
		{Name: "linkedin", Type: field.TypeString, Nullable: true},
		{Name: "unidade_id", Type: field.TypeInt},
		{Name: "endereco_id", Type: field.TypeInt},
		{Name: "imagem_perfil_url", Type: field.TypeString},
	}
	// ProfissionalsTable holds the schema information for the "profissionals" table.
	ProfissionalsTable = &schema.Table{
		Name:       "profissionals",
		Columns:    ProfissionalsColumns,
		PrimaryKey: []*schema.Column{ProfissionalsColumns[0]},
	}
	// TratamentosColumns holds the columns for the "tratamentos" table.
	TratamentosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "descricao", Type: field.TypeString},
		{Name: "profissional_tratamentos", Type: field.TypeInt, Nullable: true},
	}
	// TratamentosTable holds the schema information for the "tratamentos" table.
	TratamentosTable = &schema.Table{
		Name:       "tratamentos",
		Columns:    TratamentosColumns,
		PrimaryKey: []*schema.Column{TratamentosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tratamentos_profissionals_tratamentos",
				Columns:    []*schema.Column{TratamentosColumns[2]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "titulo", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "url_thumbnail", Type: field.TypeString},
		{Name: "profissional_videos", Type: field.TypeInt, Nullable: true},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "videos_profissionals_videos",
				Columns:    []*schema.Column{VideosColumns[4]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WhatsAppsColumns holds the columns for the "whats_apps" table.
	WhatsAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "numero", Type: field.TypeInt64},
		{Name: "principal", Type: field.TypeBool, Default: false},
		{Name: "profissional_whatsapps", Type: field.TypeInt, Nullable: true},
	}
	// WhatsAppsTable holds the schema information for the "whats_apps" table.
	WhatsAppsTable = &schema.Table{
		Name:       "whats_apps",
		Columns:    WhatsAppsColumns,
		PrimaryKey: []*schema.Column{WhatsAppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "whats_apps_profissionals_whatsapps",
				Columns:    []*schema.Column{WhatsAppsColumns[3]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EspecializacaoProfissionaisColumns holds the columns for the "especializacao_profissionais" table.
	EspecializacaoProfissionaisColumns = []*schema.Column{
		{Name: "especializacao_id", Type: field.TypeInt},
		{Name: "profissional_id", Type: field.TypeInt},
	}
	// EspecializacaoProfissionaisTable holds the schema information for the "especializacao_profissionais" table.
	EspecializacaoProfissionaisTable = &schema.Table{
		Name:       "especializacao_profissionais",
		Columns:    EspecializacaoProfissionaisColumns,
		PrimaryKey: []*schema.Column{EspecializacaoProfissionaisColumns[0], EspecializacaoProfissionaisColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "especializacao_profissionais_especializacao_id",
				Columns:    []*schema.Column{EspecializacaoProfissionaisColumns[0]},
				RefColumns: []*schema.Column{EspecializacaosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "especializacao_profissionais_profissional_id",
				Columns:    []*schema.Column{EspecializacaoProfissionaisColumns[1]},
				RefColumns: []*schema.Column{ProfissionalsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AreaSaudesTable,
		ConveniosTable,
		EspecializacaosTable,
		FotosTable,
		PodcastsTable,
		ProfissionalsTable,
		TratamentosTable,
		VideosTable,
		WhatsAppsTable,
		EspecializacaoProfissionaisTable,
	}
)

func init() {
	ConveniosTable.ForeignKeys[0].RefTable = ProfissionalsTable
	EspecializacaosTable.ForeignKeys[0].RefTable = AreaSaudesTable
	FotosTable.ForeignKeys[0].RefTable = ProfissionalsTable
	PodcastsTable.ForeignKeys[0].RefTable = ProfissionalsTable
	TratamentosTable.ForeignKeys[0].RefTable = ProfissionalsTable
	VideosTable.ForeignKeys[0].RefTable = ProfissionalsTable
	WhatsAppsTable.ForeignKeys[0].RefTable = ProfissionalsTable
	EspecializacaoProfissionaisTable.ForeignKeys[0].RefTable = EspecializacaosTable
	EspecializacaoProfissionaisTable.ForeignKeys[1].RefTable = ProfissionalsTable
}
