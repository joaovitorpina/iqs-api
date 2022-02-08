package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"materia/dtos"
	"materia/ent"
	"materia/ent/categoria"
	"materia/ent/materia"
	"materia/mappers"
	"net/http"
)

type MateriasController struct {
	Client *ent.Client
}

// ListarReduzido godoc
// @Summary      Listagem de materias pela busca
// @Description  Realiza a listagem com paginacão das materias a partir dos dados enviados para busca
// @Tags         Materia
// @Produce      json
// @Param        parametros  query     dtos.ListarMateriasQuery     true  "Parametros"
// @Success      200         {object}  dtos.ListarMateriasResponse  "Listagem Materias"
// @Router       /materias [get]
func (controller MateriasController) ListarReduzido(httpContext *gin.Context) {
	var query dtos.ListarMateriasQuery
	err := httpContext.BindQuery(query)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryDb := controller.Client.Materia.Query().
		WithProfissionalMaterias().
		Where(materia.Status(true))

	if query.Limite == 0 {
		query.Limite = 5
	}

	if query.Pagina == 0 {
		query.Pagina = 1
	}

	queryDb = queryDb.Limit(query.Limite).Offset((query.Pagina - 1) * query.Limite)

	if query.Titulo != "" {
		queryDb = queryDb.Where(materia.TituloContains(query.Titulo))
	}

	if len(query.Categorias) > 0 {
		queryDb = queryDb.Where(materia.HasCategoriaWith(categoria.IDIn(query.Categorias...)))
	}

	data, err := queryDb.All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := controller.Client.Materia.Query().Where(materia.Status(true)).Count(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseBody := dtos.ListarMateriasResponse{
		Data:       []dtos.MateriaReduzida{},
		Pagina:     query.Pagina,
		Quantidade: len(data),
		Total:      total,
	}

	for _, materiaDb := range data {
		responseBody.Data = append(responseBody.Data, dtos.MateriaReduzida{
			Titulo:        materiaDb.Titulo,
			UnidadeId:     materiaDb.UnidadeID,
			UrlAmigavel:   materiaDb.URLAmigavel,
			ImagemUrl:     materiaDb.ImagemURL,
			Profissionais: mappers.ToDomain(materiaDb.Edges.ProfissionalMaterias),
			Texto:         materiaDb.Texto,
		})
	}

	httpContext.JSON(http.StatusOK, responseBody)
}

// BuscarPorUrlAmigavel godoc
// @Summary      Busca todas as informacões da materia
// @Description  Retorna todos os detalhes da materia pelo Url Amigavel
// @Tags         Materia
// @Produce      json
// @Param        url_amigavel  path      string                      true  "Url Amigavel da Materia"
// @Success      200           {object}  dtos.BuscarMateriaResponse  "Materia"
// @Router       /materias/{url_amigavel} [get]
func (controller MateriasController) BuscarPorUrlAmigavel(httpContext *gin.Context) {
	urlAmigavel := httpContext.Param("url_amigavel")

	materiaDb, err := controller.Client.Materia.Query().
		WithProfissionalMaterias().
		Where(materia.URLAmigavel(urlAmigavel)).
		Only(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, dtos.BuscarMateriaResponse{
		Titulo:        materiaDb.Titulo,
		UnidadeId:     materiaDb.UnidadeID,
		UrlAmigavel:   materiaDb.URLAmigavel,
		ImagemUrl:     materiaDb.ImagemURL,
		Profissionais: mappers.ToDomain(materiaDb.Edges.ProfissionalMaterias),
		Texto:         materiaDb.Texto,
		Fonte:         materiaDb.Fonte,
		LinkFonte:     materiaDb.LinkFonte,
		Patrocinado:   materiaDb.Patrocinado,
	})
}
