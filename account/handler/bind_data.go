package handler

import (
	"log"

	"github.com/caogonghui/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// used to help ex
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error bidding data: %+v\n", err)
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
