package client

import (
	"bytes"
	"encoding/json"

	"github.com/jgunnink/bluebox"
)

// UserSignIn will create a new user
func (c *Client) UserSignIn(data *bluebox.UserSignInRequest) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.postHelper("/auth/sign_in", bytes.NewReader(payload))
}
