package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestCreateCampaingn(t *testing.T) {

	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test@test.com"},
	}

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)

}

func TestSaveCampaingn(t *testing.T) {

	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test@test.com"},
	}

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)

}
