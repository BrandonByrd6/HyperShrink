package main

import (
	"net/http"

	"github.com/brandonbyrd6/link-service/pkg/config"
	"github.com/brandonbyrd6/link-service/pkg/db"
	"github.com/brandonbyrd6/link-service/pkg/handlers"
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	DB := db.Init()
	//TODO: Repositories ? Not sure yet
	//? Services, Serperation of Concerns Business logic v. Data v. Handling
	c := utils.NewCounter(1, 100000, 100000)
	c.Reset()
	s := utils.NewShortener(c)

	h := handlers.NewHandler(DB, s)

	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

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
