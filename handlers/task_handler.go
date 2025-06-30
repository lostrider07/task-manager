package handlers

import (
	"net/http"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

var tasks = map[string]models.Task{} // In-memory store

// GetTasks godoc
// @Summary List tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	t := []models.Task{}
	for _, v := range tasks {
		t = append(t, v)
	}
	c.JSON(http.StatusOK, t)
}

// CreateTask godoc
// @Summary Create task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task to create"
// @Success 201 {object} models.Task
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks[task.ID] = task
	c.JSON(http.StatusCreated, task)
}

// GetTask godoc
// @Summary Get a task
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 404 {object} models.ErrorResponse
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	id := c.Param("id")
	if task, exists := tasks[id]; exists {
		c.JSON(http.StatusOK, task)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}
