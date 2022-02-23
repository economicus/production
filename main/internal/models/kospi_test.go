package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKospi(t *testing.T) {
	a := assert.New(t)
	k := NewKospi()
	a.NotEmpty(k, "kospi must not be empty")
}
