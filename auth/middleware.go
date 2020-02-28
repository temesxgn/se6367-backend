package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	"github.com/square/go-jose/jwt"
	ala "github.com/temesxgn/se6367-backend/alexa"
)

// Middleware - validates request and loads user account into the context
func Middleware() echo.MiddlewareFunc {
	var builder ala.SSMLBuilder
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := new(alexa.Request)
			if err := c.Bind(request); err != nil {
				builder.Say("Sorry there was an issue processing your request.")
				builder.Pause("100")
				builder.Say("Please try again")
				return c.JSON(http.StatusOK, ala.NewSSMLResponse("Authentication Error", builder.Build()))
			}

			c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.Session.User.AccessToken))
			usr, err := getUserFromToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				builder.Say("Error authenticating, please try again.")
				builder.Pause("100")
				builder.Say("If this issue continues please check your status on the alexa app")
				fmt.Println("ERROR getting user from token: " + err.Error())
				return c.JSON(http.StatusOK, ala.NewSSMLResponse("Authentication Error", builder.Build()))
			}

			if !usr.IsValid() {
				builder.Say("Error authenticating, please login again")
				return c.JSON(http.StatusOK, ala.NewSSMLResponse("Authentication Error", builder.Build()))
			}

			c.Set("user", usr)
			c.Set("request", request)
			return next(c)
		}
	}
}

// func Tokenify(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		service := NewService()
// 		token, err := service.GetToken()
// 		if err != nil {
// 			return err
// 		}

// 		c.Set("token", token)
// 		return nil
// 	}
// }

// func LoadUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Request()
// 		service := NewService()
// 		fmt.Println("-------------------- USER OBJECT: --------------------")
// 		fmt.Println(fmt.Sprintf("%v", c.Get("user")))
// 		fmt.Println("-------------------- USER OBJECT: --------------------")
// 		usr := c.Get("user").(string)
// 		user, _ := service.GetUser(usr)
// 		//usr.email = user.Email

// 		var passwordless bool
// 		for _, id := range user.Identities {
// 			if auth0.StringValue(id.Provider) == "email" {
// 				passwordless = true
// 				break
// 			}
// 		}

// 		if !passwordless {
// 			fmt.Println("Provisioning email passwordless user for " + fmt.Sprintf("%v", usr))
// 			_ = service.CreateUser("email", "temesxgn@gmail.com")

// 		}

// 		return nil
// 	}

// }

func getUserFromToken(token string) (*User, error) {
	if split := strings.Split(token, " "); len(split) == 2 {
		token := split[1]
		tkn, err := jwt.ParseSigned(token)
		if err != nil {
			return nil, err
		}

		usr := &User{}
		if err := tkn.UnsafeClaimsWithoutVerification(usr); err != nil {
			return nil, err
		}

		return usr, nil
	}

	return nil, errors.New("malformed authentication request")
}
