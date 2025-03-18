package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/internal/service"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Email    string `json:"email" binding:"required,email"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body RegisterRequest true "Register user"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.userService.Register(c.Request.Context(), req.Username, req.Password, req.Email); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			response.Error(c, http.StatusBadRequest, appErr.Code, appErr.Message)
			return
		}
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body LoginRequest true "Login user"
// @Success 200 {object} response.Response{data=string} "token"
// @Failure 401 {object} response.Response
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	token, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			response.Error(c, http.StatusUnauthorized, appErr.Code, appErr.Message)
			return
		}
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"token": token})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := h.userService.GetUserInfo(c.Request.Context(), uint(userID))
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			response.Error(c, http.StatusNotFound, appErr.Code, appErr.Message)
			return
		}
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, user)
}