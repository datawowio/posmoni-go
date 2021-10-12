package posmoni

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/datawowio/posmoni-go/actions"
)

// Client is like a API Gateway for Posmoni. Client will let you call available
// Posmoni's APIs. It will be used with action structures from actions sub-package.
type Client struct {
	*http.Client
	ProjectKey string
}

// NewClient is used for creates and return a Client with given project key
func NewClient(projectKey string) (*Client, error) {
	if projectKey == "" {
		return nil, errors.New("invalid project key")
	}
	return &Client{&http.Client{}, projectKey}, nil
}

// Call performs supplied operations against Posmoni's API and unmarshal response into
// given action object.
//
// In successful case, result will contain 2 main objects, data and meta. (status code and
// message) Failed case, response will contain an error message.
func (c *Client) Call(result interface{}, act actions.Action) error {
	endpoint, method, path := act.Endpoint()
	req, err := act.Payload(endpoint, method, path)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.ProjectKey)
	if req.Header.Get("Content-Type") == "" {
		c.setContentType(req)
	}

	resp, e := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if e != nil {
		return e
	}

	buffer, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Println("Error while reading response body")
	}

	switch {
	case e != nil:
		return e
	case resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated:
		err := errors.New(resp.Status)
		return err
	}

	if result != nil {
		if e := json.Unmarshal(buffer, result); e != nil {
			return e
		}
	}

	return nil
}

func (c *Client) setContentType(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
