package handler

import (
	"social-media-app/features/user"

	"net/http"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv *user.UserService) user.UserHandler {
	return &userControll{
		srv: *srv,
	}
}

func (uc *userControll) RegisterHand() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid input format")
		}

		_, err := uc.srv.RegisterServ(*ToCore(input))
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(http.StatusCreated, "register success"))
	}
}

// func (uc *userControll) LoginHand() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

// func (uc *userControll) ProfileHand() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

// func (uc *userControll) UpdateHand() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

// func (uc *userControll) RemoveHand() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }
