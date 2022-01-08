package workers

import (
	"appbar/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetWorkersFromBar(c *gin.Context) {
	id := strings.ReplaceAll(c.Param("id"), "/", "")
	data := FindAllWorkersFromBar(id)
	c.JSON(http.StatusOK, gin.H{"workers": data})
}

func WorkerAssignation(c *gin.Context) {
	workerValidator := NewWorkerValidator()

	if err := workerValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	if err := AssignWorkerToBar(&workerValidator.workerModel); err != nil {
		roleError := common.NewError("worker", err)
		c.JSON(http.StatusUnprocessableEntity, roleError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Worker asignation OK"})
}

func WorkerRemoval(c *gin.Context) {
	workerValidator := NewWorkerValidator()

	if err := workerValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	if err := DeleteWorkerFromBar(&workerValidator.workerModel); err != nil {
		roleError := common.NewError("worker", err)
		c.JSON(http.StatusUnprocessableEntity, roleError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission removal OK"})
}
