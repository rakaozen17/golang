package userService

import (
	"net/http"
	userRouter "restproject/api"
	apiError "restproject/api/err"
	"restproject/app"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *UserRepository
}

func RegisterRoute(app *app.App, e *echo.Echo) {
	userRouter.Route(e, InitUserService(app))
}

func InitUserService(app *app.App) *UserService {
	return &UserService{
		repository: InitRepository(app.DbWrite(), app.DbRead()),
	}
}

// Implements the GetPrice method of the Crypto interface.
func (cs *UserService) RegisterUser(c echo.Context) error {
	errs := apiError.InitErr()
	users := new(UserModel)
	if err := c.Bind(users); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	users.Password = string(hashedPassword)
	errors := cs.repository.RegisterAccount(*users)
	if errors != nil {
		errs.Code = 4040026
		errs.Message = "Gagal saat mendaftarkan user dengan error - Invalid Data"
		return c.JSON(http.StatusBadRequest, errs)
	}

	return c.JSON(http.StatusCreated, users)
}

// Implements the GetVolume method of the Crypto interface.
func (cs *UserService) Login(c echo.Context) error {

	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// check username and password
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
	t, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create token")
	}
	res := &LoginResponse{
		Token: t,
	}
	return c.JSON(http.StatusOK, res)

}
