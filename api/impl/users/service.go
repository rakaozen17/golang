package userService

import (
	"fmt"
	"net/http"
	userRouter "restproject/api/Router/users"
	apiRes "restproject/api/helper"
	"restproject/app"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *UserRepository
}

func RegisterRoute(app *app.App, e *echo.Echo, dbr *sqlx.DB, dbw *sqlx.DB) {
	userRouter.Route(e, InitUserService(app, dbr, dbw))
}

func InitUserService(app *app.App, dbr *sqlx.DB, dbw *sqlx.DB) *UserService {
	return &UserService{
		repository: InitRepository(dbr, dbw),
	}
}

func (cs *UserService) RegisterUser(c echo.Context) error {
	res := apiRes.InitRes()
	users := new(UserModel)
	if err := c.Bind(users); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		res.ResponseCode = 500
		res.ResponseMessage = fmt.Sprintf("Gagal saat mendaftarkan user dengan error - %v", err)
		res.Data = nil
		return err
	}
	users.Password = string(hashedPassword)
	errors := cs.repository.RegisterAccount(*users)
	if errors != nil {
		res.ResponseCode = 500
		res.ResponseMessage = fmt.Sprintf("Gagal saat mendaftarkan user dengan error - %v", err)
		res.Data = nil
		return c.JSON(http.StatusBadRequest, res)
	}

	res.ResponseCode = 200
	res.ResponseMessage = "Sucessfully Processed"
	res.Data = users

	return c.JSON(http.StatusCreated, res)
}

func (cs *UserService) Auth(c echo.Context) error {

	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	res := apiRes.InitRes()
	user, err := cs.repository.FindUserByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid username or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(viper.GetString("jwt_secretkey")))
	if err != nil {
		res.ResponseCode = 404
		res.ResponseMessage = "Users not found"
		res.Data = nil
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.ResponseCode = 200
	res.ResponseMessage = "Successfully logged in"
	res.Data = LoginResponse{
		Token:     t,
		ExpiredIn: time.Now().Add(time.Hour * 24).Unix(),
	}
	return c.JSON(http.StatusOK, res)
}
