package handlers

import (
	"fmt"
	"net/http"
	"server/pkg/models"

	"github.com/labstack/echo/v4"
)

func AddPost(c echo.Context) error {
	posts := new(models.Posts)

	err := c.Bind(&posts); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(c.Request())

  	post := models.Posts{
	Content: posts.Content,
	Likes: posts.Likes,
	UserID: posts.UserID,
  	}

	// Add to database
	b := post.AddPost()

	return c.JSON(http.StatusOK, b)
}