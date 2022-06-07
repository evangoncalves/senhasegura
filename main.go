package main

import (
	"fmt"

	senhaseguraSDK "github.com/evangoncalves/senhasegura/sdk"
)

func main() {

	client := senhaseguraSDK.NewClient("https://10.0.0.161/", "0c2fc858789bf762aacb16ef95099ced06254bf04", "43627168c56608ac1eee1ab2c1ef7889")

	for _, secret := range client.GetApplication().Secrets {
		for _, data := range secret.Data {
			for key, value := range data {
				fmt.Println(key + "=" + value)
			}
		}
	}

	app := senhaseguraSDK.NewApplication("Test2", "Demonstration", "Development", "", "", "", "", "", []string{"test"}, []string{}, []string{}, "", []string{}, []string{}, "")

	client.CreateApplication(app)

	kv := make(map[string]string)
	kv["KEY"] = "VALUE"
	kv["TEST_KEY"] = "TEST_VALUE"
	kv["NEW_KEY"] = "NEW_VALUE2"

	secret := senhaseguraSDK.NewSecretData(
		[]senhaseguraSDK.AccessKey{
			senhaseguraSDK.NewAccessKey("aws", "AKIASQGYZVJA43ISDLOK", "", "AWS_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"),
		},
		[]senhaseguraSDK.Credential{
			senhaseguraSDK.NewCredential("api-user", "", "api-server", "TEST INFO", "USER", "PASS", "HOST", "CONN"),
		},
		senhaseguraSDK.NewKeyValue(kv),
	)

	fmt.Println(client.CreateSecret(senhaseguraSDK.NewSecret("senhasegura/senhaseguraSDK", "", "Generic", "This is a Secret created through the senhaseguraSDK", "", secret)))
}
