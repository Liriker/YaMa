package chats

// UserLinkResponse - The result of a successful request to get links to the user
// Ok - Success flag.
// ID - User id in messenger.
// ChatLink - Link to chat with user.
// CallLink - Link to call user.
type UserLinkResponse struct {
	Ok       bool   `json:"ok"`
	ID       string `json:"id"`
	ChatLink string `json:"chat_link"`
	CallLink string `json:"call_link"`
}

type response struct {
	Ok          bool        `json:"ok"`
	ChatID      string      `json:"chat_id,omitempty"`
	Description interface{} `json:"description,omitempty"`
}
