package main

import (
	"context"
	"fmt"

	"github.com/brandonbyrd6/link-service/pkg/connections"

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
	//cfg := config.GetConfig()

	kv := connections.Init()
	resp, err := kv.Put(context.Background(), "Test", "Value", nil)

	if err != nil {
		fmt.Println("main")
		fmt.Println(err)
	}

	fmt.Println(resp.PrevKv.Value)
	fmt.Println(resp.Header.String())

	// c := utils.NewCounter(0, 100000, 100000)
	// s := utils.NewShortener(c)
	// r := repo.NewMemoryRepository(s)

	// app := application.NewApplication(cfg, s, r)

	// app.Start()

	//DB := db.Init()

	//c.Reset()

	//r := repo.NewPostgresRepository(DB, s)

	// h := handlers.NewHandler(r)

	// router := gin.New()
	// router.Use(middleware.Logging(), gin.Recovery())

	// v1 := router.Group("/api/v1")

	// v1.GET("/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"health": "Okay 👍"})
	// })

	// v1.POST("/", h.CreateUrl)
	// v1.GET("/:short_url", h.GetUrl)
	// v1.DELETE("/:short_url", h.DeleteUrl)

	// srv := &http.Server{
	// 	Addr:    cfg.Server.Addr + ":" + cfg.Server.Port,
	// 	Handler: router,
	// }

	// srv.ListenAndServe()
}
