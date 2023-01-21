package route

import (
	"go-clean-architecture/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	NewMetricRoute(env, timeout, db, publicRouterV1)
}
