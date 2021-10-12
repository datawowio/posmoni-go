package actions

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/posmoni/config"
)

const (
	ModerationPath = "/api/v1/moderations"
)

// Example:
//
//  list, get := &posmoni.GetModerations{}, &actions.GetModerations{
//      Page: 1,
//      Item: 20,
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Moderation: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data[0].Attributes)
//
type GetModerations struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  data, post := &posmoni.PostModeration{}, &actions.PostModeration{
//		Data: TestImageDataURL,
//  }
//
//  if err := client.Call(data, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Element: %#v\n", data.Data.Attributes)
//
type PostModeration struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

// Endpoint returns Posmoni's request url, verb and endpoint for calling Get list of
// Moderation API.
func (g *GetModerations) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "GET", ModerationPath
}

// Endpoint returns Posmoni's request url, verb and endpoint for calling Create
// Moderation API.
func (p *PostModeration) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "POST", ModerationPath
}

// Payload creates request's payload for Get list Moderation API. Returns
// http.Request which contains required query parameters.
func (g *GetModerations) Payload(endpoint, method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, string(endpoint)+path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("query", g.ID)
	}
	if g.Page != "" {
		q.Add("page", g.Page)
	}
	if g.Item != "" {
		q.Add("per_page", g.Item)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

// Payload creates request's payload for Create Moderation API. Returns
// http.Request which contains required query parameters.
func (p *PostModeration) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	if p.Data != "" {
		values.Set("data", p.Data)
	}
	if p.PostbackURL != "" {
		values.Set("postback_url", p.PostbackURL)
	}
	if p.PostbackMethod != "" {
		values.Set("postback_method", p.PostbackMethod)
	}
	if p.CustomID != "" {
		values.Set("custom_id", p.CustomID)
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
