package handler

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"api/db"
	"api/graphql"
	"api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "It works.")
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		db := db.ConnectGORM()
		user := []models.User{}
		db.Find(&user, "name=? and password=?", username, password)

		if len(user) > 0 && username == user[0].Name {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admimn"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
		return echo.ErrUnauthorized
	}
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		_ = user.Claims.(jwt.MapClaims)
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(c.Request().Body)
		query := bufBody.String()
		log.Printf(query)
		result := graphql.ExecuteQuery(query)

		return c.JSON(http.StatusOK, result)
	}
}
