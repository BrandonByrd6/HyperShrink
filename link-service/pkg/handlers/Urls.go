package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brandonbyrd6/link-service/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type CreateRequest struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"` //Optional
	UserID   string `json:"user_id"`
}

func (h *Handler) CreateUrl(c *gin.Context) {
	// I have to initalize the Validator *FACE PALM*
	validate = validator.New(validator.WithRequiredStructEnabled())

	//* 1. Extract Posted Data,
	create := CreateRequest{}

	err := c.Bind(&create)
	if err != nil {
		//c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
	}

	//* 2. Validate data
	errs := validate.Var(create.LongUrl, "required,url")
	if errs != nil {
		//c.AbortWithError(http.StatusBadRequest, errs)
		fmt.Println(err)
	}

	if create.ShortUrl == "" {
		create.ShortUrl = h.shortener.Generate()
	}

	//* Create a new url Struct
	url := models.Url{
		LongUrl:   create.LongUrl,
		ShortUrl:  create.ShortUrl,
		UserId:    create.UserID,
		ExpiresAt: time.Now().Add(time.Duration(time.Duration.Hours(60))),
	}

	if res := h.DB.Create(&url); res.Error != nil {
		fmt.Println(res.Error)
		c.AbortWithError(http.StatusBadRequest, res.Error)
	}

	//* 5. Send Data
	c.JSON(http.StatusCreated, url)
}

func (h *Handler) GetUrl(c *gin.Context) {
	short_url := c.Param("short_url")
	url := models.Url{}

	res := h.DB.Where("short_url = ?", short_url).First(&url)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.AbortWithError(http.StatusBadGateway, res.Error)
	}
	c.Redirect(http.StatusTemporaryRedirect, url.LongUrl)
}

func (h *Handler) DeleteUrl(c *gin.Context) {
	short_url := c.Param("short_url")
	url := models.Url{}

	res := h.DB.Where("short_url = ?", short_url).Delete(&url)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.AbortWithError(http.StatusBadGateway, res.Error)
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Delete": "Accepted",
	})
}
