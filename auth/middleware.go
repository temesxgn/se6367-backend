package auth

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/temesxgn/se6367-backend/config"
	"gopkg.in/auth0.v3"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetAuth0SigningKey()), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	return func(c echo.Context) error {
		return jwtMiddleware.CheckJWT(c.Response().Writer, c.Request())
	}
}

func Tokenify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		service := NewService()
		token, err := service.GetToken()
		if err != nil {
			return err
		}

		c.Set("token", token)
		return nil
	}
}

func LoadUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request()
		service := NewService()
		fmt.Println("-------------------- USER OBJECT: --------------------")
		fmt.Println(fmt.Sprintf("%v", c.Get("user")))
		fmt.Println("-------------------- USER OBJECT: --------------------")
		usr := c.Get("user").(string)
		user, _ := service.GetUser(usr)
		//usr.email = user.Email

		var passwordless bool
		for _, id := range user.Identities {
			if auth0.StringValue(id.Provider) == "email" {
				passwordless = true
				break
			}
		}

		if !passwordless {
			fmt.Println("Provisioning email passwordless user for " + fmt.Sprintf("%v", usr))
			_ = service.CreateUser("email", "temesxgn@gmail.com")

		}

		return nil
	}

}
