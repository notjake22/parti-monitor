package discord

type Webhook struct {
	Content     string        `json:"content"`
	Embeds      []Embed       `json:"embeds"`
	Username    string        `json:"username"`
	AvatarURL   string        `json:"avatar_url"`
	Attachments []interface{} `json:"attachments"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
	Footer      struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	} `json:"footer"`
	Thumbnail struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
