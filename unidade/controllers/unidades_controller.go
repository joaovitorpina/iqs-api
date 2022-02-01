package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"unidade/dtos"
	"unidade/ent"
)

type UnidadesController struct {
	Client *ent.Client
}

// ListarUnidades godoc
// @Summary      Lista todas as unidades
// @Description  Retorna uma listagem com todas as unidades
// @Tags         Unidade
// @Produce      json
// @Success      200  {object}  dtos.BuscarUnidadesResponse  "Unidades"
// @Router       /unidades [get]
func (controller UnidadesController) ListarUnidades(httpContext *gin.Context) {
	unidades, err := controller.Client.Unidade.Query().All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dtos.BuscarUnidadesResponse
	for _, unidade := range unidades {
		response = append(response, dtos.BuscarUnidadesResponse{
			Id:          unidade.ID,
			Descricao:   unidade.Descricao,
			UrlAmigavel: unidade.URLAmigavel,
			EnderecoId:  unidade.EnderecoID,
			Latitude:    unidade.Latitude,
			Longitude:   unidade.Longitude,
			Telefone:    unidade.Telefone,
			Celular:     unidade.Celular,
			Email:       unidade.Email,
			Facebook:    unidade.Facebook,
			Instagram:   unidade.Instagram,
			Youtube:     unidade.Youtube,
			Ativo:       unidade.Ativo,
		})
	}

	httpContext.JSON(http.StatusOK, response)
}
