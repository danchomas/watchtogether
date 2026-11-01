package transporthttp

import "net/http"

func (h *Handler) message(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Text string `json:"text"`
	}
	if !decode(w, r, &body) {
		return
	}
	if err := h.rooms.SendMessage(r.PathValue("code"), token(r), body.Text); err != nil {
		fail(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
