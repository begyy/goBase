package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"goBase/app/repositories"
	"goBase/app/schema"
	"goBase/app/services"
	"goBase/app/utils"
	"net/http"
)

// SignUp
// @Summary      Create User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body schema.SignUpSchemaIn true "User SignUp data"
// @Success      200  {object}  schema.SignUpSchemaIn
// @Failure      400  {object}   utils.LogicError
// @Failure      422  {array}   utils.ValidationError
// @Router       /user/sign-up/ [post]
func SignUp(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	userDTO := new(schema.SignUpSchemaIn)
	if err := c.Bind(userDTO); err != nil {
		return err
	}

	if errors := utils.ValidateAndFormat(userDTO); errors != nil {
		return c.JSON(http.StatusBadRequest, errors)
	}
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	message, err := userService.SignUp(userDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": message})
}

// SignIn
// @Summary      Login User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body schema.SignUpSchemaIn true "User SignUp data"
// @Success      200  {object}  schema.SignUpSchemaIn
// @Failure      400  {object}   utils.LogicError
// @Failure      422  {array}   utils.ValidationError
// @Router       /user/sign-up/ [post]
func SignIn(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	userDTO := new(schema.SignInSchemaIn)
	if err := c.Bind(userDTO); err != nil {
		return err
	}
	if errors := utils.ValidateAndFormat(userDTO); errors != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userID, err := userService.SignIn(userDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	tokenRepo := repositories.NewUserTokenRepository(db)
	tokenService := services.NewUserTokenService(tokenRepo)
	token, err := tokenService.GetTokenOrAddToken(*userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

// UserMe
// @Summary      Me
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  schema.UserMeSchema
// @Failure      400  {object}   utils.LogicError
// @Header       200, 400, default  {string}  Authorization "token"
// @Router       /user/me/ [get]
func UserMe(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	userID := c.Get("userID").(int)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	user, err := userService.UserMe(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)

}
