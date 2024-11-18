package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaignDTO contract.NewCampaignDTO) (string, error) {

	campaign, err := NewCampaign(newCampaignDTO.Name, newCampaignDTO.Content, newCampaignDTO.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil

}
