package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"profissional/ent/foto"
	"profissional/ent/podcast"
	profissionalQuery "profissional/ent/profissional"
	"profissional/ent/video"
	"profissional/mappers/fotos-mapper"
	podcast_mapper "profissional/mappers/podcast-mapper"
	videos_mapper "profissional/mappers/videos-mapper"
	"profissional/services"
	"strconv"
)

func BuscarFotosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fotos, err := services.DbClient.Foto.Query().Where(foto.HasProfissionalWith(profissionalQuery.ID(id))).All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, fotos_mapper.ToDomain(fotos))
}

func BuscarVideosPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	videos, err := services.DbClient.Video.Query().Where(video.HasProfissionalWith(profissionalQuery.ID(id))).All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, videos_mapper.ToDomain(videos))
}

func BuscarPodcastsPorProfissional(httpContext *gin.Context) {
	id, err := strconv.Atoi(httpContext.Param("id"))

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	podcasts, err := services.DbClient.Podcast.Query().Where(podcast.HasProfissionalWith(profissionalQuery.ID(id))).All(context.Background())

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(http.StatusOK, podcast_mapper.ToDomain(podcasts))
}
