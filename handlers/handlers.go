package handlers

import (
	"net/http"
	"retroPcPartTracker/store"
	"retroPcPartTracker/templates"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	store *store.PartStore
}

func NewHandlers(s *store.PartStore) *Handlers {
	return &Handlers{store: s}
}

// Helper function to render templ components
func render(c *gin.Context, component templ.Component) {
	c.Writer.Header().Set("Content-Type", "text/html")
	component.Render(c.Request.Context(), c.Writer)
}

func (h *Handlers) HandleHome(c *gin.Context) {
	parts, err := h.store.GetAllParts()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading parts")
		return
	}

	render(c, templates.Home(parts))
}

func (h *Handlers) HandleSearch(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	parts, err := h.store.SearchParts(query)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error searching parts")
		return
	}

	render(c, templates.SearchResults(query, parts))
}

func (h *Handlers) HandlePartsByType(c *gin.Context) {
	partType := c.Param("type")

	parts, err := h.store.GetPartsByType(partType)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading parts")
		return
	}

	render(c, templates.PartsByType(partType, parts))
}

func (h *Handlers) HandlePartDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid part ID")
		return
	}

	part, prices, err := h.store.GetPartWithPrices(id)
	if err != nil {
		c.String(http.StatusNotFound, "Part not found")
		return
	}

	render(c, templates.PartDetail(part, prices))
}
