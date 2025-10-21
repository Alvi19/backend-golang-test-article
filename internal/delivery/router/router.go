package router

import (
	"github.com/Alvi19/backend-golang-test/internal/config"
	dr "github.com/Alvi19/backend-golang-test/internal/delivery/handler"
	"github.com/Alvi19/backend-golang-test/internal/repository"
	"github.com/Alvi19/backend-golang-test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = false
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v := validator.New()

	articleRepo := repository.NewArticleRepository(db)
	articleUC := usecase.NewArticleUsecase(articleRepo)

	ah := dr.NewArticleHandler(articleUC, v)

	api := e.Group("/api/v1")
	ah.RegisterRoutes(api)

	return e
}
