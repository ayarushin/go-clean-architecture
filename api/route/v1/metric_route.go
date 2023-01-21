package route

import (
	"go-clean-architecture/api/controller"
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/domain"
	"go-clean-architecture/repository"
	"go-clean-architecture/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewMetricRoute(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	mr := repository.NewMetricRepository(db, domain.MetricTable)
	mc := &controller.MetricController{
		MetricUsecase: usecases.NewMetricUsecase(mr, timeout),
		Env:           env,
	}
	group.POST("/metric", mc.Create)
	group.GET("/metric", mc.Fetch)
	group.GET("/metric/:id", mc.GetByID)
}
