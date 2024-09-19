package integration

import (
	"context"
	"errors"
	basistheory "github.com/Basis-Theory/go-sdk"
	"github.com/google/uuid"
	"os"
	"testing"
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

	// Update Token (broken due to invalid `Content-Type`) https://basistheory.slack.com/archives/C07BZHM9JUD/p1726766102647879
	//updateCardNumber := "4242424242424242"
	//UpdateToken(t, client, tokenId, updateCardNumber)
	//GetAndValidateCardNumber(t, client, tokenId, updateCardNumber)

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

func CreateProxy(t *testing.T, manageClient *basistheoryclient.Client, appliationId string) string {
	response, err := manageClient.Proxies.Create(
		context.TODO(),
		&basistheory.CreateProxyRequest{
			Name:              "Go-" + uuid.NewString(),
			DestinationURL:    "https://echo.basistheory.com",
			RequestReactorID:  nil,
			ResponseReactorID: nil,
			RequestTransform:  nil,
			ResponseTransform: nil,
			Application: &basistheory.Application{
				ID: StringPtr(appliationId),
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

func NewPrivateClient() *basistheoryclient.Client {
	client := basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_PVT_API_KEY")),
		option.WithBaseURL(os.Getenv("BT_API_URL")),
	)
	return client
}

func CreateToken(t *testing.T, client *basistheoryclient.Client, cardNumber string) string {
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
			Name:        "Go-SDK-" + uuid.NewString(),
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
			Name: "Go-SDK-" + uuid.NewString(),
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
