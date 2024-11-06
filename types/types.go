package types

const (
	PrivateChatType = "private"
	GroupChatType   = "group"
	ChannelChatType = "channel"
)

// Update - It is used in responses to describe the message in the update.
// From - The sender of the message.
// Chat - The chat to which the message was sent.
// Text - The text of the message.
// Timestamp - The time when the message was sent by the server clock: UNIX timestamp.
// MessageID - ID of the chat message.
// UpdateID - update ID.
// File - Information about the file attached to the message.
// Images - Information about the pictures.
type Update struct {
	From      Sender    `json:"from"`
	Chat      Chat      `json:"chat"`
	Text      string    `json:"text,omitempty"`
	Timestamp int64     `json:"timestamp"`
	MessageID int64     `json:"message_id"`
	UpdateID  int64     `json:"update_id"`
	File      File      `json:"file,omitempty"`
	Images    [][]Image `json:"images,omitempty"`
}

// Button - it is used in queries to describe an inline button under a text message.
// Text - text on the inline button.
// CallbackData - the data that will be sent to the server when the button is clicked.
type Button struct {
	Text         string      `json:"text"`
	CallbackData interface{} `json:"callback_data,omitempty"`
}

// Chat - it is used in responses to describe the chat (channel).
// Type - chat type. Possible values: PrivateChatType("private") — private chat; GroupChatType("group") — group chat; ChannelChatType("channel") — the channel.
// ID - chat ID. A chat with the private type does not have a meaningful identifier. There are always two participants in such a chat — the bot and its interlocutor. The interlocutor should be identified by an object of the User type, which is usually located nearby.
type Chat struct {
	Type string `json:"type"`
	ID   string `json:"id,omitempty"`
}

// File - It is used in the responses to describe the file.
// ID - id of the file to upload via the API.
// Name - file name.
// Size - file size in bytes.
type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

