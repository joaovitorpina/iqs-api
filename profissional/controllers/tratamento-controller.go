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

func (controller TratamentosController) ListarTratamentosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Query("id"))

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
