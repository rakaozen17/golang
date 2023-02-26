package destinationService

import (
	"fmt"
	"net/http"
	DestinationRouter "restproject/api/Router/destination"
	apiRes "restproject/api/helper"
	"restproject/app"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type DestinationService struct {
	repository *DestinationRepository
}

func RegisterRoute(app *app.App, e *echo.Echo, dbr *sqlx.DB, dbw *sqlx.DB) {
	DestinationRouter.Route(e, InitDestinationService(app, dbr, dbw))
}

func InitDestinationService(app *app.App, dbr *sqlx.DB, dbw *sqlx.DB) *DestinationService {
	return &DestinationService{
		repository: InitRepository(dbr, dbw),
	}
}

func (cs *DestinationService) CreateDestination(c echo.Context) error {
	res := apiRes.InitRes()
	des := new(DestinationCreateRequest)
	if err := c.Bind(des); err != nil {
		res.ResponseCode = 403
		res.ResponseMessage = "Invalid Request"
		res.Data = nil
		return c.JSON(http.StatusBadRequest, res)
	}
	errors := cs.repository.CreateDestination(des)
	if errors != nil {
		res.ResponseCode = 500
		res.ResponseMessage = fmt.Sprintf("Gagal saat mendaftarkan destinasi dengan error - %v", errors)
		res.Data = nil
		return c.JSON(http.StatusBadRequest, res)
	}

	res.ResponseCode = 200
	res.ResponseMessage = "Sucessfully Processed"
	res.Data = &des

	return c.JSON(http.StatusCreated, res)
}

func (cs *DestinationService) UpdateDestination(c echo.Context) error {

	// req := new(LoginRequest)
	// if err := c.Bind(req); err != nil {
	// 	return err
	// }
	// res := apiRes.InitRes()
	// user, err := cs.repository.FindUserByUsername(req.Username)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = user.Username
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// t, err := token.SignedString([]byte(viper.GetString("jwt_secretkey")))
	// if err != nil {
	// 	res.ResponseCode = 404
	// 	res.ResponseMessage = "Users not found"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	// res.ResponseCode = 200
	// res.ResponseMessage = "Successfully logged in"
	// res.Data = LoginResponse{
	// 	Token:     t,
	// 	ExpiredIn: time.Now().Add(time.Hour * 24).Unix(),
	// }
	// return c.JSON(http.StatusOK, res)
	return nil
}

func (cs *DestinationService) DeleteDestination(c echo.Context) error {

	// req := new(LoginRequest)
	// if err := c.Bind(req); err != nil {
	// 	return err
	// }
	// res := apiRes.InitRes()
	// user, err := cs.repository.FindUserByUsername(req.Username)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = user.Username
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// t, err := token.SignedString([]byte(viper.GetString("jwt_secretkey")))
	// if err != nil {
	// 	res.ResponseCode = 404
	// 	res.ResponseMessage = "Users not found"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	// res.ResponseCode = 200
	// res.ResponseMessage = "Successfully logged in"
	// res.Data = LoginResponse{
	// 	Token:     t,
	// 	ExpiredIn: time.Now().Add(time.Hour * 24).Unix(),
	// }
	// return c.JSON(http.StatusOK, res)
	return nil
}

func (cs *DestinationService) GetSingleDestination(c echo.Context) error {

	// req := new(LoginRequest)
	// if err := c.Bind(req); err != nil {
	// 	return err
	// }
	// res := apiRes.InitRes()
	// user, err := cs.repository.FindUserByUsername(req.Username)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = user.Username
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// t, err := token.SignedString([]byte(viper.GetString("jwt_secretkey")))
	// if err != nil {
	// 	res.ResponseCode = 404
	// 	res.ResponseMessage = "Users not found"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	// res.ResponseCode = 200
	// res.ResponseMessage = "Successfully logged in"
	// res.Data = LoginResponse{
	// 	Token:     t,
	// 	ExpiredIn: time.Now().Add(time.Hour * 24).Unix(),
	// }
	// return c.JSON(http.StatusOK, res)
	return nil
}

func (cs *DestinationService) GetAllDestination(c echo.Context) error {

	// req := new(LoginRequest)
	// if err := c.Bind(req); err != nil {
	// 	return err
	// }
	// res := apiRes.InitRes()
	// user, err := cs.repository.FindUserByUsername(req.Username)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid username or password")
	// }

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = user.Username
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// t, err := token.SignedString([]byte(viper.GetString("jwt_secretkey")))
	// if err != nil {
	// 	res.ResponseCode = 404
	// 	res.ResponseMessage = "Users not found"
	// 	res.Data = nil
	// 	return c.JSON(http.StatusInternalServerError, res)
	// }

	// res.ResponseCode = 200
	// res.ResponseMessage = "Successfully logged in"
	// res.Data = LoginResponse{
	// 	Token:     t,
	// 	ExpiredIn: time.Now().Add(time.Hour * 24).Unix(),
	// }
	// return c.JSON(http.StatusOK, res)
	return nil
}
