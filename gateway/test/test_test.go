package test

import (
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	id := xid.New().String()
	fr, err := xid.FromString(id)
	assert.NoError(t, err)
	assert.Equal(t, fr.String(), id)
	_, err = xid.FromString("123")
	assert.Error(t, err)
}
