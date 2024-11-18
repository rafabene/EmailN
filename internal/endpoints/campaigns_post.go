package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) PostCampaigns(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err

}
