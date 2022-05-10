package dto

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"scrumlr.io/server/database"
)

// User is the response for all user requests.
type User struct {
	// The id of the user
	ID uuid.UUID `json:"id"`

	// The user's display name
	Name string `json:"name"`

	// The user's avatar configuration
	Avatar json.RawMessage `json:"avatar"`
}

func (u *User) From(user database.User) *User {
	u.ID = user.ID
	u.Name = user.Name
	u.Avatar = user.Avatar
	return u
}

func (*User) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type UserUpdateRequest struct {
	ID     uuid.UUID       `json:"-"`
	Name   string          `json:"name"`
	Avatar json.RawMessage `json:"avatar"`
}
