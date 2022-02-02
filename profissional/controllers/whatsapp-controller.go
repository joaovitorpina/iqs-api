package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/ent"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/whatsapp"
	whatsappmapper "profissional/mappers/whatsapp-mapper"
	"strconv"
)

type WhatsappController struct {
	Client *ent.Client
}

// BuscarWhatsappsPorIdProfissional godoc
// @Summary      Busca os WhatsApp`s do profissional
// @Description  Buscar os WhatsApp`s do profisional por id
// @Tags         WhatsApp
// @Produce      json
// @Param        id   path      int                             true  "Id do profissional"
// @Success      200           {object}  []dtos.BuscarWhatsappsResponse  "WhatsApp`s"
// @Router       /admin/profissionais/{id}/whatsapps [get]
func (controller WhatsappController) BuscarWhatsappsPorIdProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	whatsapps, err := controller.Client.WhatsApp.Query().
		Where(whatsapp.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, whatsappmapper.ToDomain(whatsapps))
}

// BuscarWhatsappsPorUrlAmigavelProfissional godoc
// @Summary      Busca os WhatsApp`s do profissional
// @Description  Buscar os WhatsApp`s do profisional por Url Amigavel
// @Tags         WhatsApp
// @Produce      json
// @Param        url_amigavel  path      string                          true  "Url Amigavel do profissional"
// @Success      200  {object}  []dtos.BuscarWhatsappsResponse  "WhatsApp`s"
// @Router       /profissionais/{url_amigavel}/whatsapps [get]
func (controller WhatsappController) BuscarWhatsappsPorUrlAmigavelProfissional(httpContext *gin.Context) {
	urlAmigavel := httpContext.Param("url_amigavel")

	whatsapps, err := controller.Client.WhatsApp.Query().
		Where(whatsapp.HasProfissionalWith(profissionalQuery.URLAmigavel(urlAmigavel))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, whatsappmapper.ToDomain(whatsapps))
}
