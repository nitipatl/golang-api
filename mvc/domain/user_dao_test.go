package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id = 0")

	assert.NotNil(t, err, "we were expecting an eror with id = 0")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user id is not found")

	// if user != nil {
	// 	t.Error("we were not expecting a user with id = 0")
	// }

	// if err == nil {
	// 	t.Error("we were expecting an eror with id = 0")
	// }

	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user id is not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Nitipat", user.FirstName)
	assert.EqualValues(t, "L", user.LastName)
	assert.EqualValues(t, "iamsvz@gmail.com", user.Email)
}
