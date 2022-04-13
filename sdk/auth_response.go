package sdk

import "encoding/json"

type AuthResponse struct {
	ID        string   `json:"id,omitempty"`
	Error     string   `json:"error,omitempty"`
	Message   string   `json:"message,omitempty"`
	Reason    string   `json:"reason,omitempty"`
	ExpiresIn int      `json:"expires_in,omitempty"`
	Signature string   `json:"signature,omitempty"`
	TokenType string   `json:"token_type,omitempty"`
	Token     string   `json:"access_token"`
	Response  Response `json:"response,omitempty"`
}

func (r *AuthResponse) Unmarshal(resBody []byte) error {
	err := json.Unmarshal(resBody, &r)
	if err != nil {
		return err
	}

	return nil
}
