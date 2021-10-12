package posmoni

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/datawowio/posmoni-go/actions"
	"github.com/datawowio/posmoni-go/config"
	a "github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	TestProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
	TestImageID    = "5a44671ab3957c2ab5c33326"
	TestImageData  = "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png"
)

func readFile(path string) []byte {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	return result
}

type mockPayload struct {
	ID       string
	CustomID string
}

func (*mockPayload) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "method", "path"
}

func (m *mockPayload) Payload(endpoint, method, path string) (*http.Request, error) {
	return nil, errors.New("Mock error for payload testing")
}

func TestNewClient(t *testing.T) {
	c, err := NewClient(TestProjectKey)
	a.Nil(t, err)
	a.NotNil(t, c)
}

func TestNewClient_ErrorInvalidKey(t *testing.T) {
	c, err := NewClient("")
	a.NotNil(t, err)
	a.Nil(t, c)
}

func TestClient_CallGetListModeration(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m, getM := &GetModerations{}, &actions.GetModerations{
		ID: TestImageID,
	}

	endpoint, _, path := getM.Endpoint()
	mockResp := readFile("./testdata/moderations.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(m, getM)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, m)
	a.NotNil(t, m.Data[0].ID)
	a.NotNil(t, m.Data[0].Attributes.Answer)
	a.Equal(t, getM.ID, m.Data[0].ID)
}

func TestClient_CallPostModeration(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m, postM := &PostModeration{}, &actions.PostModeration{
		Data: TestImageData,
	}

	endpoint, _, path := postM.Endpoint()
	mockResp := readFile("./testdata/post_moderation.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(m, postM)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, m)
	a.Equal(t, postM.Data, m.Data.Attributes.Source)
}

func TestClient_InvalidCall(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m, getM := &GetModerations{}, &actions.GetModerations{}

	endpoint, _, path := getM.Endpoint()
	gock.New(endpoint).
		Get(path).
		Reply(401)

	e := c.Call(m, getM)
	a.EqualError(t, e, e.Error())
}

func TestClient_InvalidPayload(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m := &mockPayload{}
	e := c.Call(nil, m)
	a.NotNil(t, e)
	a.EqualError(t, e, "Mock error for payload testing")
}
