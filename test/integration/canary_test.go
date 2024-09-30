package integration

import (
	"context"
	"crypto/tls"
	"errors"
	basistheory "github.com/Basis-Theory/go-sdk"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)
import (
	basistheoryclient "github.com/Basis-Theory/go-sdk/client"
	"github.com/Basis-Theory/go-sdk/option"
)

func TestTenantSelf(t *testing.T) {
	client := NewManagementClient()

	response, _ := client.Tenants.Self.Get(
		context.TODO(),
	)

	if response.Name == nil || *response.Name != `SDK Integration Tests` {
		t.Errorf("Expected SDK Integration Tests. Got %s", *response.Name)
	}
}

func TestTokenCrud(t *testing.T) {
	client := NewPrivateClient()
	manageClient := NewManagementClient()

	// Create Token
	createCardNumber := "6011000990139424"
	tokenId := CreateToken(t, client, createCardNumber)
	GetAndValidateCardNumber(t, client, tokenId, createCardNumber)

	// Update Token
	updateCardNumber := "4242424242424242"
	UpdateToken(t, client, tokenId, updateCardNumber)
	GetAndValidateCardNumber(t, client, tokenId, updateCardNumber)

	// Create Application
	applicationId := CreateApplication(t, manageClient)

	// Create / Delete Proxy
	proxyId := CreateProxy(t, manageClient, applicationId)
	// Missing ability to call proxies from SDK
	// The definition is missing from the Basis Theory OpenApi specs
	DeleteProxy(t, manageClient, proxyId)

	// Reactors
	reactorId := CreateReactor(t, manageClient, applicationId)
	React(t, client, reactorId)
	DeleteReactor(t, manageClient, reactorId)

	DeleteApplication(t, manageClient, applicationId)

	// Delete Token
	DeleteToken(t, client, tokenId)
	EnsureTokenDeleted(t, client, tokenId)
}

func TestIdempotencyHeader(t *testing.T) {
	client := NewPrivateClient()
	idempotencyKey := uuid.NewString()

	firstTokenId := CreateToken(t, client, "6011000990139424", option.WithIdempotencyKey(StringPtr(idempotencyKey)))
	secondTokenId := CreateToken(t, client, "4242424242424242", option.WithIdempotencyKey(StringPtr(idempotencyKey)))

	if firstTokenId != secondTokenId {
		t.Errorf("Expected firstTokenId to be %s but was %s", firstTokenId, secondTokenId)
	}
}

func TestListV1PaginationWithIteration(t *testing.T) {
	client := NewPrivateClient()

	pageSize := 3
	page, err := client.Tokens.List(
		context.TODO(),
		&basistheory.TokensListRequest{
			Page: IntPtr(1),
			Size: IntPtr(pageSize),
		},
	)
	FailIfError(t, "Unable to list tokens", err)

	count := 0
	iter := page.Iterator()
	for iter.Next(context.TODO()) {
		count++
		if count > pageSize {
			break
		}
	}
	if count <= pageSize {
		t.Errorf("Expected at least %d tokens. Got %d", pageSize, count)
	}
}

func TestListV2PaginationWithIteration(t *testing.T) {
	client := NewPrivateClient()

	pageSize := 3
	page, err := client.Tokens.ListV2(
		context.TODO(),
		&basistheory.TokensListV2Request{
			Start: nil,
			Size:  IntPtr(pageSize),
		},
	)
	FailIfError(t, "Unable to list tokens", err)

	count := 0
	iter := page.Iterator()
	for iter.Next(context.TODO()) {
		count++
		if count > pageSize {
			break
		}
	}
	if count <= pageSize {
		t.Errorf("Expected at least %d tokens. Got %d", pageSize, count)
	}
}

func TestWebhooks(t *testing.T) {
	client := NewManagementClient()

	url := "https://echo.basistheory.com/" + uuid.NewString()
	webhookId := CreateWebhook(t, client, url)

	GetWebhookAssertUrl(t, client, webhookId, url)

	time.Sleep(2 * time.Second) // Required to avoid error `The webhook subscription is undergoing another concurrent operation. Please wait a few seconds, then try again.`

	updateUrl := "https://echo.basistheory.com/" + uuid.NewString()
	UpdateWebhook(t, client, webhookId, updateUrl)

	GetWebhookAssertUrl(t, client, webhookId, updateUrl)

	time.Sleep(2 * time.Second) // Required to avoid error `The webhook subscription is undergoing another concurrent operation. Please wait a few seconds, then try again.`

	DeleteWebhook(t, client, webhookId)

	_, err := client.Webhooks.Get(
		context.TODO(),
		webhookId)
	if err == nil {
		t.Errorf("Expected error when trying to get a webhook that doesn't exist")
	}
	var notFoundError basistheory.NotFoundError
	if errors.As(err, &notFoundError) {
		t.Errorf("Expected error to be Not Found")
	}
}

