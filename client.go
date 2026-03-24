package tdarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the Tdarr API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// Option is a functional option for configuring a Client.
type Option func(*Client)

// WithToken sets the Bearer token (API key or JWT) used for authenticated requests.
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// WithTimeout sets a custom HTTP timeout (default: 30s).
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// WithHTTPClient replaces the underlying http.Client entirely.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// New creates a new Tdarr client. baseURL should be e.g. "http://localhost:8265".
//
// Example:
//
//	client := tdarr.New("http://localhost:8265",
//	    tdarr.WithToken("tapi_xxxx"),
//	    tdarr.WithTimeout(10*time.Second),
//	)
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SetToken sets the Bearer token at runtime (e.g. after login).
func (c *Client) SetToken(token string) {
	c.token = token
}

// APIError is returned when the server responds with a non-2xx status.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("tdarr: HTTP %d: %s", e.StatusCode, e.Message)
}

// newRequest builds an http.Request with auth headers attached.
func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return req, nil
}

// post builds and sends a POST with a JSON body.
func (c *Client) post(url string, payload interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(b)
	}
	req, err := c.newRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

// get builds and sends a GET.
func (c *Client) get(url string) (*http.Response, error) {
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

// checkStatus returns an *APIError if the response is non-2xx.
func checkStatus(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	// Try to extract a JSON {"error":"..."} message.
	var apiErr struct {
		Error string `json:"error"`
	}
	if json.Unmarshal(body, &apiErr) == nil && apiErr.Error != "" {
		return &APIError{StatusCode: resp.StatusCode, Message: apiErr.Error}
	}
	return &APIError{StatusCode: resp.StatusCode, Message: string(body)}
}

// decode reads a JSON response body into a RawMessage.
func (c *Client) decode(resp *http.Response) (json.RawMessage, error) {
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	var out json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// decodeString reads a plain-text or JSON-string response body.
func (c *Client) decodeString(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// decodeBool reads a boolean JSON response.
func (c *Client) decodeBool(resp *http.Response) (bool, error) {
	defer resp.Body.Close()
	if err := checkStatus(resp); err != nil {
		return false, err
	}
	var out bool
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return false, err
	}
	return out, nil
}


// decodeJSON decodes a response body into a typed value.
// The caller is responsible for calling resp.Body.Close() and checkStatus before this.
func decodeJSON(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
