package actions

import (
	"testing"

	"github.com/datawowio/posmoni-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetListModerationEndpoint(t *testing.T) {
	aiConsensus := &GetModerations{}
	endpoint, method, path := aiConsensus.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.PosmoniAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/v1/moderations", path)
}

func TestPostModerationEndpoint(t *testing.T) {
	aiConsensus := &PostModeration{}
	endpoint, method, path := aiConsensus.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.PosmoniAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/v1/moderations", path)
}

func TestGetListModerationPayload(t *testing.T) {
	g := &GetModerations{
		ID:   "5a44671ab3957c2ab5c33326",
		Item: "1",
		Page: "5",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("query"))
	a.Equal(t, g.Item, queryValues.Get("per_page"))
	a.Equal(t, g.Page, queryValues.Get("page"))
}

func TestPostModerationPayload(t *testing.T) {
	p := &PostModeration{
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		PostbackURL:    "http://someUrl.url",
		PostbackMethod: "POST",
		CustomID:       "custom_id",
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
