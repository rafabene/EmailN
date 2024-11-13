package campaign

import (
	internalerrors "emailn/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID       string    `validate:"required"`
	Name     string    `validate:"min=5,max=24"`
	CreateOn time.Time `validate:"required"`
	Content  string    `validate:"min=5,max=1024"`
	Contacts []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, val := range emails {
		contacts[index].Email = val
	}
	campaign := &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreateOn: time.Now(),
		Contacts: contacts,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return campaign, err
}
