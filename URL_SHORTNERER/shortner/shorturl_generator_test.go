package shortner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "41e5f550-872c-44a6-bc12-9c9c6a67aa39"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "http://google.com"
	initialLink_2 := "https://www.youtube.com/watch?v=Y7rkTu8mWmE"
	initialLink_3 := "https://github.com"

	shortLink1 := GenerateShortUrl(initialLink_1, UserId)
	shortLink2 := GenerateShortUrl(initialLink_2, UserId)
	shortLink3 := GenerateShortUrl(initialLink_3, UserId)

	// Test URL length
	assert.Equal(t, 8, len(shortLink1))
	assert.Equal(t, 8, len(shortLink2))
	assert.Equal(t, 8, len(shortLink3))

	// Test different inputs produce different outputs
	assert.NotEqual(t, shortLink1, shortLink2)
	assert.NotEqual(t, shortLink2, shortLink3)
	assert.NotEqual(t, shortLink1, shortLink3)

	// Test consistency
	assert.Equal(t, shortLink1, GenerateShortUrl(initialLink_1, UserId))
	assert.Equal(t, shortLink2, GenerateShortUrl(initialLink_2, UserId))
	assert.Equal(t, shortLink3, GenerateShortUrl(initialLink_3, UserId))
}
