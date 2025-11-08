package app

import (
	"vacanciesParser/internal/core/api/handlers"
	logic_skilltree "vacanciesParser/internal/core/logic/skilltree"
	"vacanciesParser/internal/core/repository"

	"github.com/gin-gonic/gin"
)

func PrepareRouter() *gin.Engine {
	router := gin.Default()

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
