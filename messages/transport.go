package messages

type response struct {
	Ok          bool   `json:"ok"`
	MessageID   int64  `json:"message_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type getFileRequest struct {
	FileID int64 `json:"file_id"`
}
