package controllers

import (
	"encoding/base64"

	"github.com/goadesign/goa"
	"github.com/m0a/goagooglelogin"
	"github.com/m0a/goagooglelogin/examples/simple/app"
)

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
	Accounts *map[string]Account
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service, ac *map[string]Account) *JWTController {
	return &JWTController{
		Controller: service.NewController("JWTController"),
		Accounts:   ac,
	}
}

// Secure runs the secure action.
func (c *JWTController) Secure(ctx *app.SecureJWTContext) error {
	// jwtContext := jwt.ContextJWT(ctx)
	// if jwtContext == nil {
	// 	return errors.New("jwtContext is nil")
	// }
	// claims, ok := jwtContext.Claims.(jwtgo.MapClaims)
	// if !ok {
	// 	return ctx.Unauthorized()
	// }
	// googleID, ok := claims["sub"].(string)
	// if !ok {
	// 	return ctx.Unauthorized()
	// }
	googleID, err := goagooglelogin.GoogleIDByJWTContext(ctx)

	if err != nil {
		return ctx.Unauthorized()
	}

	if c.Accounts == nil {
		return ctx.Unauthorized()
	}
	account, ok := (*c.Accounts)[googleID]
	if !ok {
		return ctx.Unauthorized()
	}

	img := base64.StdEncoding.EncodeToString(account.Image)
	res := app.GoaExamplesSecuritySecure{
		Name:  &account.Name,
		Email: &account.Email,
		Image: &img,
	}
	return ctx.OK(&res)
}

// Unsecure runs the unsecure action.
func (c *JWTController) Unsecure(ctx *app.UnsecureJWTContext) error {
	// JWTController_Unsecure: start_implement
	// JWTController_Unsecure: end_implement
	res := &app.Success{}
	return ctx.OK(res)
}
