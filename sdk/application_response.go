package sdk

import "encoding/json"

type ApplicationResponse struct {
	ClientID     string      `json:"id,omitempty"`
	ClientSecret string      `json:"signature,omitempty"`
	Application  Application `json:"application,omitempty"`
	Response     Response    `json:"response,omitempty"`
}

type Application struct {
	Name        string   `json:"name,omitempty"`
	System      string   `json:"system,omitempty"`
	Environment string   `json:"environment,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Secrets     []Secret `json:"secrets,omitempty"`
}

type Secret struct {
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
