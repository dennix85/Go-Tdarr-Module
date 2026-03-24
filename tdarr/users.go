package tdarr

import (
	"encoding/json"
)

// AdminRegisterRequest is the request body for this endpoint.
type AdminRegisterRequest struct {
	Username string `json:"username"`
	Roles []string `json:"roles"`
}

// AdminResetPasswordRequest is the request body for this endpoint.
type AdminResetPasswordRequest struct {
	Username string `json:"username"`
}

// AuthResetPasswordRequest is the request body for this endpoint.
type AuthResetPasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	Password string `json:"password"`
}

// PublicAuthLoginRequest is the request body for this endpoint.
type PublicAuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PublicAuthRegisterRequest is the request body for this endpoint.
type PublicAuthRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminRegister - For admin registering a new user
func (c *Client) AdminRegister(req AdminRegisterRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/admin/register", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// AdminResetPassword - For admin resetting a user password
func (c *Client) AdminResetPassword(req AdminResetPasswordRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/admin/reset-password", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// AuthLogout - For logging out a user
func (c *Client) AuthLogout() (string, error) {
	resp, err := c.post(c.baseURL + "/api/v2/auth/logout", nil)
	if err != nil {
		return nil, err
	}
	return c.decodeString(resp)
}

// AuthResetPassword - For resetting a user password
func (c *Client) AuthResetPassword(req AuthResetPasswordRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/auth/reset-password", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// AuthVerifyToken - For verifying a user token
func (c *Client) AuthVerifyToken() (json.RawMessage, error) {
	resp, err := c.get(c.baseURL + "/api/v2/auth/verify-token")
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// PublicAuthLogin - For logging in a user
func (c *Client) PublicAuthLogin(req PublicAuthLoginRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/public/auth/login", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}

// PublicAuthRegister - For registering a new user
func (c *Client) PublicAuthRegister(req PublicAuthRegisterRequest) (json.RawMessage, error) {
	resp, err := c.post(c.baseURL + "/api/v2/public/auth/register", req)
	if err != nil {
		return nil, err
	}
	return c.decode(resp)
}
