package v1

import (
	"fmt"
	"github.com/Hellsker/Todo-List/internal/entity"
	"github.com/Hellsker/Todo-List/internal/logger"
	"github.com/Hellsker/Todo-List/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TaskRoutes struct {
	log  logger.Interface
	serv service.TaskInterface
}

// NewTaskRoutes constructor
func NewTaskRoutes(handler *gin.RouterGroup, logger logger.Interface, service service.TaskInterface) *TaskRoutes {
	r := TaskRoutes{logger, service}
	h := handler.Group("/todo")
	{
		h.GET("/task/:id", r.GetById)
		h.GET("/task", r.GetByDateWithFilter)
		h.GET("/tasks", r.GetAllWithPagination)
		h.POST("/task", r.Post)
		h.PUT("/task/:id", r.Put)
		h.DELETE("/task/:id", r.Delete)

	}
	return &r
}

// @Summary     Get Task by ID
// @Description Get Task by ID
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param 		id	 path	 int	true	"Task ID"
// @Success     200 {object} entity.TaskResponse
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/task/{id} [get]
func (r *TaskRoutes) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Error("TaskRoutes - GetById - strconv.ParseUint:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid task ID!")
		return
	}
	taskRes, err := r.serv.GetById(c.Request.Context(), id)
	if err != nil {
		r.log.Info(err.Error())
		errorResponse(c, http.StatusNotFound, "Task with this ID does not exist!")
		return
	}
	c.JSON(http.StatusOK, taskRes)
}

// @Summary     Get Task by Date With Filter
// @Description Get Task by Date With Filter
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param       date   	  query     string false	"on what date to issue the tasks"
// @Param       status    query     bool  false  "filter parameter completed"
// @Success     200 {object} []entity.TaskResponse
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/task [get]
func (r *TaskRoutes) GetByDateWithFilter(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format(time.DateOnly))
	statusFilter := c.DefaultQuery("status", "")
	dateLimit, err := time.Parse(time.DateOnly, date)
	if err != nil {
		r.log.Error("TaskRoutes - GetByDateWithFilter- time.Parse:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid date parameter!")
		return
	}
	if statusFilter == "true" || statusFilter == "false" || statusFilter == "" {
		tasks, err := r.serv.GetByDateWithFilter(c.Request.Context(), dateLimit, statusFilter)
		if err != nil {
			r.log.Error(err.Error())
			errorResponse(c, http.StatusInternalServerError, "Database error!")
			return
		}
		fmt.Println(tasks)
		c.JSON(http.StatusOK, tasks)
		return
	}
	errorResponse(c, http.StatusBadRequest, "Invalid status parameter!")
	return

}

// @Summary     Get Tasks With Pagination
// @Description Get Tasks With Pagination
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param        page    query     int  false  "pagination parameter page"
// @Param        pageSize    query     int  false  "pagination parameter page size"
// @Param        status    query     bool  false  "filter parameter completed"
// @Success     200 {object} []entity.TaskResponse
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/tasks [get]
func (r *TaskRoutes) GetAllWithPagination(c *gin.Context) {
	// Pagination param
	page, pageSize := c.DefaultQuery("page", "1"), c.DefaultQuery("pageSize", "10")
	r.log.Warn(page)
	// Status filter param
	statusFilter := c.DefaultQuery("status", "")
	pageNum, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		r.log.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, "Invalid pagination parameter!")
		return
	}
	limit, err := strconv.ParseUint(pageSize, 10, 64)
	if err != nil {
		r.log.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, "Invalid pagination parameter!")
		return
	}
	offset := (pageNum - 1) * limit
	if statusFilter == "true" || statusFilter == "false" || statusFilter == "" {
		tasks, err := r.serv.GetAllWithPagination(c.Request.Context(), limit, offset, statusFilter)
		if err != nil {
			r.log.Error(err.Error())
			errorResponse(c, http.StatusInternalServerError, "Database error!")
			return
		}
		c.JSON(http.StatusOK, tasks)
		return
	}
	errorResponse(c, http.StatusBadRequest, "Invalid status parameter!")
	return

}

// @Summary     Create Task
// @Description Create Task
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param 		request body entity.TaskRequest true "body params"
// @Success     200 {object} string
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/task [post]
func (r *TaskRoutes) Post(c *gin.Context) {
	var task entity.TaskRequest
	err := c.ShouldBindJSON(&task)
	if err != nil {
		r.log.Error("TaskRoutes - Post - c.ShouldBindJSON:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid body!")
		return
	}
	err = r.serv.Save(c.Request.Context(), task)
	if err != nil {
		r.log.Error(err.Error())
		errorResponse(c, http.StatusInternalServerError, "Database error!")
		return
	}
	c.JSON(http.StatusOK, "Successfully created!")
	return
}

// @Summary     Update Task
// @Description Update Task
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param		id	 path	 int	true	"Task ID"
// @Param 		request body entity.TaskRequest true "body params"
// @Success     200 {object} message
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/task/{id} [put]
func (r *TaskRoutes) Put(c *gin.Context) {
	var task entity.TaskRequest
	err := c.ShouldBindJSON(&task)
	if err != nil {
		r.log.Error("TaskRoutes - Put - c.ShouldBindJSON:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid body!")
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Error("TaskRoutes - Put - strconv.Atoi:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid task ID!")
		return
	}
	err = r.serv.Update(c.Request.Context(), id, task)
	if err != nil {
		r.log.Error(err.Error())
		errorResponse(c, http.StatusNotFound, "Task with this ID does not exist!")
		return
	}
	messageResponse(c, http.StatusOK, "Successfully updated!")
	return
}

// @Summary     Delete Task
// @Description Delete Task
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param		id	 path	 int	true	"Task ID"
// @Success     200 {object} string
// @Failure     400  {object}  error
// @Failure     404  {object}  error
// @Failure     500  {object}  error
// @Router      /todo/task/{id} [delete]
func (r *TaskRoutes) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		r.log.Error("TaskRoutes - Delete - strconv.ParseUint:", err)
		errorResponse(c, http.StatusBadRequest, "Invalid task ID!")
		return
	}
	err = r.serv.Delete(c.Request.Context(), id)
	if err != nil {
		r.log.Info(err.Error())
		errorResponse(c, http.StatusNotFound, "Task with this ID does not exist!")
		return
	}
	messageResponse(c, http.StatusOK, "Successfully deleted!")
}
