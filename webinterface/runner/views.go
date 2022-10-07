package runner

import (
	"net/http"
	"test_go/common"
	"test_go/pipelines"

	"github.com/gin-gonic/gin"
)

type RunnerView struct {
	ParentGroup *gin.RouterGroup
	Pipeline    *pipelines.DispatchPipeline
}

func (view *RunnerView) RegisterView() {
	router := view.ParentGroup.Group("/runners")

	router.POST("", func(c *gin.Context) {
		var task pipelines.DispatchTask

		common.LoadJSONToModel(c, &task)

		view.Pipeline.DipatchChannel <- task
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

}
