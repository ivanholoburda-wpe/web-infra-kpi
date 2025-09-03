package handlers

import (
	"api-service/internal/dto"
	"api-service/internal/models"
	"api-service/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteHandler struct {
	service pkg.SiteService
}

func NewSiteHandler(service pkg.SiteService) *SiteHandler {
	return &SiteHandler{
		service: service,
	}
}

func (h *SiteHandler) Create(c *gin.Context) {
	var req dto.CreateSiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	siteModel := models.Site{
		Name: req.Name,
		Url:  req.URL,
	}

	createdSite, err := h.service.Create(c.Request.Context(), &siteModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapModelToResponse(createdSite))
}

func (h *SiteHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	site, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapModelToResponse(site))
}

func (h *SiteHandler) GetAll(c *gin.Context) {
	sites, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]dto.SiteResponse, 0, len(sites))
	for _, site := range sites {
		responses = append(responses, mapModelToResponse(&site))
	}

	c.JSON(http.StatusOK, responses)
}

func (h *SiteHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func mapModelToResponse(site *models.Site) dto.SiteResponse {
	return dto.SiteResponse{
		ID:           site.ID,
		URL:          site.Url,
		HttpStatus:   site.HttpStatus,
		ResponseTime: site.ResponseTime,
		CreatedAt:    site.CreatedAt,
		UpdatedAt:    site.UpdatedAt,
	}
}
