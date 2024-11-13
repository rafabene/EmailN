package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body Hi!"
	contacts = []string{"email1@e.com", "email2@e.com"}
	fake     = faker.New()
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

func TestNewCampaign_MustValidateNameMin(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", error.Error())

}

func TestNewCampaign_MustValidateNameMax(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign(fake.Lorem().Text(25), content, contacts)

	assert.Equal("name is required with max 24", error.Error())

}

func TestNewCampaign_MustValidateContentMin(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", error.Error())

}

func TestNewCampaign_MustValidateContentMax(t *testing.T) {

	assert := assert.New(t)

	_, error := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", error.Error())

}

func Test_NewCampaign_CreatedOnMustBenNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreateOn, now)

}

func Test_NewCampaign_MustValidatContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", error.Error())
}

func Test_NewCampaign_MustValidatContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", error.Error())
}
