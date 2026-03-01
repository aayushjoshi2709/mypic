package comment

import "github.com/gin-gonic/gin"

type Handler struct{
	repo *Repository
}

func (h *Handler) New(repo *Repository){
	h.repo = repo
}

func (h *Handler) get(ctx *gin.Context) {}

func (h *Handler) getAll(ctx *gin.Context) {}

func (h *Handler) create(ctx *gin.Context) {}

func (h *Handler) update(ctx *gin.Context) {}

func (h *Handler) delete(ctx *gin.Context) {}
