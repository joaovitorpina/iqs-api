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

func (controller WhatsappController) BuscarWhatsappsPorProfissional(httpContext *gin.Context) {
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
