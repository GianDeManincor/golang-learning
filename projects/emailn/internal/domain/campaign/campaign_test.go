package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign XYZ"
	content := "Body"
	contacts := []string{"email1@email.com", "email2@email.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Contains(campaign.Content, content)
	assert.Contains(contacts, "email2@email.com")
	assert.Equal(len(contacts), len(campaign.Contacts))
}
