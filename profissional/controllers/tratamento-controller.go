package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/ent"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/tratamento"
	"strconv"
)

type TratamentosController struct {
	Client *ent.Client
}

// ListarTratamentosPorIdProfissional godoc
// @Summary      Busca os tratamentos do profissional
// @Description  Buscar os tratamentos do profisional por id
// @Tags         Tratamentos
// @Produce      json
// @Param        id   path      int       true  "Id do profissional"
// @Success      200           {object}  []string  "Tratamentos"
// @Router       /admin/profissionais/{id}/tratamentos [get]
func (controller TratamentosController) ListarTratamentosPorIdProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tratamentos, err := controller.Client.Tratamento.Query().
		Where(tratamento.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []string
	for _, t := range tratamentos {
		response = append(response, t.Descricao)
	}

	httpContext.JSON(http.StatusOK, response)
}

// ListarTratamentosPorUrlAmigavelProfissional godoc
// @Summary      Busca os tratamentos do profissional
// @Description  Buscar os tratamentos do profisional por Url Amigavel
// @Tags         Tratamentos
// @Produce      json
// @Param        url_amigavel  path      string    true  "Url Amigavel do profissional"
// @Success      200  {object}  []string  "Tratamentos"
// @Router       /profissionais/{url_amigavel}/tratamentos [get]
func (controller TratamentosController) ListarTratamentosPorUrlAmigavelProfissional(httpContext *gin.Context) {
	urlAmigavel := httpContext.Param("url_amigavel")

	tratamentos, err := controller.Client.Tratamento.Query().
		Where(tratamento.HasProfissionalWith(profissionalQuery.URLAmigavel(urlAmigavel))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []string
	for _, t := range tratamentos {
		response = append(response, t.Descricao)
	}

	httpContext.JSON(http.StatusOK, response)
}
