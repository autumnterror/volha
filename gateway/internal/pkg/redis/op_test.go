package redis

import (
	"gateway/config"
	"gateway/internal/views"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDictionariesGood(t *testing.T) {
	c := New(config.Test())

	td := &views.Dictionaries{
		Brands:     []views.Brand{},
		Categories: []views.Category{},
		Countries:  []views.Country{},
		Materials:  []views.Material{},
		Colors:     []views.Color{},
		MinPrice:   1,
		MaxPrice:   2,
		MinWidth:   3,
		MaxWidth:   4,
		MinHeight:  5,
		MaxHeight:  6,
		MinDepth:   7,
		MaxDepth:   8,
	}

	assert.NoError(t, c.SetDictionaries(td))

	d, err := c.GetDictionaries()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	log.Println(d)

	assert.NoError(t, c.CleanDictionaries())

	d, err = c.GetDictionaries()
	assert.Error(t, err)
	assert.Nil(t, d)
}

func TestDictionariesNoCache(t *testing.T) {
	c := New(config.Test())

	assert.NoError(t, c.CleanDictionaries())

	d, err := c.GetDictionaries()
	assert.Error(t, err)
	assert.Nil(t, d)
}
