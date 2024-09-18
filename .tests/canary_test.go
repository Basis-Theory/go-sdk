package _tests

import (
	"context"
	"os"
	"testing"
)
import (
	basistheoryclient "github.com/basis-theory/go-sdk/client"
	"github.com/basis-theory/go-sdk/option"
)

func TestMyFunction(t *testing.T) {
	client := basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_MGT_API_KEY")),
		option.WithBaseURL("https://api.flock-dev.com"),
	)

	response, _ := client.Tenants.Self.Get(
		context.TODO(),
	)

	if response.Name == nil || *response.Name != `SDK Integration Tests` {
		t.Errorf("Expected SDK Integration Tests. Got %s", *response.Name)
	}
}
