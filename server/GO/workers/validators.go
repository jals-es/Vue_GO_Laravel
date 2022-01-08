package workers

import (
	"appbar/bars"
	"appbar/common"
	"appbar/role_permissions"
	"appbar/users"
	"github.com/gin-gonic/gin"
)

type WorkerValidator struct {
	Worker struct {
		User users.UserModel       `form:"bar" json:"user" binding:"required"`
		Bar  bars.BarModel         `form:"bar" json:"bar" binding:"required"`
		Role role_permissions.Role `form:"role" json:"role" binding:"required"`
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
