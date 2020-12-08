package climacell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("c0ffee", nil)
	assert.NoError(t, err)
	assert.True(t, c.httpClient != nil)
}

func TestNewClient_emptyBaseURL(t *testing.T) {
	backupBaseURL := BaseURL
	BaseURL = ""

	defer func() {
		// Reset BaseURL to prevent side effects
		BaseURL = backupBaseURL
	}()

	c, err := NewClient("c0ffee", nil)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidBaseURL)
	assert.Nil(t, c)
}

func TestNewClient_invalidAPIKey(t *testing.T) {
	c, err := NewClient("", nil)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidAPIKey)
	assert.Nil(t, c)
}
