package services

import (
	"encoding/json"
	"fmt"
	"linkmatch-be/database/models"
	"linkmatch-be/requests"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunNodeParser(cvPath string) (err error) {
	// Full command: node prisma/app.js cvPath
	cmd := exec.Command("node", "prisma/app.js", cvPath)

	// Optional: capture output for logging/debugging
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Failed to run parser: %v\n", err)
		return err
	}

	// Print output (stdout and stderr combined)
	fmt.Println(string(output))
	return nil
}

func GetJsonData(uuidName string) gin.H {
	// Construct the absolute path to the JSON file
	path := filepath.Join("prisma", "parsed-cv", uuidName+".json")

	// Read the JSON file
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("❌ Error reading JSON file:", err)
		return gin.H{"error": "Failed to read JSON file"}
	}

	// Parse JSON into a generic map
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("❌ Error parsing JSON:", err)
		return gin.H{"error": "Invalid JSON format"}
	}

	// Return the parsed data
	return jsonData
}

func UpdateProfile(db *gorm.DB, username string, body *requests.ProfileRequest) models.UserPublic {
	var user models.UserPublic
	result := db.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return models.UserPublic{}
	}

	user.Name = body.Name
	user.ProfessionalTitle = body.ProfessionalTitle
	user.CVJson = body.CVJson
	user.CompanyName = body.CompanyName
	user.Location = body.Location
	user.Description = body.Description

	db.Save(&user)
	return user
}
