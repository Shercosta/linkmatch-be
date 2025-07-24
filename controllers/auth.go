package controllers

import (
	"linkmatch-be/requests"
	"linkmatch-be/responses"
	"linkmatch-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Register a new user
// @Description Handles user registration
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "Register body"
// @Router /auth/register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body requests.RegisterRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		result := services.Register(
			db,
			c,
			&body,
		)

		if result != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, result.Error(), nil)
			return
		}

		responses.JSONSuccess(c.Writer, "Success", nil, nil)
	}
}

// @Summary Login a user
// @Description Handles user login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "Login body"
// @Router /auth/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body requests.RegisterRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		var user = services.GetUser(db, body.Username)
		if user == nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, "User not found", nil)
			return
		}

		result, err := services.Login(
			db,
			&body,
			user,
		)

		if err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		responses.JSONSuccess(c.Writer, result, nil, nil)
	}
}
