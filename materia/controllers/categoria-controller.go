package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"materia/dtos"
	"materia/ent"
	"net/http"
)

type CategoriaController struct {
	Client *ent.Client
}

// Listar godoc
// @Summary      Listagem todas categorias de materias
// @Description  Realiza a listagem de categorias
// @Tags         Materia
// @Produce      json
// @Success      200  {object}  []dtos.ListarCategoriasResponse  "Listagem Categorias"
// @Router       /categorias [get]
func (controller CategoriaController) Listar(httpContext *gin.Context) {
	categorias, err := controller.Client.Categoria.Query().
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dtos.ListarCategoriasResponse

	for _, categoria := range categorias {
		response = append(response, dtos.ListarCategoriasResponse{
			Id:        categoria.ID,
			Descricao: categoria.Descricao,
		})
	}

	httpContext.JSON(http.StatusOK, response)
}
