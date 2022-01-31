package controllers

import (
	"context"
	"endereco/dtos"
	"endereco/ent"
	enderecoQuery "endereco/ent/endereco"
	"endereco/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BuscarEnderecoPorId godoc
// @Summary      Busca um endereco por id
// @Description  Retorna o endereco pelo id enviado
// @Tags         Endereco
// @Produce      json
// @Param        id   path      int                          true  "Id do Endereco"
// @Success      200  {object}  dtos.BuscarEnderecoResponse  "Endereco"
// @Router       /enderecos/{id} [get]
func BuscarEnderecoPorId(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	endereco, err := services.DbClient.Endereco.Query().Where(enderecoQuery.ID(id)).WithCep(func(cepQuery *ent.CepQuery) {
		cepQuery.WithCidade(func(cidadeQuery *ent.CidadeQuery) {
			cidadeQuery.WithEstado()
		})
	}).First(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, dtos.BuscarEnderecoResponse{
		Logradouro: endereco.Edges.Cep.Logradouro,
		Bairro:     endereco.Edges.Cep.Bairro,
		Numero:     endereco.Numero,
		Cep:        endereco.Edges.Cep.ID,
		Cidade:     endereco.Edges.Cep.Edges.Cidade.Nome,
		Estado:     endereco.Edges.Cep.Edges.Cidade.Edges.Estado.Nome,
	})
}
