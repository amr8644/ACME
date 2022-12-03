package handlers

import (
	"fmt"
	"net/http"
	"server/pkg/models"
	"server/pkg/utils"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := session.Get("session", c)
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		  }
		_, ok := session.Values["userID"]

		if strings.Split(c.Path(),"/")[1] == "login" ||
		 strings.Split(c.Path(),"/")[1] == "register"{
			return next(c)
		}
		
		if !ok {
			http.Redirect(c.Response(),c.Request(),"/",http.StatusFound)
			return c.String(http.StatusForbidden, "Unauthorized")
		}
		if session.Values["authenticated"] == false{
			http.Redirect(c.Response(),c.Request(),"/",http.StatusFound)
			return c.String(http.StatusForbidden, "Unauthorized")
		}
		return next(c)
	}
}

func RegisterUser(c echo.Context) error {

	u := new(models.Users)

	err := c.Bind(&u); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}	
	
	// Hash password
	hash, _ := utils.HashPassword(u.Password)

	utils.ServeFrames(u.Profile)
	fmt.Println(utils.ServeFrames(u.Profile))
  	user := models.Users{
	Fullname: u.Fullname,
    Username: u.Username,
    Password: hash,
    Email: u.Email,
	Profile: utils.ServeFrames(u.Profile),
  	}

	// Add to database
	b := user.RegisterUser()


	// Set user as authenticated
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	  }

	session.Values["userID"] = u.ID
	session.Values["authenticated"] = true

	session.Save(c.Request(), c.Response())


	return c.JSON(http.StatusOK, b)
}

func LoginUser(c echo.Context) error {
	u := new(models.Users)
	
    err := c.Bind(&u); if err!= nil {
        return c.String(http.StatusBadRequest, "bad request")
    }
	user := models.Users{
		Password:u.Password,
	}
	
	b := u.LoginUser()
	
	// Check Password Hash
	match := utils.CheckPasswordHash(user.Password,b.Password)

	if!match {
		return c.String(http.StatusForbidden,"Password mismatch")
	}

	// Set user as authenticated
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
  	}
	session.Values["userID"] = u.ID
	session.Values["authenticated"] = true


	session.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, b)
}

func Logout(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge: -1,
		HttpOnly: true,
	  }

	delete(session.Values,"userID")
	
	session.Values["authenticated"] = false

    // Revoke users authentication
    session.Save(c.Request(),c.Response())
	return c.JSON(http.StatusOK, "Logout")

}

func GetAllUsers(c echo.Context) error {
	allUsers := models.GetAllUsers()
	return c.JSON(http.StatusOK,allUsers)
}

func GetUserById(c echo.Context) error {
	id,err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,"Bad Request")
	}
	user := models.GetUserById(id);
	return c.JSON(http.StatusOK,user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
    deletedUser:= models.DeleteUser(id)
    return c.JSON(http.StatusOK,deletedUser)
}