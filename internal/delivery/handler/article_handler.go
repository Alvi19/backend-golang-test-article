package http

import (
	"net/http"
	"strconv"

	"github.com/Alvi19/backend-golang-test/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	uc        usecase.ArticleUsecase
	validator *validator.Validate
}

func NewArticleHandler(uc usecase.ArticleUsecase, v *validator.Validate) *ArticleHandler {
	return &ArticleHandler{uc: uc, validator: v}
}

// Register all routes
func (h *ArticleHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/article", h.CreateArticle)
	g.GET("/article/:limit/:offset", h.ListArticles)
	g.GET("/article/:id", h.GetArticleByID)
	g.PUT("/article/:id", h.UpdateArticle)
	g.DELETE("/article/:id", h.DeleteArticle)
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with title, content, category, and status
// @Tags Articles
// @Accept json
// @Produce json
// @Param request body usecase.CreateArticleRequest true "Article Data"
// @Success 201 {object} domain.Article
// @Failure 400 {object} map[string]interface{}
// @Router /v1/article [post]
func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	var req usecase.CreateArticleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	article, err := h.uc.Create(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, article)
}

// ListArticles godoc
// @Summary Get list of articles
// @Description Get paginated list of articles using limit and offset
// @Tags Articles
// @Produce json
// @Param limit path int true "Limit"
// @Param offset path int true "Offset"
// @Success 200 {array} domain.Article
// @Router /v1/article/{limit}/{offset} [get]
func (h *ArticleHandler) ListArticles(c echo.Context) error {
	limit, err1 := strconv.Atoi(c.Param("limit"))
	offset, err2 := strconv.Atoi(c.Param("offset"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid limit or offset"})
	}

	articles, err := h.uc.List(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, articles)
}

// GetArticleByID godoc
// @Summary Get article by ID
// @Description Get single article by its ID
// @Tags Articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} domain.Article
// @Failure 404 {object} map[string]interface{}
// @Router /v1/article/{id} [get]
func (h *ArticleHandler) GetArticleByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article id"})
	}

	article, err := h.uc.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "article not found"})
	}

	return c.JSON(http.StatusOK, article)
}

// UpdateArticle godoc
// @Summary Update an article
// @Description Update an article's title, content, category, or status
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param request body usecase.UpdateArticleRequest true "Update Article"
// @Success 200 {object} domain.Article
// @Router /v1/article/{id} [put]
func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article id"})
	}

	var req usecase.UpdateArticleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	article, err := h.uc.Update(c.Request().Context(), uint(id), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article by its ID
// @Tags Articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]string
// @Router /v1/article/{id} [delete]
func (h *ArticleHandler) DeleteArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article id"})
	}

	if err := h.uc.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "article deleted successfully"})
}
