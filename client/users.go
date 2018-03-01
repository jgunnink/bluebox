package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jgunnink/bluebox"
)

// UserByID will get a user from an ID
func (c *Client) UserByID(id int) (*bluebox.User, error) {
	resp, err := http.Get(c.BaseURL + "/users/" + strconv.Itoa(id) + "/get")
	if err != nil {
		return nil, err
	}
	result := &bluebox.User{}
	json.NewDecoder(resp.Body).Decode(result)
	return result, nil
}

// UserCreate will create a new user
func (c *Client) UserCreate(data *bluebox.UserCreateRequest) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.postHelper("/users/create", bytes.NewReader(payload))
}
