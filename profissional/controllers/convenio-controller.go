package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/dtos"
	"profissional/ent"
	"profissional/ent/convenio"
	profissionalQuery "profissional/ent/profissional"
	"strconv"
)

type ConvenioController struct {
	Client *ent.Client
}

// BuscarConveniosPorProfissional godoc
// @Summary      Busca os convenios do profissional
// @Description  Buscar os convenios do profisional por id
// @Tags         Convenio
// @Produce      json
// @Param        id   path      int                           true  "Id do profissional"
// @Success      200  {object}  dtos.BuscarConveniosResponse  "Convenios"
// @Router       /profissionais/{id}/convenios [get]
func (controller ConvenioController) BuscarConveniosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	convenios, err := controller.Client.Convenio.Query().
		Where(convenio.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []string

	for _, c := range convenios {
		data = append(data, c.Nome)
	}

	httpContext.JSON(http.StatusOK, dtos.BuscarConveniosResponse{
		Data: data,
	})
}
