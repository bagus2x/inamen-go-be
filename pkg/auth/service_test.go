package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAccessToken(t *testing.T) {
	service := NewService("10i2sqns91w17aafefef")
	claims, err := service.ParseAccessToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJpbmFtZW4udmVyY2VsLmFwcCIsInN1YiI6IjEiLCJpYXQiOjE2MjE2ODM2MTcsImV4cCI6MTYyMTY4NzIxNywidXNlcm5hbWUiOiJiYWd1cyJ9.CyCmOgRR8ikIkCmI6LZ3TKS5xoc0m97qHMM_ebIurys")
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	t.Log(claims)
}
