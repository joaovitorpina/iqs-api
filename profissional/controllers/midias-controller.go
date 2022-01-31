package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/ent"
	"profissional/ent/foto"
	"profissional/ent/podcast"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/video"
	"profissional/mappers/fotos-mapper"
	podcast_mapper "profissional/mappers/podcast-mapper"
	videos_mapper "profissional/mappers/videos-mapper"
	"strconv"
)

type MidiasController struct {
	Client *ent.Client
}

// BuscarFotosPorProfissional godoc
// @Summary      Busca as fotos do profissional
// @Description  Buscar as fotos do profisional por id
// @Tags         Midias
// @Produce      json
// @Param        id   path      int                         true  "Id do profissional"
// @Success      200  {object}  []dtos.BuscarFotosResponse  "Fotos"
// @Router       /profissionais/{id}/midias/fotos [get]
func (controller MidiasController) BuscarFotosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fotos, err := controller.Client.Foto.Query().
		Where(foto.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, fotos_mapper.ToDomain(fotos))
}

// BuscarVideosPorProfissional godoc
// @Summary      Busca os videos do profissional
// @Description  Buscar os videos do profisional por id
// @Tags         Midias
// @Produce      json
// @Param        id   path      int                          true  "Id do profissional"
// @Success      200  {object}  []dtos.BuscarVideosResponse  "Videos"
// @Router       /profissionais/{id}/midias/videos [get]
func (controller MidiasController) BuscarVideosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	videos, err := controller.Client.Video.Query().
		Where(video.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, videos_mapper.ToDomain(videos))
}

// BuscarPodcastsPorProfissional godoc
// @Summary      Busca os podcasts do profissional
// @Description  Buscar os podcasts do profisional por id
// @Tags         Midias
// @Produce      json
// @Param        id   path      int                            true  "Id do profissional"
// @Success      200  {object}  []dtos.BuscarPodcastsResponse  "Podcasts"
// @Router       /profissionais/{id}/midias/podcasts [get]
func (controller MidiasController) BuscarPodcastsPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	podcasts, err := controller.Client.Podcast.Query().
		Where(podcast.HasProfissionalWith(profissionalQuery.ID(id))).
		All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, podcast_mapper.ToDomain(podcasts))
}
