package webinterface

import (
	"test_go/common"
	"test_go/pipelines"
	"test_go/webinterface/runner"
	"test_go/webinterface/users"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := common.GetDB()

	router := gin.Default()
	api := router.Group("/api/v1")

	pipeline := &pipelines.DispatchPipeline{}

	views := &[...]common.WebView{
		&users.UsersView{ParentGroup: api, DB: db},
		&runner.RunnerView{ParentGroup: api, Pipeline: pipeline},
	}

	for _, v := range *views {
		v.RegisterView()
	}
	pipeline.Run()

	router.Run(":3030")
}
