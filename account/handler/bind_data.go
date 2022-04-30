package handler

import (
	"log"

	"github.com/caogonghui/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// used to help extract validation error
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

//为了复用 从signup handler里面提取出来
// return false if binding failed
func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error bidding data: %+v\n", err)
		//如果是标准的validation error 那么就提取出来
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}
			err := apperrors.NewBadRequest("Invalid request paramters,See invalidArgs")
			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}
		// add code for validating max body size here

		// if not validating errors return an internal server error
		fallBack := apperrors.NewInternal()
		c.JSON(fallBack.Status(), gin.H{"error": fallBack})
		return false
	}
	return true
}
