package services

import (
	"testing"

	"github.com/jasonnchann24/gogin-rest/entities"
	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

func getVideo() entities.Video {
	return entities.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	service := New()

	service.Save(getVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
