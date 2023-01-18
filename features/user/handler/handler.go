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

func (uc *userControll) LoginHand() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginReq{}
		if err := c.Bind(&input); err != nil {
			// log.Println("Bind error", err.Error())
			return c.JSON(http.StatusBadRequest, "Invalid input format") //http.StatusBadRequest bisa diganti dengan 400
		}

		token, res, err := uc.srv.LoginServ(input.Email, input.Password)
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(http.StatusOK, "login success", res, token))
	}
}

func (uc *userControll) ProfileHand() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := uc.srv.ProfileServ(token)
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(http.StatusOK, "berhasil lihat profil", res))
	}
}

func (uc *userControll) UpdateHand() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := UpdateReq{}
		if err := c.Bind(&input); err != nil {
			// log.Println("Bind error", err.Error())
			return c.JSON(http.StatusBadRequest, "Invalid input format")
		}
		
		res, err := uc.srv.UpdateServ(token, *ToCore(input))
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(http.StatusOK, "update success", res))
	}
}

func (uc *userControll) RemoveHand() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		err := uc.srv.RemoveServ(token)
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, "Akun berhasil dihapus")
	}
}
