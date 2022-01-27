package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/dtos"
	"profissional/ent"
	"profissional/ent/areasaude"
	"profissional/ent/especializacao"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/whatsapp"
	"profissional/mappers/especializacao-mapper"
	"strconv"
)

type ProfissionaisController struct {
	Client *ent.Client
}

func (controller ProfissionaisController) ListarProfissionaisReduzido(httpContext *gin.Context) {
	var form dtos.BuscarListagemProfissionaisQuery
	err := httpContext.BindQuery(&form)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := controller.Client.Profissional.Query().
		WithEspecializacoes(func(especializacaoQuery *ent.EspecializacaoQuery) {
			especializacaoQuery.WithAreasaude()
		}).WithWhatsapps(func(whatsAppQuery *ent.WhatsAppQuery) {
		whatsAppQuery.Where(whatsapp.Principal(true)).Limit(1)
	})

	if form.Limite >= 0 {
		form.Limite = 50
	}

	if form.Pagina >= 0 {
		form.Pagina = 1
	}

	query = query.Limit(form.Limite).Offset((form.Pagina - 1) * form.Limite)

	if form.UnidadeId != 0 {
		query = query.Where(profissionalQuery.UnidadeID(form.UnidadeId))
	}

	if form.Nome != "" {
		query = query.Where(profissionalQuery.NomeContains(form.Nome))
	}

	if form.TipoProfissionalId != 0 {
		query = query.Where(profissionalQuery.HasEspecializacoesWith(especializacao.HasAreasaudeWith(areasaude.ID(form.TipoProfissionalId))))
	}

	if form.EspecialidadeId != 0 {
		query = query.Where(profissionalQuery.HasEspecializacoesWith(especializacao.ID(form.EspecialidadeId)))
	}

	data, err := query.All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseBody := dtos.BuscarListagemProfissionaisResponse{Data: []dtos.ProfissionalReduzido{}, Pagina: form.Pagina, Quantidade: len(data)}

	for _, profissionalDb := range data {
		profissionalResponse := dtos.ProfissionalReduzido{
			Id:              profissionalDb.ID,
			Nome:            profissionalDb.Nome,
			UrlAmigavel:     profissionalDb.URLAmigavel,
			Tipo:            profissionalDb.Edges.Especializacoes[0].Edges.Areasaude.Descricao,
			Especialidades:  especializacao_mapper.ToDomain(profissionalDb.Edges.Especializacoes),
			Recomendado:     profissionalDb.Recomendado,
			ImagemPerfilUrl: profissionalDb.ImagemPerfilURL,
			UnidadeId:       profissionalDb.UnidadeID,
			Facebook:        profissionalDb.Facebook,
			WhatsApp:        profissionalDb.Edges.Whatsapps[0].Numero,
			Instagram:       profissionalDb.Instagram,
			Email:           profissionalDb.Email,
			Site:            profissionalDb.Site,
		}

		responseBody.Data = append(responseBody.Data, profissionalResponse)
	}

	httpContext.JSON(http.StatusOK, responseBody)
}

func (controller ProfissionaisController) BuscarProfissionalPorId(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profissional, err := controller.Client.Profissional.Query().
		Where(profissionalQuery.ID(id)).
		WithEspecializacoes().
		First(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, dtos.BuscarProfissionalResponse{
		Id:                  profissional.ID,
		Nome:                profissional.Nome,
		ImagemPerfilUrl:     profissional.ImagemPerfilURL,
		EnderecoId:          profissional.EnderecoID,
		Tipo:                profissional.Edges.Especializacoes[0].Edges.Areasaude.Descricao,
		Especialidades:      especializacao_mapper.ToDomain(profissional.Edges.Especializacoes),
		Telefone:            profissional.Telefone,
		Celular:             profissional.Celular,
		Facebook:            profissional.Facebook,
		Instagram:           profissional.Instagram,
		Email:               profissional.Email,
		Site:                profissional.Site,
		Sobre:               profissional.Sobre,
		Conselho:            profissional.Conselho,
		NumeroIdentificacao: profissional.NumeroIdentificacao,
	})
}
