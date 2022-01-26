package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/whatsapp"
	whatsappmapper "profissional/mappers/whatsapp-mapper"
	"profissional/services"
	"strconv"
)

func BuscarWhatsappsPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	whatsapps, err := services.DbClient.WhatsApp.Query().Where(whatsapp.HasProfissionalWith(profissionalQuery.ID(id))).All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, whatsappmapper.ToDomain(whatsapps))
}
