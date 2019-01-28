package models

// SlackMessage .
type SlackMessage struct {
	Fallback string `json:"fallback"` //Required!
	Text     string `json:"text"`
	Color    string `json:"color"`
	Fields   struct {
		Title string `json:"title"`
		Value string `json:"value"`
		Short bool   `json:"short"`
	}
}

// Channel  string `json:"channel"`
// Emoji           string    `json:"emoji"`
// AssignmentNotification to ass.
type AssignmentNotification struct {
	SlackMessage    *SlackMessage `json:"slackMessage"`
	AssignmentLink  string        `json:"assignmentLink"`
	HighligtedNames []string      `json:"highlightedNames"`
}

// SimpleSlackMsg to ass.
type SimpleSlackMsg struct {
	Text string `json:"text"`
}

// SlackUser ..
type SlackUser struct {
	DisplayName []string `json:"displayName"`
	NumOfUsers  int      `json:"numOfUsers"`
}
