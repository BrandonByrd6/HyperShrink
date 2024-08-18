package handlers

import (
	"fmt"
	"net/http"

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

	url, err := h.r.CreateUrl(create.LongUrl, create.UserID)
	if err != nil {
		fmt.Println(err)
	}

	//* 5. Send Data
	c.JSON(http.StatusCreated, url)
}

func (h *Handler) GetUrl(c *gin.Context) {
	short_url := c.Param("short_url")

	url, err := h.r.GetByShortUrl(short_url)
	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusTemporaryRedirect, url.LongUrl)
}

func (h *Handler) DeleteUrl(c *gin.Context) {
	short_url := c.Param("short_url")

	err := h.r.DeleteUrlByShortURL(short_url)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Delete": "Accepted",
	})
}
