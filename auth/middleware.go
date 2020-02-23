package auth

// // HasAdminSecret - Check if context has admin secret
// func HasAdminSecret(ctx context.Context) bool {
// 	if secret := ctx.Value(AdminSecretCtxKey); secret != nil && secret == config.GetHasuraSecret() {
// 		return true
// 	}

// 	return false
// }

// // HasAuthorization - Check if context has api key & secret
// func HasAuthorization(ctx context.Context) bool {
// 	if key := ctx.Value(APIKey); key != nil && key == config.GetAPIKey() {
// 		if secret := ctx.Value(APISecret); secret != nil && secret == config.GetAPISecret() {
// 			return true
// 		}
// 	}

// 	return false
// }

// // GetUserFromToken - convert JWT token to user object
// func GetUserFromToken(token string) *User {
// 	if split := strings.Split(token, " "); len(split) == 2 {
// 		token := split[1]
// 		tkn, _ := jwt.ParseSigned(token)

// 		usr := &User{}
// 		if err := tkn.UnsafeClaimsWithoutVerification(usr); err == nil {
// 			return usr
// 		}
// 	}

// 	return &User{}
// }

// // Middleware - Auth middleware integration
// func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := SetValuesFromHeaders(c.Request())
// 		if HasAdminSecret(ctx) || HasAuthorization(ctx) {
// 			return next(c)
// 		}

// 		return echo.NewHTTPError(http.StatusUnauthorized, "Access Denied")
// 	}
// }
