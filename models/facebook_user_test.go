package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacebookUser(t *testing.T) {
	fbuser := FacebookUser{UID: "olivere", ID: "Take Five", Name: "Bla"}

	assert.Equal(t, fbuser.ID, "Take Five")
	assert.Equal(t, fbuser.UID, "olivere")
}

func TestFacebookUserToJson(t *testing.T) {
	fbuser := FacebookUser{UID: "olivere", ID: "Take Five", Name: "Bla"}

	assert.Equal(t, fbuser.ToJson(), "{\"uid\":\"olivere\",\"id\":\"Take Five\",\"name\":\"Bla\"}")
}
