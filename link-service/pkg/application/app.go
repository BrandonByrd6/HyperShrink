package application

import (
	"net/http"

	"github.com/brandonbyrd6/link-service/pkg/config"
	"github.com/brandonbyrd6/link-service/pkg/handlers"
	"github.com/brandonbyrd6/link-service/pkg/middleware"
	"github.com/brandonbyrd6/link-service/pkg/repo"
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Application struct {
	cfg *config.Config
	s   utils.ShortenerInterface
	r   repo.Repository
}

func NewApplication(cfg *config.Config, s utils.ShortenerInterface, r repo.Repository) *Application {
	return &Application{cfg: cfg, s: s, r: r}
}

func (a *Application) Start() {
	//TODO Build Handlers
	//TODO Other stuff

	h := handlers.NewHandler(a.r)

	router := gin.New()
	router.Use(middleware.Logging(), gin.Recovery())

	v1 := router.Group("/api/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "Okay 👍"})
	})

	v1.POST("/", h.CreateUrl)
	v1.GET("/:short_url", h.GetUrl)
	v1.DELETE("/:short_url", h.DeleteUrl)

	srv := &http.Server{
		Addr:    a.cfg.Server.Addr + ":" + a.cfg.Server.Port,
		Handler: router,
	}

	srv.ListenAndServe()
}
