package routes

import (
	"log"

	"github.com/WantToBePro31/rand-pass/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("views/*.html")
	r.GET("/", controllers.RedirectToRandPassPage)
	r.GET("/random-password-generator", controllers.PasswordGeneratorPage)
	r.POST("/random-password-generator", controllers.SubmitRequestHandler)
	r.GET("/random-password-generator/result", controllers.PasswordResultPage)
	r.POST("/random-password-generator/result", controllers.DownloadPassword)
	r.GET("/random-password-generator/result/downloaded", controllers.PasswordDownloadedPage)
	log.Println("Routes Set")
}
