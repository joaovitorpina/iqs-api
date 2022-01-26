package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/dtos"
	"profissional/ent/convenio"
	profissionalQuery "profissional/ent/profissional"
	"profissional/services"
	"strconv"
)

func BuscarConveniosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	convenios, err := services.DbClient.Convenio.Query().Where(convenio.HasProfissionalWith(profissionalQuery.ID(id))).All(context.Background())

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