// Image - It is used in responses to describe the image.
// FileID - id of the file to upload via the API.
// Width - The width of the image.
// Height - The height of the image.
// Size - file size in bytes.
// Name - The name of the file (as it was when it was uploaded).
type Image struct {
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"size,omitempty"`
	Name   string `json:"name,omitempty"`
}

// Sender - It is used in responses to describe the sender of the message.
// Login - The username of the user who sent the message. Specified for messages from chats.
// ID - id of the channel whose administrator sent the message. Specified for messages in channels.
// You will receive only one of the login or id parameters in the response, depending on where the message was sent — to the chat or channel.
// DisplayName - sender's display name.
// Robot - indicates whether the sender is a bot.
type Sender struct {
	Login       string `json:"login,omitempty"`
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Robot       bool   `json:"robot,omitempty"`
}

// Vote - It is used in responses to describe the person who voted in the survey.
// Timestamp - Voice ID.
// User - The user who voted.
type Vote struct {
	Timestamp int64  `json:"timestamp"`
	User      string `json:"user"`
}

// User - It is used in queries to describe the user.
// Login - user's login. For Yandex accounts (domain yandex.ru ) logins can be used without specifying a domain.
// For accounts created on other domains, the full login form <login>@<domain> is specified.
// The mailing address of a group or division can also be specified as login, then the group or division will be used as User.
type User struct {
	Login string `json:"login"`
}

// NewChat - struct to chat/channel creation.
// Name - name of the chat (channel). No more than 200 characters.
// Description - description of the chat (channel) No more than 500 characters, an empty string is allowed.
// AvatarUrl - chat icon (channel) image URL.
// Admins - list of chat administrators (channel).
// Members - the list of chat participants. Should be empty if a channel is being created instead of a chat (channel=true).
// Channel - flag for creating a channel instead of a chat
// Subscribers - the list of subscribers of the channel. Should be empty if a chat is being created (channel=false).
type NewChat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	Admins      []User `json:"admins,omitempty"`
	Members     []User `json:"members,omitempty"`
	Channel     bool   `json:"channel,omitempty"`
	Subscribers []User `json:"subscribers,omitempty"`
}

// ChatUpdate - struct for chat/channel management
// ChatID - ID of the chat (channel). The bot must be in the chat (channel)
// Members - The list of users who need to be made participants in the chat Bot must be in the chat. If at least one of the users (User) was an administrator, the bot must be an admin.
// Admins - The list of users who need to be made chat (channel) administrators. The bot must be the administrator of the chat (channel).
// Subscribers - The list of users who need to be made subscribers of the Bot channel must be in the chat. Users cannot be chat administrators.
// Remove - The list of users to be removed from the chat (channel) To remove administrators, the bot must be the administrator of the chat (channel).
// The members, admins, subscribers and remove parameters are optional, but at least one of the lists must be set.
type ChatUpdate struct {
	ChatID      string `json:"chat_id"`
	Members     []User `json:"members,omitempty"`
	Admins      []User `json:"admins,omitempty"`
	Subscribers []User `json:"subscribers,omitempty"`
	Remove      []User `json:"remove,omitempty"`
}

// NewMessage - struct for message creation.
// ChatID - group chat ID The bot must be a chat participant.
// Login - user login.
// The chat_id and login parameters are optional, but at least one of the two must be filled in:
// When filling in the chat_id, the message will be sent to the group chat specified by this ID.
// When filling in the login, a message will be sent to the user in a private chat.
// Text - message text.
// PayloadID - request ID The ID must be unique for each request. Requests with the same ID are treated as duplicates.
// ReplyMessageID - ID of the message to be answered. The message must be from the same chat.
// DisableNotification -  Whether to disable the notification. Default value: false.
// Important -  Is the message important. Default value: false.
// DisableWebPagePreview - Disable link disclosure in the message. Default value: false.
// ThreadID - ID of the thread (timestamp of the message).
// InlineKeyboard - An array of inline buttons under the message, which can be used to send a quick response.
type NewMessage struct {
	ChatID                string   `json:"chat_id,omitempty"`
	Login                 string   `json:"login,omitempty"`
	Text                  string   `json:"text"`
	PayloadID             string   `json:"payload_id,omitempty"`
	ReplyMessageID        string   `json:"reply_message_id,omitempty"`
	DisableNotification   bool     `json:"disable_notification,omitempty"`
	Important             bool     `json:"important,omitempty"`
	DisableWebPagePreview bool     `json:"disable_web_page_preview,omitempty"`
	ThreadID              string   `json:"thread_id,omitempty"`
	InlineKeyboard        []Button `json:"inline_keyboard,omitempty"`
}

// NewFileMessage - struct for file-message creation.
// ChatID - group chat ID The bot must be a chat participant.
// Login - user login.
// The chat_id and login parameters are optional, but at least one of the two must be filled in:
// When filling in the chat_id, the message will be sent to the group chat specified by this ID.
// When filling in the login, a message will be sent to the user in a private chat.
// Document - file contents.
// ThreadID - ID of the thread (timestamp of the message).
type NewFileMessage struct {
	ChatID   string `json:"chat_id,omitempty"`
	Login    string `json:"login,omitempty"`
	Document []byte `json:"document"`
	ThreadID int64  `json:"thread_id,omitempty"`
}

// NewImageMessage - struct for image-message creation.
// ChatID - group chat ID The bot must be a chat participant.
// Login - user login.
// The chat_id and login parameters are optional, but at least one of the two must be filled in:
// When filling in the chat_id, the message will be sent to the group chat specified by this ID.
// When filling in the login, a message will be sent to the user in a private chat.
// Image - file contents.
// ThreadID - ID of the thread (timestamp of the message).
type NewImageMessage struct {
	ChatID   string `json:"chat_id,omitempty"`
	Login    string `json:"login,omitempty"`
	Image    []byte `json:"image"`
	ThreadID int64  `json:"thread_id,omitempty"`
}

// NewGalleryMessage - struct for image-slice-message creation.
// ChatID - group chat ID The bot must be a chat participant.
// Login - user login.
// The chat_id and login parameters are optional, but at least one of the two must be filled in:
// When filling in the chat_id, the message will be sent to the group chat specified by this ID.
// When filling in the login, a message will be sent to the user in a private chat.
// Image - slice of files.
// ThreadID - ID of the thread (timestamp of the message).
type NewGalleryMessage struct {
	ChatID   string   `json:"chat_id,omitempty"`
	Login    string   `json:"login,omitempty"`
	Images   [][]byte `json:"images"`
	ThreadID int64    `json:"thread_id,omitempty"`
}

// NewDeleteMessageRequest - struct for message deleting.
// ChatID - group chat ID The bot must be a chat participant.
// Login - user login.
// The chat_id and login parameters are optional, but at least one of the two must be filled in:
// When filling in the chat_id, the message will be sent to the group chat specified by this ID.
// When filling in the login, a message will be sent to the user in a private chat.
// MessageID - ID of the message (timestamp).
// ThreadID - ID of the thread (timestamp of the message).
type NewDeleteMessageRequest struct {
	ChatID    string `json:"chat_id,omitempty"`
	Login     string `json:"login,omitempty"`
	MessageID int64  `json:"message_id"`
	ThreadID  int64  `json:"thread_id,omitempty"`
}
