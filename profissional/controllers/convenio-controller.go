package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/ent"
	"profissional/ent/convenio"
	profissionalQuery "profissional/ent/profissional"
	"strconv"
)

type ConvenioController struct {
	Client *ent.Client
}

// BuscarConveniosPorIdProfissional godoc
// @Summary      Busca os convenios do profissional
// @Description  Buscar os convenios do profisional por id
// @Tags         Convenio
// @Produce      json
// @Param        id   path      int       true  "Id do profissional"
// @Success      200           {object}  []string  "Convenios"
// @Router       /admin/profissionais/{id}/convenios [get]
func (controller ConvenioController) BuscarConveniosPorIdProfissional(httpContext *gin.Context) {
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

	httpContext.JSON(http.StatusOK, data)
}

// BuscarConveniosPorUrlAmigavelProfissional godoc
// @Summary      Busca os convenios do profissional
// @Description  Buscar os convenios do profisional por Url Amigavel
// @Tags         Convenio
// @Produce      json
// @Param        url_amigavel  path      string    true  "Url Amigavel do profissional"
// @Success      200  {object}  []string  "Convenios"
// @Router       /profissionais/{url_amigavel}/convenios [get]
func (controller ConvenioController) BuscarConveniosPorUrlAmigavelProfissional(httpContext *gin.Context) {
	urlAmigavel := httpContext.Param("url_amigavel")

	convenios, err := controller.Client.Convenio.Query().
		Where(convenio.HasProfissionalWith(profissionalQuery.URLAmigavel(urlAmigavel))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []string

	for _, c := range convenios {
		data = append(data, c.Nome)
	}

	httpContext.JSON(http.StatusOK, data)
}
