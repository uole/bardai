package bardai

type chatRequest struct {
	Prompt      string `json:"prompt"`
	AccessToken string `json:"access_token"`
}

type chatResponse struct {
	Prompt   string `json:"prompt"`
	Answer   string `json:"answer"`
	Duration string `json:"duration"`
}
