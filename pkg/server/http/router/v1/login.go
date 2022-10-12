package v1

import (
	"net/http"
	"time"

	engine "github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/pkg/domain/claim"
	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/server/http/router"
	"github.com/Calmantara/go-fga/pkg/usecase/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
}

func NewLoginRouter(ginEngine engine.HttpServer) router.Router {

	routerGroup := ginEngine.GetGin().Group("/v1/login")
	return &LoginImpl{ginEngine: ginEngine, routerGroup: routerGroup}
}

func (u *LoginImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("", func(ctx *gin.Context) {
		// check body
		body := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		// binding body
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "failed to bind payload",
			})
			return
		}
		// check username dan password
		if body.Password != "PASSWORD" || body.Username != "CALMAN" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "password or username is not matched",
			})
			return
		}

		// jika username dan password valid, akan mengeluarkan access token
		timeNow := time.Now()
		claimAccess := claim.Access{
			JWTID:          uuid.New(),
			Subject:        "CALMAN",
			Issuer:         "go-fga.com",
			Audience:       "user.go-fga.com",
			Scope:          "user",
			Type:           "ACCESS_TOKEN",
			IssuedAt:       timeNow.Unix(),
			NotValidBefore: timeNow.Unix(),
			ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
		}
		accessToken := crypto.CreateJWT(ctx, claimAccess)

		claimRefresh := claim.Access{
			JWTID:          uuid.New(),
			Subject:        "CALMAN",
			Issuer:         "go-fga.com",
			Audience:       "user.go-fga.com",
			Scope:          "user",
			Type:           "REFRESH_TOKEN",
			IssuedAt:       timeNow.Unix(),
			NotValidBefore: timeNow.Unix(),
			ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
		}
		refreshToken := crypto.CreateJWT(ctx, claimRefresh)

		// access token ini akan digunakan selanjutnya
		ctx.JSON(http.StatusOK, message.Response{
			Code:    0,
			Message: "successfully login",
			Data: map[string]any{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		})
	})
}

func (u *LoginImpl) Routers() {
	u.post()
}
