package tdarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the Tdarr API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// New creates a new Tdarr client. baseURL should be e.g. "http://localhost:8265".
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// SetToken sets the Bearer token for authenticated requests.
func (c *Client) SetToken(token string) {
	c.token = token
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

// do executes a prepared request.
func (c *Client) do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// post is a helper that builds + sends a POST with a JSON body.
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
	return c.do(req)
}

// get is a helper that builds + sends a GET.
func (c *Client) get(url string) (*http.Response, error) {
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// decode reads a JSON response body into a RawMessage.
func (c *Client) decode(resp *http.Response) (json.RawMessage, error) {
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("tdarr: HTTP %d: %s", resp.StatusCode, string(body))
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
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("tdarr: HTTP %d: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// decodeBool reads a boolean JSON response.
func (c *Client) decodeBool(resp *http.Response) (bool, error) {
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("tdarr: HTTP %d: %s", resp.StatusCode, string(body))
	}
	var out bool
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return false, err
	}
	return out, nil
}

// Ensure unused imports don't break compilation.
var _ = bytes.NewReader
var _ = fmt.Sprintf
