package controllers

import (
	"net/http"
	"os"
	"strings"

	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
var CreateAccount = func(ctx echo.Context) error {

	account := &models.Account{}
	err := ctx.Bind(account)
	// err := json.NewDecoder(ctx.Request().Body).Decode(account)
	if err != nil {
		// u.Respond(ctx, u.Message(false, "Invalid request"))
		return err
	}

	resp := account.Create()
	// u.Respond(w, resp)
	return ctx.JSON(http.StatusOK, resp)
}

var Authenticate = func(ctx echo.Context) error {
	// user := ctx.Get("id").(*jwt.Token)
	tok := ctx.Request().Header.Get("Authorization")
	tokenPart := strings.Split(tok, " ")[1]
	tk := &models.Token{}
	// tokenPart := user.Claims.(*models.Token)
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})
	if err != nil {
		// u.Respond(w, u.Message(false, "Invalid request"))
		return err
	}
	if !token.Valid {
		return ctx.JSON(http.StatusForbidden, "")
	}
	account := &models.Account{}
	er := ctx.Bind(account)
	if er != nil {
		// u.Respond(w, u.Message(false, "Invalid request"))
		return er
	}
	resp := models.Login(account.Email, account.Password)
	return ctx.JSON(http.StatusOK, resp)
}
