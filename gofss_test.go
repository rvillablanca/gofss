package gofss

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// This test requires you run the flying-saucer-service. Check target create-service in Makefile.
func TestFlayingSaucerService(t *testing.T) {
	c := New("http://localhost:8082/convert")
	require.NotNil(t, c)
	pdf, err := c.GeneratePDF("")
	assert.Error(t, err)
	assert.Nil(t, pdf)

	pdf, err = c.GeneratePDF("badhtm")
	assert.Error(t, err)
	assert.Nil(t, pdf)

	pdf, err = c.GeneratePDF("<p>hello world</p>")
	assert.NoError(t, err)
	assert.NotNil(t, pdf)
}
