package main

import (
	"net/http"

	"github.com/brandonbyrd6/link-service/pkg/config"
	"github.com/brandonbyrd6/link-service/pkg/handlers"
	"github.com/brandonbyrd6/link-service/pkg/middleware"
	"github.com/brandonbyrd6/link-service/pkg/repo"
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func init() {
	// logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	// if err != nil {
	// 	logLevel = log.InfoLevel
	// }

	logLevel := log.InfoLevel
	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	cfg := config.GetConfig()

	//DB := db.Init()
	//TODO: Repositories ? Not sure yet
	//? Services, Serperation of Concerns Business logic v. Data v. Handling
	c := utils.NewCounter(0, 100000, 100000)
	//c.Reset()
	s := utils.NewShortener(c)

	r := repo.NewMemoryRepository(s)

	h := handlers.NewHandler(r)

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
		Addr:    cfg.Server.Addr + ":" + cfg.Server.Port,
		Handler: router,
	}

	srv.ListenAndServe()
}
