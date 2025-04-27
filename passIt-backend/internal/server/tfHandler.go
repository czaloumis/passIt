package server

import (
	"fmt"
	"log"
	"net/http"
	"passIt/internal/core"

	"github.com/gin-gonic/gin"
)

func (s *Server) TerraformInitHandler(c *gin.Context) {
	configPath := `C:\Users\chrys\projects\demo-eks\terraform-deploy-eks` // Workdir
	execPath := `C:\Users\chrys\Tools\terraform.exe`                      // Replace with your actual exec path
	tf := core.NewTerraform(configPath, execPath)
	// Initialize Terraform
	log.Println("Starting Terraform Init...")
	err := tf.Init()
	if err != nil {
		log.Println("Terraform Init failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Terraform"})
		return
	}
	log.Println("Terraform Init completed successfully")

	log.Println("Getting Terraform version...")
	// Get the Terraform version
	version, err := tf.Show()
	if err != nil {
		log.Println("Failed to get Terraform version:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Terraform version"})
		return
	}
	fmt.Println("Terraform version:", version)

	c.JSON(http.StatusOK, gin.H{"version": version})
}
