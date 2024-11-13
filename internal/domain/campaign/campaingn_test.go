package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func TestNewCampaignCreated(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func TestNewCampaign_IdNotNil(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func TestNewCampaign_DateIsnow(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	now := time.Now().Add(-time.Minute)

	assert.Greater(campaign.CreateOn, now)

}

func TestNewCampaign_MustValidateName(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign("", content, contacts)

	assert.Equal("name is required", error.Error())

}

func TestNewCampaign_MustValidateContent(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign(name, "", contacts)

	assert.Equal("content is required", error.Error())

}
