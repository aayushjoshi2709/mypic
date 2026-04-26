package group

import "github.com/gin-gonic/gin"

type Handler struct {
	repos map[string]any
}

func (h *Handler) New(repos map[string]any) {
	h.repos = repos
}

func (h *Handler) add(ctx *gin.Context){

}

func (h *Handler) get(ctx *gin.Context){

}

func (h *Handler) getAll(ctx *gin.Context){

}

func (h *Handler) delete(ctx *gin.Context){

}

func (h *Handler) addImage(ctx *gin.Context){

}

func (h *Handler) addUser(ctx *gin.Context){

}
