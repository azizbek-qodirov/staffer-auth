package api

import (
	"log"

	"github.com/Azizbek-Qodirov/hr_platform_auth_service/api/handler"
	"github.com/Azizbek-Qodirov/hr_platform_auth_service/api/middleware"
	_ "github.com/Azizbek-Qodirov/hr_platform_auth_service/docs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	r := gin.Default()


	r.Use(middleware.NewAuth(ca))
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	u := r.Group("/user")
	u.POST("/register", h.Register)
	u.GET("/refreshTokenRequest", h.RefreshToken)
	u.POST("/login", h.Login)
	u.DELETE("/logout", h.Logout)
	u.PUT("admin/update", h.UpdateForAdmin)
	u.PUT("/update", h.Update)

	return r
}
