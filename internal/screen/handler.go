package screen

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.List)
	rg.POST("", h.Create)
	rg.GET("/:id", h.GetByID)
	rg.PUT("/:id", h.Update)
	rg.DELETE("/:id", h.Delete)
}

func (h *Handler) List(c *gin.Context) {
	screens, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if screens == nil {
		screens = []Screen{}
	}
	c.JSON(http.StatusOK, screens)
}

func (h *Handler) ListByAppID(c *gin.Context) {
	screens, err := h.repo.ListByApp(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if screens == nil {
		screens = []Screen{}
	}
	c.JSON(http.StatusOK, screens)
}

func (h *Handler) GetByID(c *gin.Context) {
	sc, err := h.repo.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if sc == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, sc)
}

func (h *Handler) Create(c *gin.Context) {
	var sc Screen
	if err := c.ShouldBindJSON(&sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(c.Request.Context(), &sc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sc)
}

func (h *Handler) Update(c *gin.Context) {
	var sc Screen
	if err := c.ShouldBindJSON(&sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sc.ID = c.Param("id")
	if err := h.repo.Update(c.Request.Context(), &sc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sc)
}

func (h *Handler) Delete(c *gin.Context) {
	if err := h.repo.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
