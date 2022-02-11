package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/dtos"
	"profissional/ent"
)

type AreaSaudeController struct {
	Client *ent.Client
}

// Listar godoc
// @Summary      Lista todas os tipos de profissionais
// @Description  Retorna uma listagem com todas os tipos de profissionais com as especializacoes
// @Tags         Profissional
// @Produce      json
// @Success      200  {object}  []dtos.ListarAreasSaudeResponse  "Tipos de Profissionais"
// @Router       /tipos-profissional [get]
func (controller AreaSaudeController) Listar(httpContext *gin.Context) {
	areasSaude, err := controller.Client.AreaSaude.Query().WithEspecializacoes().All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dtos.ListarAreasSaudeResponse

	for _, areaSaude := range areasSaude {
		response = append(response, dtos.ListarAreasSaudeResponse{
			Id:              areaSaude.ID,
			Descricao:       areaSaude.Descricao,
			Especializacoes: modelToDomain(areaSaude.Edges.Especializacoes),
		})
	}

	httpContext.JSON(http.StatusOK, response)
}

func modelToDomain(model []*ent.Especializacao) []dtos.Especializacao {
	var domain []dtos.Especializacao
	for _, especializacao := range model {
		domain = append(domain, dtos.Especializacao{
			Id:        especializacao.ID,
			Descricao: especializacao.Descricao,
		})
	}

	return domain
}
