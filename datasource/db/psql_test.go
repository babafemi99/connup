package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	client := GetClient()
	assert.NotNil(t, client)
}
