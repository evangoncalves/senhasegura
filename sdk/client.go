package sdk

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

type Client struct {
	URL          string
	ClientID     string
	ClientSecret string
	Token        string
}

func NewClient(url string, clientID string, clientSecret string) Client {
	if url == "" {
		log.Fatal("URL cannot be empty")
	}
	if clientID == "" {
		log.Fatal("Client ID cannot be empty")
	}
	if clientSecret == "" {
		log.Fatal("Client Secret cannot be empty")
	}

	c := Client{
		url,
		clientID,
		clientSecret,
		"",
	}

	return c
}

func (c *Client) Authenticate() {
	endpoint := "/iso/oauth2/token"

	reqBody := url.Values{}
	reqBody.Set("grant_type", "client_credentials")
	reqBody.Set("client_id", c.ClientID)
	reqBody.Set("client_secret", c.ClientSecret)

	var response AuthResponse

	fmt.Println("Authenticating...")

	err := c.Request(http.MethodPost, endpoint, reqBody, &response)
	if err != nil {
		log.Fatalf("An error has ocurred when trying to authenticate %v", err)
	}

	c.Token = response.TokenType + " " + response.Token
}

func (c Client) GetApplication() Application {

	c.Authenticate()

	endpoint := "/iso/dapp/application"

	var response ApplicationResponse

	err := c.Request(http.MethodGet, endpoint, url.Values{}, &response)
	if err != nil {
		log.Fatalf("An error has ocurred when trying to get application %v", err)
	}

	return response.Application
}

func (c *Client) CreateApplication(application ApplicationObject) {

	if application.Name == "" {
		log.Fatal("Application name not set")
	}
	if application.System == "" {
		log.Fatal("Application system not set")
	}
	if application.Environment == "" {
		log.Fatal("Application environment not set")
	}

	c.Authenticate()

	endpoint := "/iso/dapp/application"

	reqBody, err := query.Values(application)
	if err != nil {
		log.Fatal("Error when trying to create an application")
	}

	var response ApplicationResponse

	err = c.Request(http.MethodPost, endpoint, reqBody, &response)
	if err != nil {
		log.Fatalf("An error has ocurred when trying to create an application %v", err)
	}

	c.ClientID = response.ClientID
	c.ClientSecret = response.ClientSecret
}

func (c Client) GetSecret() Secret {
	return nil
}

func (c Client) CreateSecret(secret SecretObject) Application {
	if secret.Data == "" {
		log.Fatal("Secret data not set")
	}

	c.Authenticate()

	endpoint := "/iso/sctm/secret"

	reqBody, err := query.Values(secret)

	var response ApplicationResponse

	err = c.Request(http.MethodPost, endpoint, reqBody, &response)
	if err != nil {
		log.Fatal("An error has ocurred when trying to create a secret")
	}

	return response.Application
}

func (c Client) Request(method string, endpoint string, reqBody url.Values, response IResponse) error {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Content-Length"] = strconv.Itoa(len(reqBody.Encode()))
	if c.Token != "" {
		headers["Authorization"] = c.Token
	}

	url, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return err
	}

	url.Path = endpoint

	uri := url.String()

	fmt.Println(fmt.Sprintf("%s to %v", method, uri))

	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest(method, uri, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return err
	}

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return response.Unmarshal(resBody)
}
