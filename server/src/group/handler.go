package group


type Handler struct {
	repo map[string]any
}

func (h *Handler) New(repos map[string]any) {
	h.repo = repos
}