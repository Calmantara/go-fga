package v1

import (
	engine "github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/pkg/domain/user"
	"github.com/Calmantara/go-fga/pkg/server/http/middleware"
	"github.com/Calmantara/go-fga/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {

	// setiap yang /v1/user
	// harus melakukan pengecheckan auth
	// sehingga kita bisa meletakkan middleware di dalam group kita

	routerGroup := ginEngine.GetGin().Group("/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	// all path for get method are here
}

func (u *UserRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("",
		middleware.CheckJwtAuth, u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
}
