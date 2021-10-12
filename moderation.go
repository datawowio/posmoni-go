package posmoni

// GetModerations represents the response object that returned from Get List of Moderation API.
type GetModerations struct {
	Data Moderations `json:"data"`
	Meta Meta        `json:"meta"`
}

// PostModeration respresents the created object that returned from Create Moderation API.
type PostModeration struct {
	Data Moderation `json:"data"`
	Meta Meta       `json:"meta"`
}

// Moderation represents Moderation object.
type Moderation struct {
	ID         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes ModerationAttributes `json:"attributes"`
}

// Moderations represents list of Moderation object.
type Moderations []Moderation

// ModerationAttributes represents available attributes of Moderation
type ModerationAttributes struct {
	Answer         string `json:"answer"`
	CustomID       string `json:"custom_id"`
	Source         string `json:"data"`
	Postback       bool   `json:"postback"`
	PostbackURL    string `json:"postback_url"`
	PostbackMethod string `json:"postback_method"`
	ProcessedAt    string `json:"processed_at"`
	ProjectID      int    `json:"project_id"`
	Status         string `json:"status"`
}
