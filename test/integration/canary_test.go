package integration

import (
	"context"
	basistheory "github.com/Basis-Theory/go-sdk"
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

	// Create Token
	createCardNumber := "6011000990139424"
	tokenId := CreateToken(t, client, createCardNumber)
	GetAndValidateCardNumber(t, client, tokenId, createCardNumber)

	// Update Token
	//updateCardNumber := "4242424242424242"
	//UpdateToken(t, client, tokenId, updateCardNumber)
	//GetAndValidateCardNumber(t, client, tokenId, updateCardNumber)

	// Delete Token
	DeleteToken(t, client, tokenId)

	_, err := client.Tokens.Get(
		context.TODO(),
		tokenId)
	if err == nil {
		t.Errorf("Expected error when trying to get a token that doesn't exist")
	}
	if _, ok := err.(basistheory.NotFoundError); ok {
		t.Errorf("Expected error to be Not Found")
	}
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
			Privacy:               nil,
			Metadata:              nil,
			SearchIndexes:         nil,
			FingerprintExpression: nil,
			Mask:                  nil,
			ExpiresAt:             nil,
			DeduplicateToken:      nil,
			Containers:            nil,
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
