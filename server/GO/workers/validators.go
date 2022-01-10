package workers

import (
	"appbar/common"
	"github.com/gin-gonic/gin"
)

type WorkerValidator struct {
	Worker struct {
		User string `form:"user" json:"user" binding:"required"`
		Bar  string `form:"bar" json:"bar" binding:"required"`
		Role string `form:"role" json:"role" binding:"required"`
	} `json:"worker"`
	workerModel Worker `json:"-"`
}

func (worker *WorkerValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, worker)

	if err != nil {
		return err
	}
	worker.workerModel.User = worker.Worker.User
	worker.workerModel.Bar = worker.Worker.Bar
	worker.workerModel.Role = worker.Worker.Role

	return nil
}

func NewWorkerValidator() WorkerValidator {
	return WorkerValidator{}
}