func CreateProxy(t *testing.T, manageClient *basistheoryclient.Client, applicationId string) string {
	response, err := manageClient.Proxies.Create(
		context.TODO(),
		&basistheory.CreateProxyRequest{
			Name:              "(Deletable) go-sdk-" + uuid.NewString(),
			DestinationURL:    "https://echo.basistheory.com",
			RequestReactorID:  nil,
			ResponseReactorID: nil,
			RequestTransform:  nil,
			ResponseTransform: nil,
			Application: &basistheory.Application{
				ID: StringPtr(applicationId),
			},
			Configuration: nil,
			RequireAuth:   nil,
		},
	)
	FailIfError(t, "Failed to create proxy", err)
	proxyId := *response.ID
	return proxyId
}

func NewManagementClient() *basistheoryclient.Client {
	client := basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_MGT_API_KEY")),
		option.WithBaseURL(os.Getenv("BT_API_URL")),
	)
	return client
}

// This management client can be used with a capturing proxy to inspect the requests
// Set the environment variable `HTTP_PROXY` to the proxy URL; ie `http://127.0.0.1:8008`
// This client will disable SSL verification to avoid self-signed (or unavailable) certs on the proxy
func NewManagementClientDebug() *basistheoryclient.Client {
	proxyURL, _ := url.Parse(os.Getenv("HTTP_PROXY"))
	client := basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_MGT_API_KEY")),
		option.WithBaseURL(os.Getenv("BT_API_URL")),
		option.WithHTTPClient(&http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxyURL),
		}}),
	)
	return client
}

func NewPrivateClient() *basistheoryclient.Client {
	client := basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_PVT_API_KEY")),
		option.WithBaseURL(os.Getenv("BT_API_URL")),
	)
	return client
}

func CreateToken(t *testing.T, client *basistheoryclient.Client, cardNumber string, opts ...option.IdempotentRequestOption) string {
	tokenCreatedResponse, err := client.Tokens.Create(
		context.TODO(),
		&basistheory.CreateTokenRequest{
			ID:   nil,
			Type: StringPtr("card"),
			Data: map[string]interface{}{
				"number":           cardNumber,
				"expiration_month": 4,
				"expiration_year":  2025,
				"cvc":              "123",
			},
			Metadata: map[string]*string{
				"customer_id": StringPtr("3181"),
			},
			SearchIndexes: []string{
				"{{ data.expiration_month }}", "{{ data.number | last4 }}",
			},
			FingerprintExpression: StringPtr("{{ data.number }}"),
			Mask: map[string]interface{}{
				"number": StringPtr("{{ data.number, reveal_last: 4 }}"),
				"cvc":    StringPtr("{{ data.cvc }}"),
			},
			DeduplicateToken: BoolPtr(false),
			ExpiresAt:        nil,
			Containers:       []string{"/pci/high/"},
		},
		opts...,
	)
	FailIfError(t, "Failed to create token", err)
	tokenId := *tokenCreatedResponse.ID
	return tokenId
}

func UpdateToken(t *testing.T, client *basistheoryclient.Client, tokenId string, updateCardNumber string) {
	_, err := client.Tokens.Update(
		context.TODO(),
		tokenId,
		&basistheory.UpdateTokenRequest{
			Data: map[string]interface{}{
				"number":           updateCardNumber,
				"expiration_month": 6,
				"expiration_year":  2026,
				"cvc":              "987",
			},
		},
	)
	FailIfError(t, "Failed to update token", err)
}

func DeleteToken(t *testing.T, client *basistheoryclient.Client, tokenId string) {
	err := client.Tokens.Delete(
		context.TODO(),
		tokenId,
	)
	FailIfError(t, "Failed to delete token", err)
}

func EnsureTokenDeleted(t *testing.T, client *basistheoryclient.Client, tokenId string) {
	_, err := client.Tokens.Get(
		context.TODO(),
		tokenId)
	if err == nil {
		t.Errorf("Expected error when trying to get a token that doesn't exist")
	}
	var notFoundError basistheory.NotFoundError
	if errors.As(err, &notFoundError) {
		t.Errorf("Expected error to be Not Found")
	}
}

