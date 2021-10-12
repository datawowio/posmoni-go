package actions

import "net/http"

// Action represents necessary methods for each actions.
// The interface contains 2 methods which are Endpoint for getting the endpoint, path,
// and API verb and Payload for creating request payload in each request.
type Action interface {
	Endpoint() (string, string, string)
	Payload(string, string, string) (*http.Request, error)
}
