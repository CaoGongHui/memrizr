package handler

import (
	"net/http"

	"github.com/caogonghui/memrizr/account/model"
	"github.com/gin-gonic/gin"
)

// 主要持有所有服务的属性 包括一些业务核心代码的实现方法
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// 配置需要在handler包中初始化的值
type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
	BaseURL      string
}

// 一个工厂方法 在main函数中调用初始化路由route
func NewHandler(c *Config) {
	// Create an account group
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	} // currently has no properties

	// Create a group, or base url for all routes
	// 使用组的方式 ACCOUNT_API_URL通过配置文件传递
	g := c.R.Group(c.BaseURL)

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
