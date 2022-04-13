package sdk

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type SecretObject struct {
	Identity    string `url:"identity,omitempty"`
	Name        string `url:"name,omitempty"`
	Engine      string `url:"engine,omitempty"`
	Expiration  string `url:"expiration_date"`
	Description string `url:"description"`
	Data        string `url:"data,omitempty"`
}

type SecretObjectData struct {
	AccessKeys  []AccessKeyContainer  `json:"access_keys,omitempty"`
	Credentials []CredentialContainer `json:"credentials,omitempty"`
	KeyValues   KeyValue              `json:"key_value,omitempty"`
}

type AccessKeyContainer struct {
	AccessKey AccessKey `json:"access_key,omitempty"`
}

type AccessKey struct {
	Type            string         `json:"type,omitempty"`
	AccessKeyFields AccessKeyField `json:"fields,omitempty"`
}

type AccessKeyField struct {
	ID          string `json:"access_key_id,omitempty"`
	Secret      string `json:"secret_access_key,omitempty"`
	IDLabel     string `json:"access_key_id_label,omitempty"`
	SecretLabel string `json:"secret_access_key_label,omitempty"`
}

type CredentialContainer struct {
	Credential Credential `json:"credential,omitempty"`
}

type Credential struct {
	CredentialFields CredentialField `json:"fields,omitempty"`
}

type CredentialField struct {
	Username                   string `json:"user,omitempty"`
	Password                   string `json:"password,omitempty"`
	Hostname                   string `json:"host,omitempty"`
	AdditionalInformation      string `json:"additional_information,omitempty"`
	UsernameLabel              string `json:"user_label,omitempty"`
	PasswordLabel              string `json:"password_label,omitempty"`
	HostnameLabel              string `json:"host_label,omitempty"`
	AdditionalInformationLabel string `json:"additional_information_label,omitempty"`
}

type KeyValue struct {
	Fields map[string]string `json:"fields,omitempty"`
}

func NewSecretObject(identity string, name string, engine string, description string, expiration string, secret SecretObjectData) SecretObject {
	return SecretObject{
		identity,
		name,
		engine,
		expiration,
		description,
		secret.Encode(),
	}
}

func NewSecretObjectData(accessKeys []AccessKey, credentials []Credential, keyValues KeyValue) SecretObjectData {
	var accessKeysContainer []AccessKeyContainer
	var credentialsContainer []CredentialContainer

	for _, accessKey := range accessKeys {
		accessKeysContainer = append(accessKeysContainer, AccessKeyContainer{accessKey})
	}

	for _, credential := range credentials {
		credentialsContainer = append(credentialsContainer, CredentialContainer{credential})
	}

	return SecretObjectData{
		accessKeysContainer,
		credentialsContainer,
		keyValues,
	}
}

func NewAccessKey(accessKeyType string, accessKeyID string, secretAccessKey string, accessKeyIDLabel string, secretAccessKeyLabel string) AccessKey {
	accessKeyFields := AccessKeyField{
		accessKeyID,
		secretAccessKey,
		accessKeyIDLabel,
		secretAccessKeyLabel,
	}

	return AccessKey{
		accessKeyType,
		accessKeyFields,
	}
}

func NewCredential(username string, password string, hostname string, addInfo string, usernameLabel string, passwordLabel string, hostnameLabel string, addInfoLabel string) Credential {
	credentials := CredentialField{
		username,
		password,
		hostname,
		addInfo,
		usernameLabel,
		passwordLabel,
		hostnameLabel,
		addInfoLabel,
	}

	return Credential{
		credentials,
	}
}

func NewKeyValue(fields map[string]string) KeyValue {
	return KeyValue{
		fields,
	}
}

func (s SecretObjectData) Encode() string {
	sJSON, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Error converting object to JSON")
	}
	return base64.StdEncoding.EncodeToString(sJSON)
}
