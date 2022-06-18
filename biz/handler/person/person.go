package person

import (
	"strings"

	"net/http"

	"github.com/LynnWonder/gin_prac/biz/handler"
	"github.com/LynnWonder/gin_prac/biz/service/person"
	"github.com/gin-gonic/gin"
)

type ListAllRequest struct {
	Query    string `form:"query" binding:"max=50"`
	Page     int    `form:"page" binding:"gte=1"`
	PageSize int    `form:"pageSize"`
}

// 绑定 JSON
type CreatePersonRequest struct {
	Name       string `json:"name" binding:"required,lte=50"`
	// description 不是必填参数
	Description string `json:"description" binding:"lte=200"`
}

type UpdatePersonRequest struct {
	PersonId    int `uri:"personId" binding:"required"`
	Name        string `json:"name" binding:"lte=50"`
	Description string `json:"description" binding:"lte=200"`
}

type DeletePersonRequest struct {
	PersonId int `uri:"personId" binding:"required"`
}

type ListQueryRequest struct {
	Page     int `form:"page" binding:"gte=1"`
	PageSize int `form:"pageSize"`
}

// handler 函数传入的是 gin.Context 指针
func ListAll(c *gin.Context) {
	var q ListAllRequest
	// 验证请求参数
	if err := handler.QueryValidate(&q, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}
	Query := q.Query

	var acc = make(map[string]interface{})
	for _, raw := range strings.Split(Query, ",") {
		if strings.Contains(raw, ":") {
			pair := strings.Split(raw, ":")
			k, v := pair[0], pair[1]
			acc[k] = v
		}
	}
	persons, count, err := person.ListPersons(q.Page, q.PageSize, acc)

	if err != nil {
		handler.ErrorsHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"total": count, "data": persons})
}

func CreatePerson(c *gin.Context) {
	var r CreatePersonRequest

	if err := handler.JSONValidate(&r, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}

	data, err := person.CreatePerson(r.Name, r.Description)

	if err != nil {
		handler.ErrorsHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, data)
}

func UpdatePerson(c *gin.Context) {
	var r UpdatePersonRequest

	if err := handler.ParamsValidate(&r, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}
	if err := handler.JSONValidate(&r, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}

	data, err := person.UpdatePerson(r.PersonId, r.Name, r.Description)

	if err != nil {
		handler.ErrorsHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, data)
}

func DeletePerson(c *gin.Context) {
	var r DeletePersonRequest

	if err := handler.ParamsValidate(&r, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}

	err := person.DeletePerson(r.PersonId)

	if err != nil {
		handler.ErrorsHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"msg": "OK"})
}
