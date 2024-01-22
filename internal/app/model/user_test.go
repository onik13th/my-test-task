package model_test

import (
	"github.com/onik13th/my-test-task/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty surname",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Surname = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty patronymic",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Patronymic = ""
				return u
			},
			isValid: true,
		},
		{
			name: "invalid",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "12345"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}

		})
	}
}
