package controller

import "go-clean-architecture/interface/controller"

type AppController struct {
	Metric interface{ controller.MetricController }
}
