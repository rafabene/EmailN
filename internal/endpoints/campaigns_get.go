package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)

	campaigns := h.CampaignService.Repository.Get()

	render.JSON(w, r, campaigns)
}
