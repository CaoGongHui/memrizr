package handler

import (
	"log"
	"net/http"

	"github.com/caogonghui/memrizr/account/model"
	"github.com/caogonghui/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" bidding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	var req signupReq
	if ok := bindData(c, &req); !ok {
		return
	}
	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.UserService.Signup(c, u)
	if err != nil {
		log.Printf("Faild to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		// may eventually implement rollback logic here
		// meaning, if we fail to create tokens after creating a user,
		// we make sure to clear/delete the created user in the database

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
