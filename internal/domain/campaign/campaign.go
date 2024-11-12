package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))

	for i, value := range emails {
		contacts[i].Email = value
	}
	return &Campaign{
		ID:        "1",
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
