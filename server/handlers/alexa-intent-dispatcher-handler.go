package handlers

import (
	"errors"
	"fmt"
	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	"github.com/square/go-jose/jwt"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth"
	"net/http"
	"strings"
)

// TODO optimize/refactor code
func AlexaIntentHandler(c echo.Context) error {
	var builder ala.SSMLBuilder
	u := new(alexa.Request)
	if err := c.Bind(u); err != nil {
		builder.Say("Sorry error processing your request.")
		builder.Pause("100")
		builder.Say("Please try again")
		return c.JSON(http.StatusOK, builder.Build())
	}

	var res alexa.Response
	var err error
	switch u.Body.Intent.Name {
	case ala.GetMyEventsForTodayIntentType.String():
		c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.Session.User.AccessToken))
		usr, err := GetUserFromToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			builder.Say("Error authenticating, please try again")
			c.JSON(http.StatusOK, ala.NewSSMLResponse("My Events Today", builder.Build()))
			return nil
		}

		if !usr.IsValid() {
			builder.Say("Error authenticating, please login again")
			c.JSON(http.StatusOK, ala.NewSSMLResponse("My Events Today", builder.Build()))
			return nil
		}

		service := auth.NewService()
		token, err := service.GetToken()
		if err != nil {
			builder.Say("Error authenticating, please try again")
			c.JSON(http.StatusOK, ala.NewSSMLResponse("My Events Today", builder.Build()))
			return nil
		}

		c.Set("token", token)

		//user, err := service.GetUser(usr.Sub)
		//if err != nil {
		//	fmt.Println("ERRR " + err.Error())
		//} else {
		//	fmt.Println("USER " + auth0.StringValue(user.Email))
		//}

		//usr.email = user.Email

		//var passwordless bool
		//for _, id := range user.Identities {
		//	if auth0.StringValue(id.Provider) == "email" {
		//		passwordless = true
		//		break
		//	}
		//}

		//if !passwordless {
		//	fmt.Println("Provisioning email passwordless user for " + fmt.Sprintf("%v", usr))
		//	_ = service.CreateUser("email", "temesxgn@gmail.com")
		//}

		res, err = GetMyEventsForTodayIntent(usr)
	default:
		res, err = HandleHelpIntent(u)
	}

	//res, err := IntentDispatcher(u)
	if err != nil {
		return c.JSON(http.StatusOK, ala.NewSSMLResponse("My Events Today", err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

//
//// IntentDispatcher -
//func IntentDispatcher(request *alexa.Request) (alexa.Response, error) {
//	switch request.Body.Intent.Name {
//	case ala.GetMyEventsForTodayIntentType.String():
//		return GetMyEventsForTodayIntent(request)
//	default:
//		return HandleHelpIntent(request)
//	}
//}

func GetUserFromToken(token string) (*auth.User, error) {
	if split := strings.Split(token, " "); len(split) == 2 {
		token := split[1]
		tkn, err := jwt.ParseSigned(token)
		if err != nil {
			return nil, err
		}

		usr := &auth.User{}
		if err := tkn.UnsafeClaimsWithoutVerification(usr); err == nil {
			return usr, nil
		}
	}

	return nil, errors.New("malformed authentication request")
}
