package controllers

import (
	"linkmatch-be/requests"
	"linkmatch-be/responses"
	"linkmatch-be/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary Get Profile
// @Description Returns a JSON with the user that requested it
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security Bearer
// @Router /api/profile [get]
func Profile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")

		user := services.GetUserPublic(db, username)

		responses.JSONSuccess(c.Writer, user, nil, nil)
	}
}

// @Summary Parse CV
// @Description Returns a JSON with the parsed CV
// @Tags Profile
// @Accept multipart/form-data
// @Produce json
// @Param cv formData file true "CV file"
// @Success 200 {object} map[string]interface{}
// @Security Bearer
// @Router /api/profile/parse-cv [post]
func ParseResume() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("cv")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uuidFilenameWithoutPDF := uuid.NewString()
		uuidFilename := uuidFilenameWithoutPDF + ".pdf"
		saveDir := "./prisma/cv/"
		savePath := saveDir + uuidFilename

		// create directory
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// save file
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := services.RunNodeParser(savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jsonData := services.GetJsonData(uuidFilenameWithoutPDF)

		// delete file
		if err := os.Remove(savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := os.Remove("./prisma/parsed-cv/" + uuidFilenameWithoutPDF + ".json"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, jsonData)
	}
}

// @Summary Update Profile
// @Description Updates the profile with the given body
// @Tags Profile
// @Accept json
// @Produce json
// @Param request body requests.ProfileRequest true "Profile body"
// @Success 200 {object} models.UserPublic
// @Security Bearer
// @Router /api/profile [put]
func UpdateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body requests.ProfileRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		username := c.GetString("username")

		result := services.UpdateProfile(db, username, &body)

		responses.JSONSuccess(c.Writer, result, nil, nil)
	}
}
