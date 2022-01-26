package controllers

import (
	"context"
	"endereco/dtos"
	"endereco/ent"
	enderecoQuery "endereco/ent/endereco"
	"endereco/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuscarEnderecoPorId(httpContext *gin.Context) {
	var query dtos.BuscarEnderecoQuery
	err := httpContext.BindQuery(&query)

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	endereco, err := services.DbClient.Endereco.Query().Where(enderecoQuery.ID(query.Id)).WithCep(func(cepQuery *ent.CepQuery) {
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
