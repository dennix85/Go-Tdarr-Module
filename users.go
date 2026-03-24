package tdarr

// AdminRegisterRequest is the request body for AdminRegister.
type AdminRegisterRequest struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// AdminResetPasswordRequest is the request body for AdminResetPassword.
type AdminResetPasswordRequest struct {
	Username string `json:"username"`
}

// AuthResetPasswordRequest is the request body for AuthResetPassword.
type AuthResetPasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

// PublicAuthLoginRequest is the request body for PublicAuthLogin.
type PublicAuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PublicAuthRegisterRequest is the request body for PublicAuthRegister.
type PublicAuthRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminRegister - For admin registering a new user.
func (c *Client) AdminRegister(req AdminRegisterRequest) (*TempPasswordResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/admin/register", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out TempPasswordResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdminResetPassword - For admin resetting a user password.
func (c *Client) AdminResetPassword(req AdminResetPasswordRequest) (*TempPasswordResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/admin/reset-password", map[string]interface{}{"data": req})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out TempPasswordResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AuthLogout - For logging out a user.
func (c *Client) AuthLogout() (string, error) {
	resp, err := c.post(c.baseURL+"/api/v2/auth/logout", nil)
	if err != nil {
		return "", err
	}
	return c.decodeString(resp)
}

// AuthResetPassword - For resetting a user password. Returns the new JWT token.
func (c *Client) AuthResetPassword(req AuthResetPasswordRequest) (*TokenResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/auth/reset-password", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out TokenResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AuthVerifyToken - For verifying a user token.
func (c *Client) AuthVerifyToken() (*VerifyTokenResponse, error) {
	resp, err := c.get(c.baseURL + "/api/v2/auth/verify-token")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out VerifyTokenResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PublicAuthLogin - For logging in a user. Returns a JWT token.
// After a successful login, call client.SetToken(result.Token) to authenticate future requests.
func (c *Client) PublicAuthLogin(req PublicAuthLoginRequest) (*TokenResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/public/auth/login", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out TokenResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PublicAuthRegister - For registering a new user. Returns a JWT token.
func (c *Client) PublicAuthRegister(req PublicAuthRegisterRequest) (*TokenResponse, error) {
	resp, err := c.post(c.baseURL+"/api/v2/public/auth/register", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out TokenResponse
	if err := decodeJSON(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
