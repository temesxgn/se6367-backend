package handlers

import (
	"fmt"
	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	"github.com/square/go-jose/jwt"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth"
	"net/http"
	"strings"
)

func AlexaIntentHandler(c echo.Context) error {
	var builder ala.SSMLBuilder
	u := new(alexa.Request)
	if err := c.Bind(u); err != nil {
		builder.Say("Sorry error processing your request.")
		builder.Pause("100")
		builder.Say("Please try again")
		return c.JSON(http.StatusBadRequest, builder.Build())
	}

	var res alexa.Response
	var err error
	switch u.Body.Intent.Name {
	case ala.GetMyEventsForTodayIntentType.String():
		c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.Session.User.AccessToken))

		//jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		//		return []byte("WEW0t4UqoC1-vaeSCrcyyPOUdRXdH792r-Xl7F2aZuQG1zu9nFv8vdtPVfsGmN95"), nil
		//	},
		//	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		//	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		//	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//	SigningMethod: jwt.SigningMethodRS256,
		//})
		//
		//if err := jwtMiddleware.CheckJWT(c.Response().Writer, c.Request()); err != nil {
		//	builder.Say("Failed authenticating your account")
		//	return c.JSON(http.StatusBadRequest, builder.Build())
		//}

		//tkn, err := jwt.Parse(u.Session.User.AccessToken, func(token *jwt.Token) (interface{}, error) {
		//	// Don't forget to validate the alg is what you expect:
		//	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//	}
		//
		//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		//	return []byte("WEW0t4UqoC1-vaeSCrcyyPOUdRXdH792r-Xl7F2aZuQG1zu9nFv8vdtPVfsGmN95"), nil
		//})

		usr := GetUserFromToken(u.Session.User.AccessToken)

		//c.Set("user", tkn)

		service := auth.NewService()
		token, err := service.GetToken()
		if err != nil {
			builder.Say(err.Error())
			c.JSON(http.StatusInternalServerError, builder.Build())
		}

		c.Set("token", token)

		//user, _ := service.GetUser(usr.Sub)
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
		return c.JSON(http.StatusInternalServerError, err.Error())
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

func GetUserFromToken(token string) *auth.User {
	if split := strings.Split(token, " "); len(split) == 2 {
		token := split[1]
		tkn, _ := jwt.ParseSigned(token)

		usr := &auth.User{}
		if err := tkn.UnsafeClaimsWithoutVerification(usr); err == nil {
			return usr
		}
	}

	return &auth.User{}
}
