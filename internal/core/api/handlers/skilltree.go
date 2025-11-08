package handlers

import (
	"context"
	"net/http"
	"vacanciesParser/internal/core/logic/skilltree"

	"github.com/gin-gonic/gin"
)

type SkillTreeService interface {
	GetSkillTree(ctx context.Context) (*skilltree.Node, error)
	CreateNode(ctx context.Context, node *skilltree.NodePath) error
}

type SkillTreeHandler struct {
	service SkillTreeService
}

func NewSkillTreeHandler(service SkillTreeService) *SkillTreeHandler {
	return &SkillTreeHandler{
		service: service,
	}
}

func (h *SkillTreeHandler) GetSkillTree(ctx *gin.Context) {
	tree, err := h.service.GetSkillTree(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tree)
}

func (h *SkillTreeHandler) CreateNode(ctx *gin.Context) {
	var newNode *skilltree.NodePath
	if err := ctx.ShouldBindJSON(&newNode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateNode(ctx, newNode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
