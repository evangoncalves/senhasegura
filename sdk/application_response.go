package sdk

import "encoding/json"

type ApplicationResponse struct {
	ClientID     string                  `json:"id,omitempty"`
	ClientSecret string                  `json:"signature,omitempty"`
	Application  ApplicationResponseData `json:"application,omitempty"`
	Response     Response                `json:"response,omitempty"`
}

type ApplicationResponseData struct {
	Name        string               `json:"name,omitempty"`
	System      string               `json:"system,omitempty"`
	Environment string               `json:"environment,omitempty"`
	Description string               `json:"description,omitempty"`
	Tags        []string             `json:"tags,omitempty"`
	Secrets     []SecretResponseData `json:"secrets,omitempty"`
}

// An object containing all secret data from senhasegura, including its name, identity, version, expiration date, engine and sensitive information as data property
type SecretResponseData struct {
	ID         string              `json:"secret_id,omitempty"`
	Name       string              `json:"secret_name,omitempty"`
	Identity   string              `json:"identity,omitempty"`
	Version    string              `json:"version,omitempty"`
	Expiration string              `json:"expiration_date,omitempty"`
	Engine     string              `json:"engine,omitempty"`
	Data       []map[string]string `json:"data,omitempty"`
}

func (r *ApplicationResponse) Unmarshal(resBody []byte) error {
	err := json.Unmarshal(resBody, &r)
	if err != nil {
		return err
	}

	return nil
}
