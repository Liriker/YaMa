package updates

import "github.com/Liriker/YaMa/types"

type updateRequest struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type updateResponse struct {
	Ok          bool           `json:"ok"`
	Updates     []types.Update `json:"updates,omitempty"`
	Description string         `json:"description,omitempty"`
}

type webhookRequest struct {
	WebhookUrl string `json:"webhook_url"`
}

type webhookResponse struct {
	Ok            bool   `json:"ok"`
	Id            string `json:"id"`
	DisplayName   string `json:"display_name"`
	WebhookUrl    string `json:"webhook_url,omitempty"`
	Organizations []int  `json:"organizations"`
	Login         string `json:"login"`
}
