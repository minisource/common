package helper

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			GenerateBaseResponseWithValidationError(nil, false, ValidationError, err))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(TranslateErrorToStatusCode(err),
			GenerateBaseResponseWithError(nil, false, InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, GenerateBaseResponse(res, true, 0))
}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			GenerateBaseResponseWithValidationError(nil, false, ValidationError, err))
		return
	}

	res, err := caller(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(TranslateErrorToStatusCode(err),
			GenerateBaseResponseWithError(nil, false, InternalError, err))
		return
	}
	c.JSON(http.StatusOK, GenerateBaseResponse(res, true, 0))
}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			GenerateBaseResponse(nil, false, ValidationError))
		return
	}

	err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(TranslateErrorToStatusCode(err),
			GenerateBaseResponseWithError(nil, false, InternalError, err))
		return
	}
	c.JSON(http.StatusOK, GenerateBaseResponse(nil, true, 0))
}

func GetById[To any](c *gin.Context, caller func(c context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			GenerateBaseResponse(nil, false, ValidationError))
		return
	}

	res, err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(TranslateErrorToStatusCode(err),
			GenerateBaseResponseWithError(nil, false, InternalError, err))
		return
	}
	c.JSON(http.StatusOK, GenerateBaseResponse(res, true, 0))
}

func GetByFilter[Ti any, To any](c *gin.Context, caller func(c context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			GenerateBaseResponseWithValidationError(nil, false, ValidationError, err))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(TranslateErrorToStatusCode(err),
			GenerateBaseResponseWithError(nil, false, InternalError, err))
		return
	}
	c.JSON(http.StatusOK, GenerateBaseResponse(res, true, 0))
}
