package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var teststoreservice *StoreService

func init() {
	teststoreservice = InitstoreService()
}
func TestInsertation(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shorturl := "Jsz4k57oAX"

	SaveUrlMapping(shorturl, initialLink, userUUId)

	retrieveUrl := RetrieveInitialUrl(shorturl)

	assert.Equal(t, initialLink, retrieveUrl)
}