func GetAndValidateCardNumber(t *testing.T, client *basistheoryclient.Client, tokenId string, createCardNumber string) {
	token, err := client.Tokens.Get(
		context.TODO(),
		tokenId,
	)
	FailIfError(t, "Failed to get token", err)
	data, ok := token.Data.(map[string]interface{})
	number, ok := data["number"].(string)
	if !ok {
		t.Fatalf("Data does not contain card number: %v", err)
	}
	if number != createCardNumber {
		t.Fatalf("Failed to create token: %v", err)
	}
}

func CreateApplication(t *testing.T, manageClient *basistheoryclient.Client) string {
	x, err := manageClient.Applications.Create(
		context.TODO(),
		&basistheory.CreateApplicationRequest{
			Name:        "(Deletable) Go-SDK-" + uuid.NewString(),
			Type:        "private",
			Permissions: []string{"token:use"},
		},
	)
	FailIfError(t, "Failed to create application", err)
	applicationId := *x.ID
	return applicationId
}

func DeleteApplication(t *testing.T, manageClient *basistheoryclient.Client, applicationId string) {
	e := manageClient.Applications.Delete(
		context.TODO(),
		applicationId,
	)
	FailIfError(t, "Failed to delete application", e)
}

func DeleteProxy(t *testing.T, manageClient *basistheoryclient.Client, proxyId string) {
	err := manageClient.Proxies.Delete(
		context.TODO(),
		proxyId,
	)
	FailIfError(t, "Failed to delete proxy", err)
}

func CreateReactor(t *testing.T, manageClient *basistheoryclient.Client, applicationId string) string {
	x, err2 := manageClient.Reactors.Create(
		context.TODO(),
		&basistheory.CreateReactorRequest{
			Name: "(Deletable) Go-SDK-" + uuid.NewString(),
			Code: "module.exports = function (req) {return {raw: req.args}}",
			Application: &basistheory.Application{
				ID: StringPtr(applicationId),
			},
			Configuration: nil,
		})
	FailIfError(t, "Failed to create reactor", err2)
	return *x.ID
}

func React(t *testing.T, client *basistheoryclient.Client, reactorId string) {
	args := map[string]interface{}{
		"key1": StringPtr("Key1-" + uuid.New().String()),
		"key2": StringPtr("Key2-" + uuid.New().String()),
	}
	x, err := client.Reactors.React(
		context.TODO(),
		reactorId,
		&basistheory.ReactRequest{
			Args: args,
		},
	)
	FailIfError(t, "Failed to react in Reactor", err)
	actual, _ := x.Raw.(map[string]interface{})
	if *args["key1"].(*string) != actual["key1"].(string) {
		t.Fatalf("Invalid reactor response: %v", err)
	}
	if *args["key2"].(*string) != actual["key2"].(string) {
		t.Fatalf("Expected <" + *args["key2"].(*string) + ">; Actual <" + actual["key1"].(string) + ">")
	}
}

func DeleteReactor(t *testing.T, manageClient *basistheoryclient.Client, reactorId string) {
	e := manageClient.Reactors.Delete(
		context.TODO(),
		reactorId,
	)
	FailIfError(t, "Failed to delete reactor", e)
}

func CreateWebhook(t *testing.T, client *basistheoryclient.Client, url string) string {
	response, err := client.Webhooks.Create(
		context.TODO(),
		&basistheory.CreateWebhookRequest{
			Name:   "(Deletable) Webhook - " + uuid.NewString(),
			URL:    url,
			Events: []string{"token.created"},
		})
	FailIfError(t, "Could not create webhook", err)
	return response.ID
}

func UpdateWebhook(t *testing.T, client *basistheoryclient.Client, webhookId string, updateUrl string) {
	_, err := client.Webhooks.Update(
		context.TODO(),
		webhookId,
		&basistheory.UpdateWebhookRequest{
			Name:   "(Deletable) Updated -" + uuid.NewString(),
			URL:    updateUrl,
			Events: []string{"token.created", "token.updated"},
		})
	FailIfError(t, "Unable to update Webhook", err)
}

func DeleteWebhook(t *testing.T, client *basistheoryclient.Client, webhookId string) {
	err := client.Webhooks.Delete(
		context.TODO(),
		webhookId)
	FailIfError(t, "Unable to delete webhook", err)
}

func GetWebhookAssertUrl(t *testing.T, client *basistheoryclient.Client, webhookId string, url string) {
	response, err := client.Webhooks.Get(
		context.TODO(),
		webhookId)
	FailIfError(t, "Unable to get webhook", err)
	if url != response.URL {
		t.Errorf("Expected webhook URL to be %s. Got %s", url, response.URL)
	}
}

func FailIfError(t *testing.T, message string, err error) {
	if err != nil {
		t.Fatalf(message+": %v", err)
	}
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntPtr(i int) *int {
	return &i
}
