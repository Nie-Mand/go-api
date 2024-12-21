package domain

type Hello struct {
	World string `json:"world"`
}

func (h *Hello) IsNil() bool {
	return h.World == ""
}
