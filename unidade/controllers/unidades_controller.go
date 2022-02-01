package controllers

import (
	"github.com/gin-gonic/gin"
	"unidade/ent"
)

type UnidadesController struct {
	Client *ent.Client
}

func (controller UnidadesController) ListarUnidades(httpContext *gin.Context) {

}
