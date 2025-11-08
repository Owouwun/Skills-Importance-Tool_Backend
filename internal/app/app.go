package app

import (
	"time"
	"vacanciesParser/internal/core/api/handlers"
	logic_skilltree "vacanciesParser/internal/core/logic/skilltree"
	"vacanciesParser/internal/core/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PrepareRouter() *gin.Engine {
	router := gin.Default()

	router.Use(newCORS())

	repo := repository.NewRepository()
	skilltreeService := logic_skilltree.NewService(repo)
	handlers_skilltree := handlers.NewSkillTreeHandler(skilltreeService)

	apiSkills := router.Group("/api/skills")
	{
		apiSkills.GET("", handlers_skilltree.GetSkillTree)
		apiSkills.POST("", handlers_skilltree.CreateNode)
	}

	return router
}

func newCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://localhost:3000",
	}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Access-Control-Allow-Origin"}
	config.MaxAge = 12 * time.Hour

	return cors.New(config)
}
