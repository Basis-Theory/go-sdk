package integration

import (
	"context"
	"errors"
	"os"
	"regexp"
	"testing"
	"time"

	basistheory "github.com/Basis-Theory/go-sdk/v5"
	"github.com/Basis-Theory/go-sdk/v5/agentic"
	"github.com/Basis-Theory/go-sdk/v5/agentic/agents"
	"github.com/Basis-Theory/go-sdk/v5/agentic/agents/instructions"
	"github.com/Basis-Theory/go-sdk/v5/agentic/enrollments"
	basistheoryclient "github.com/Basis-Theory/go-sdk/v5/client"
	"github.com/Basis-Theory/go-sdk/v5/option"
	"github.com/google/uuid"
)

func newAgenticClient() *basistheoryclient.Client {
	return basistheoryclient.NewClient(
		option.WithAPIKey(os.Getenv("BT_PVT_TEST_API_KEY")),
		option.WithBaseURL(os.Getenv("BT_API_URL")),
	)
}

func createCardToken(t *testing.T, client *basistheoryclient.Client, cardNumber string) string {
	token, err := client.Tokens.Create(
		context.TODO(),
		&basistheory.CreateTokenRequest{
			Type: StringPtr("card"),
			Data: map[string]interface{}{
				"number":           cardNumber,
				"expiration_month": 12,
				"expiration_year":  2030,
				"cvc":              "123",
			},
		},
	)
	FailIfError(t, "Failed to create card token", err)
	return *token.ID
}

func newDeviceContext() *basistheory.DeviceContext {
	return &basistheory.DeviceContext{
		ScreenHeight:      1080,
		ScreenWidth:       1920,
		UserAgentString:   "go-sdk-test",
		LanguageCode:      "en-US",
		TimeZone:          "America/New_York",
		JavaScriptEnabled: true,
		ClientDeviceID:    uuid.NewString(),
		ClientReferenceID: uuid.NewString(),
		PlatformType:      basistheory.DeviceContextPlatformTypeWeb,
	}
}

func createAndVerifyEnrollment(t *testing.T, client *basistheoryclient.Client, cardNumber string, email string, agentIDs []string) *basistheory.Enrollment {
	tokenID := createCardToken(t, client, cardNumber)

	req := &agentic.CreateEnrollmentRequest{
		TokenID:  tokenID,
		Consumer: &basistheory.Consumer{Email: email},
	}
	if len(agentIDs) > 0 {
		req.AgentIDs = agentIDs
	}

	enrollment, err := client.Agentic.Enrollments.Create(context.TODO(), req)
	FailIfError(t, "Failed to create enrollment", err)

	// Start verification
	verifyResponse, err := client.Agentic.Enrollments.Verify.Start(
		context.TODO(),
		*enrollment.ID,
		&basistheory.StartVerificationRequest{
			DeviceContext: newDeviceContext(),
		},
	)
	FailIfError(t, "Failed to start enrollment verification", err)

	// Auto-approved cards are already verified after verify.start
	if verifyResponse.Status == nil || string(*verifyResponse.Status) != "approved" {
		// OTP flow: select method, submit OTP, complete
		if len(verifyResponse.Methods) > 0 {
			_, err = client.Agentic.Enrollments.Verify.Method(
				context.TODO(),
				*enrollment.ID,
				&enrollments.SelectMethodRequest{
					MethodID: *verifyResponse.Methods[0].ID,
				},
			)
			FailIfError(t, "Failed to select verification method", err)
		}

		_, err = client.Agentic.Enrollments.Verify.Otp(
			context.TODO(),
			*enrollment.ID,
			&enrollments.SubmitOtpRequest{
				OtpCode: "123456",
			},
		)
		FailIfError(t, "Failed to submit OTP", err)

		_, err = client.Agentic.Enrollments.Verify.Complete(
			context.TODO(),
			*enrollment.ID,
			&enrollments.CompleteVerificationRequest{},
		)
		FailIfError(t, "Failed to complete verification", err)
	}

	result, err := client.Agentic.Enrollments.Get(context.TODO(), *enrollment.ID)
	FailIfError(t, "Failed to get enrollment after verification", err)
	return result
}

func findEnrollment(t *testing.T, client *basistheoryclient.Client, enrollmentID string) *basistheory.Enrollment {
	page, err := client.Agentic.Enrollments.List(
		context.TODO(),
		&agentic.EnrollmentsListRequest{},
	)
	FailIfError(t, "Failed to list enrollments", err)

	iter := page.Iterator()
	for iter.Next(context.TODO()) {
		item := iter.Current()
		if item.ID != nil && *item.ID == enrollmentID {
			return item
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("Error iterating enrollments: %v", err)
	}
	return nil
}

func findInstruction(t *testing.T, client *basistheoryclient.Client, agentID string, instructionID string, enrollmentID *string) *basistheory.Instruction {
	req := &agents.InstructionsListRequest{}
	if enrollmentID != nil {
		req.EnrollmentID = enrollmentID
	}

	page, err := client.Agentic.Agents.Instructions.List(
		context.TODO(),
		agentID,
		req,
	)
	FailIfError(t, "Failed to list instructions", err)

	iter := page.Iterator()
	for iter.Next(context.TODO()) {
		item := iter.Current()
		if item.ID != nil && *item.ID == instructionID {
			return item
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("Error iterating instructions: %v", err)
	}
	return nil
}

func TestAgentLifecycle(t *testing.T) {
	client := newAgenticClient()
	ctx := context.TODO()

	// Create
	agentName := "(Deletable) go-SDK-agent-" + uuid.NewString()
	agent, err := client.Agentic.Agents.Create(ctx, &agentic.CreateAgentRequest{
		Name: agentName,
	})
	FailIfError(t, "Failed to create agent", err)

	if agent.ID == nil {
		t.Fatal("Expected agent ID to be defined")
	}
	if agent.Name == nil || *agent.Name != agentName {
		t.Fatalf("Expected agent name to be %s, got %v", agentName, agent.Name)
	}
	if agent.Status == nil || *agent.Status != "active" {
		t.Fatalf("Expected agent status to be active, got %v", agent.Status)
	}
	if agent.EnrollmentIDs == nil {
		t.Fatal("Expected enrollment IDs to be defined")
	}
	if agent.CreatedAt == nil {
		t.Fatal("Expected created at to be defined")
	}

	// Get and verify all fields match
	retrieved, err := client.Agentic.Agents.Get(ctx, *agent.ID)
	FailIfError(t, "Failed to get agent", err)

	if *retrieved.ID != *agent.ID {
		t.Fatalf("Expected retrieved ID %s to match %s", *retrieved.ID, *agent.ID)
	}
	if *retrieved.Name != *agent.Name {
		t.Fatalf("Expected retrieved name %s to match %s", *retrieved.Name, *agent.Name)
	}
	if *retrieved.Status != "active" {
		t.Fatalf("Expected retrieved status to be active, got %s", *retrieved.Status)
	}
	if !retrieved.CreatedAt.Equal(*agent.CreatedAt) {
		t.Fatalf("Expected created at times to match")
	}

	// Update
	updatedName := "(Deletable) go-SDK-agent-updated-" + uuid.NewString()
	updated, err := client.Agentic.Agents.Update(ctx, *agent.ID, &agentic.UpdateAgentRequest{
		Name: StringPtr(updatedName),
	})
	FailIfError(t, "Failed to update agent", err)

	if *updated.ID != *agent.ID {
		t.Fatalf("Expected updated ID %s to match %s", *updated.ID, *agent.ID)
	}
	if *updated.Name != updatedName {
		t.Fatalf("Expected updated name to be %s, got %s", updatedName, *updated.Name)
	}
	if *updated.Status != "active" {
		t.Fatalf("Expected updated status to be active, got %s", *updated.Status)
	}
	if !updated.CreatedAt.Equal(*agent.CreatedAt) {
		t.Fatalf("Expected created at times to match after update")
	}

	// Delete
	err = client.Agentic.Agents.Delete(ctx, *agent.ID)
	FailIfError(t, "Failed to delete agent", err)

	// Verify deleted
	_, err = client.Agentic.Agents.Get(ctx, *agent.ID)
	if err == nil {
		t.Fatal("Expected error when getting deleted agent")
	}
	var notFoundError *basistheory.NotFoundError
	if !errors.As(err, &notFoundError) {
		t.Fatalf("Expected NotFoundError, got %T: %v", err, err)
	}
}

func TestEnrollmentAutoApprovedLifecycle(t *testing.T) {
	client := newAgenticClient()
	ctx := context.TODO()

	// Create and verify enrollment (passkey bypass card)
	enrollment := createAndVerifyEnrollment(t, client, "4000056655665556", "sdk-test@example.com", nil)

	if enrollment.ID == nil {
		t.Fatal("Expected enrollment ID to be defined")
	}
	if enrollment.Status == nil || string(*enrollment.Status) != "active" {
		t.Fatalf("Expected enrollment status to be active, got %v", enrollment.Status)
	}
	if enrollment.Provider == nil {
		t.Fatal("Expected provider to be defined")
	}
	if enrollment.CreatedAt == nil {
		t.Fatal("Expected created at to be defined")
	}

	// Verify card object fields
	if enrollment.Card == nil {
		t.Fatal("Expected card to be defined")
	}
	if enrollment.Card.Brand == nil || string(*enrollment.Card.Brand) != "visa" {
		t.Fatalf("Expected card brand to be visa, got %v", enrollment.Card.Brand)
	}
	if enrollment.Card.Bin == nil {
		t.Fatal("Expected card bin to be defined")
	}
	if enrollment.Card.Last4 == nil || *enrollment.Card.Last4 != "5556" {
		t.Fatalf("Expected card last4 to be 5556, got %v", enrollment.Card.Last4)
	}
	if enrollment.Card.ExpirationMonth == nil {
		t.Fatal("Expected card expiration month to be defined")
	}
	if enrollment.Card.ExpirationYear == nil {
		t.Fatal("Expected card expiration year to be defined")
	}

	// Get enrollment and verify all fields match
	retrieved, err := client.Agentic.Enrollments.Get(ctx, *enrollment.ID)
	FailIfError(t, "Failed to get enrollment", err)

	if *retrieved.ID != *enrollment.ID {
		t.Fatalf("Expected retrieved ID to match")
	}
	if string(*retrieved.Status) != "active" {
		t.Fatalf("Expected retrieved status to be active, got %s", string(*retrieved.Status))
	}
	if string(*retrieved.Provider) != string(*enrollment.Provider) {
		t.Fatalf("Expected provider to match")
	}
	if string(*retrieved.Card.Brand) != string(*enrollment.Card.Brand) {
		t.Fatalf("Expected card brand to match")
	}
	if *retrieved.Card.Bin != *enrollment.Card.Bin {
		t.Fatalf("Expected card bin to match")
	}
	if *retrieved.Card.Last4 != *enrollment.Card.Last4 {
		t.Fatalf("Expected card last4 to match")
	}
	if *retrieved.Card.ExpirationMonth != *enrollment.Card.ExpirationMonth {
		t.Fatalf("Expected card expiration month to match")
	}
	if *retrieved.Card.ExpirationYear != *enrollment.Card.ExpirationYear {
		t.Fatalf("Expected card expiration year to match")
	}
	if !retrieved.CreatedAt.Equal(*enrollment.CreatedAt) {
		t.Fatalf("Expected created at to match")
	}

	// List enrollments (paginated) and verify structure
	listed := findEnrollment(t, client, *enrollment.ID)
	if listed == nil {
		t.Fatal("Expected to find enrollment in list")
	}
	if string(*listed.Status) != "active" {
		t.Fatalf("Expected listed status to be active, got %s", string(*listed.Status))
	}
	if *listed.Card.Last4 != "5556" {
		t.Fatalf("Expected listed card last4 to be 5556, got %s", *listed.Card.Last4)
	}

	// Delete enrollment
	err = client.Agentic.Enrollments.Delete(ctx, *enrollment.ID)
	FailIfError(t, "Failed to delete enrollment", err)
}

func TestEnrollmentOtpVerificationFlow(t *testing.T) {
	client := newAgenticClient()
	ctx := context.TODO()

	// Create a card token (OTP challenge Visa test card)
	tokenID := createCardToken(t, client, "4000000000000002")

	// Create enrollment (will be pending_verification)
	enrollment, err := client.Agentic.Enrollments.Create(ctx, &agentic.CreateEnrollmentRequest{
		TokenID: tokenID,
		Consumer: &basistheory.Consumer{
			Email: "sdk-test-otp@example.com",
		},
	})
	FailIfError(t, "Failed to create enrollment", err)

	if enrollment.ID == nil {
		t.Fatal("Expected enrollment ID to be defined")
	}
	if enrollment.Status == nil || string(*enrollment.Status) != "pending_verification" {
		t.Fatalf("Expected enrollment status to be pending_verification, got %v", enrollment.Status)
	}
	if enrollment.Provider == nil {
		t.Fatal("Expected provider to be defined")
	}
	if enrollment.CreatedAt == nil {
		t.Fatal("Expected created at to be defined")
	}

	// Verify card object on initial create response
	if enrollment.Card == nil {
		t.Fatal("Expected card to be defined")
	}
	if enrollment.Card.Brand == nil || string(*enrollment.Card.Brand) != "visa" {
		t.Fatalf("Expected card brand to be visa, got %v", enrollment.Card.Brand)
	}
	if enrollment.Card.Bin == nil {
		t.Fatal("Expected card bin to be defined")
	}
	if enrollment.Card.Last4 == nil || *enrollment.Card.Last4 != "0002" {
		t.Fatalf("Expected card last4 to be 0002, got %v", enrollment.Card.Last4)
	}
	if enrollment.Card.ExpirationMonth == nil {
		t.Fatal("Expected card expiration month to be defined")
	}
	if enrollment.Card.ExpirationYear == nil {
		t.Fatal("Expected card expiration year to be defined")
	}

	// Start verification — expect challenge with OTP methods
	verifyResponse, err := client.Agentic.Enrollments.Verify.Start(
		ctx,
		*enrollment.ID,
		&basistheory.StartVerificationRequest{
			DeviceContext: newDeviceContext(),
		},
	)
	FailIfError(t, "Failed to start verification", err)

	if verifyResponse.Status == nil || string(*verifyResponse.Status) != "challenge" {
		t.Fatalf("Expected verify status to be challenge, got %v", verifyResponse.Status)
	}
	if len(verifyResponse.Methods) == 0 {
		t.Fatal("Expected verification methods to be present")
	}
	if verifyResponse.Methods[0].ID == nil {
		t.Fatal("Expected method ID to be defined")
	}
	if verifyResponse.Methods[0].Type == nil {
		t.Fatal("Expected method type to be defined")
	}

	// Select verification method
	if len(verifyResponse.Methods) > 0 {
		_, err = client.Agentic.Enrollments.Verify.Method(
			ctx,
			*enrollment.ID,
			&enrollments.SelectMethodRequest{
				MethodID: *verifyResponse.Methods[0].ID,
			},
		)
		FailIfError(t, "Failed to select verification method", err)
	}

	// Submit OTP (mock accepts any code)
	otpResponse, err := client.Agentic.Enrollments.Verify.Otp(
		ctx,
		*enrollment.ID,
		&enrollments.SubmitOtpRequest{
			OtpCode: "123456",
		},
	)
	FailIfError(t, "Failed to submit OTP", err)

	if otpResponse.Status == nil || string(*otpResponse.Status) != "device_bound" {
		t.Fatalf("Expected OTP response status to be device_bound, got %v", otpResponse.Status)
	}

	// Complete verification
	completeResponse, err := client.Agentic.Enrollments.Verify.Complete(
		ctx,
		*enrollment.ID,
		&enrollments.CompleteVerificationRequest{},
	)
	FailIfError(t, "Failed to complete verification", err)

	if completeResponse.Status == nil || string(*completeResponse.Status) != "verified" {
		t.Fatalf("Expected complete response status to be verified, got %v", completeResponse.Status)
	}

	// Verify enrollment is now active with all fields
	completed, err := client.Agentic.Enrollments.Get(ctx, *enrollment.ID)
	FailIfError(t, "Failed to get completed enrollment", err)

	if *completed.ID != *enrollment.ID {
		t.Fatalf("Expected completed ID to match")
	}
	if string(*completed.Status) != "active" {
		t.Fatalf("Expected completed status to be active, got %s", string(*completed.Status))
	}
	if string(*completed.Provider) != string(*enrollment.Provider) {
		t.Fatalf("Expected provider to match")
	}
	if string(*completed.Card.Brand) != "visa" {
		t.Fatalf("Expected card brand to be visa")
	}
	if *completed.Card.Last4 != "0002" {
		t.Fatalf("Expected card last4 to be 0002, got %s", *completed.Card.Last4)
	}
	if !completed.CreatedAt.Equal(*enrollment.CreatedAt) {
		t.Fatalf("Expected created at to match")
	}

	// Cleanup
	err = client.Agentic.Enrollments.Delete(ctx, *enrollment.ID)
	FailIfError(t, "Failed to delete enrollment", err)
}

func TestInstructionLifecycleWithCredentials(t *testing.T) {
	client := newAgenticClient()
	ctx := context.TODO()

	// Setup: create agent and auto-approved enrollment
	agent, err := client.Agentic.Agents.Create(ctx, &agentic.CreateAgentRequest{
		Name: "(Deletable) go-SDK-instruction-agent-" + uuid.NewString(),
	})
	FailIfError(t, "Failed to create agent", err)

	enrollment := createAndVerifyEnrollment(t, client, "4000056655665556", "sdk-test-instructions@example.com", []string{*agent.ID})
	if string(*enrollment.Status) != "active" {
		t.Fatalf("Expected enrollment status to be active, got %s", string(*enrollment.Status))
	}

	// Create instruction
	expiresAt := time.Now().UTC().AddDate(0, 0, 7)

	instruction, err := client.Agentic.Agents.Instructions.Create(ctx, *agent.ID, &agents.CreateInstructionRequest{
		EnrollmentID: *enrollment.ID,
		Amount: &basistheory.Amount{
			Value:    "25.00",
			Currency: StringPtr("USD"),
		},
		Description: "(Deletable) go-SDK test purchase",
		ExpiresAt:   expiresAt,
	})
	FailIfError(t, "Failed to create instruction", err)

	if instruction.ID == nil {
		t.Fatal("Expected instruction ID to be defined")
	}
	if instruction.EnrollmentID == nil || *instruction.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected enrollment ID to match")
	}
	if instruction.Status == nil || string(*instruction.Status) != "pending_verification" {
		t.Fatalf("Expected instruction status to be pending_verification, got %v", instruction.Status)
	}
	if instruction.Amount == nil {
		t.Fatal("Expected amount to be defined")
	}
	if instruction.Amount.Value != "25.00" {
		t.Fatalf("Expected amount value to be 25.00, got %s", instruction.Amount.Value)
	}
	if instruction.Amount.Currency == nil || *instruction.Amount.Currency != "USD" {
		t.Fatalf("Expected amount currency to be USD, got %v", instruction.Amount.Currency)
	}
	if instruction.ExpiresAt == nil {
		t.Fatal("Expected expires at to be defined")
	}
	if instruction.CreatedAt == nil {
		t.Fatal("Expected created at to be defined")
	}

	// Get instruction and verify all fields match
	retrieved, err := client.Agentic.Agents.Instructions.Get(ctx, *agent.ID, *instruction.ID)
	FailIfError(t, "Failed to get instruction", err)

	if *retrieved.ID != *instruction.ID {
		t.Fatalf("Expected retrieved ID to match")
	}
	if *retrieved.EnrollmentID != *instruction.EnrollmentID {
		t.Fatalf("Expected retrieved enrollment ID to match")
	}
	if string(*retrieved.Status) != string(*instruction.Status) {
		t.Fatalf("Expected retrieved status to match")
	}
	if retrieved.Amount.Value != instruction.Amount.Value {
		t.Fatalf("Expected retrieved amount value to match")
	}
	if *retrieved.Amount.Currency != *instruction.Amount.Currency {
		t.Fatalf("Expected retrieved amount currency to match")
	}
	if !retrieved.CreatedAt.Equal(*instruction.CreatedAt) {
		t.Fatalf("Expected retrieved created at to match")
	}

	// Update instruction
	updated, err := client.Agentic.Agents.Instructions.Update(ctx, *agent.ID, *instruction.ID, &agents.UpdateInstructionRequest{
		Amount: &basistheory.Amount{
			Value:    "50.00",
			Currency: StringPtr("USD"),
		},
	})
	FailIfError(t, "Failed to update instruction", err)

	if *updated.ID != *instruction.ID {
		t.Fatalf("Expected updated ID to match")
	}
	if updated.Amount.Value != "50.00" {
		t.Fatalf("Expected updated amount value to be 50.00, got %s", updated.Amount.Value)
	}
	if *updated.Amount.Currency != "USD" {
		t.Fatalf("Expected updated amount currency to be USD")
	}
	if *updated.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected updated enrollment ID to match")
	}
	if !updated.CreatedAt.Equal(*instruction.CreatedAt) {
		t.Fatalf("Expected updated created at to match")
	}

	// List instructions (paginated) and verify structure
	listed := findInstruction(t, client, *agent.ID, *instruction.ID, nil)
	if listed == nil {
		t.Fatal("Expected to find instruction in list")
	}
	if *listed.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected listed enrollment ID to match")
	}
	if listed.Amount.Value != "50.00" {
		t.Fatalf("Expected listed amount value to be 50.00, got %s", listed.Amount.Value)
	}

	// Verify instruction (required before credentials can be retrieved)
	instrVerifyResponse, err := client.Agentic.Agents.Instructions.Verify.Start(
		ctx,
		*agent.ID,
		*instruction.ID,
		&basistheory.StartVerificationRequest{
			DeviceContext: newDeviceContext(),
		},
	)
	FailIfError(t, "Failed to start instruction verification", err)

	// Passkey bypass card returns { status: "verified" } immediately
	if instrVerifyResponse.Status == nil || string(*instrVerifyResponse.Status) != "verified" {
		t.Fatalf("Expected instruction verify status to be verified, got %v", instrVerifyResponse.Status)
	}

	// Confirm instruction is now approved
	approved, err := client.Agentic.Agents.Instructions.Get(ctx, *agent.ID, *instruction.ID)
	FailIfError(t, "Failed to get approved instruction", err)

	if string(*approved.Status) != "approved" {
		t.Fatalf("Expected instruction status to be approved, got %s", string(*approved.Status))
	}
	if *approved.ID != *instruction.ID {
		t.Fatalf("Expected approved ID to match")
	}
	if *approved.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected approved enrollment ID to match")
	}

	// Get credentials
	credentials, err := client.Agentic.Agents.Instructions.Credentials.Create(
		ctx,
		*agent.ID,
		*instruction.ID,
		&instructions.GetCredentialsRequest{
			Merchant: &basistheory.AgenticMerchant{
				Name:        "(Deletable) Test Merchant",
				URL:         "https://example.com",
				CountryCode: "US",
			},
		},
	)
	FailIfError(t, "Failed to get credentials", err)

	if credentials.Card == nil {
		t.Fatal("Expected credentials card to be defined")
	}
	if credentials.Card.Number == nil {
		t.Fatal("Expected card number to be defined")
	}
	if credentials.Card.ExpirationMonth == nil {
		t.Fatal("Expected card expiration month to be defined")
	}
	if credentials.Card.ExpirationYear == nil {
		t.Fatal("Expected card expiration year to be defined")
	}
	if credentials.Card.Cvc == nil {
		t.Fatal("Expected card CVC to be defined")
	}

	// Verify mock virtual card number format: 400000100000 + last 4
	matched, _ := regexp.MatchString(`^400000100000\d{4}$`, *credentials.Card.Number)
	if !matched {
		t.Fatalf("Expected card number to match pattern 400000100000XXXX, got %s", *credentials.Card.Number)
	}

	// Delete instruction
	err = client.Agentic.Agents.Instructions.Delete(ctx, *agent.ID, *instruction.ID)
	FailIfError(t, "Failed to delete instruction", err)

	// Cleanup
	err = client.Agentic.Enrollments.Delete(ctx, *enrollment.ID)
	FailIfError(t, "Failed to delete enrollment", err)

	err = client.Agentic.Agents.Delete(ctx, *agent.ID)
	FailIfError(t, "Failed to delete agent", err)
}

func TestInstructionListFilteredByEnrollment(t *testing.T) {
	client := newAgenticClient()
	ctx := context.TODO()

	// Setup
	agent, err := client.Agentic.Agents.Create(ctx, &agentic.CreateAgentRequest{
		Name: "(Deletable) go-SDK-filter-agent-" + uuid.NewString(),
	})
	FailIfError(t, "Failed to create agent", err)

	enrollment := createAndVerifyEnrollment(t, client, "4000056655665556", "sdk-test-filter@example.com", []string{*agent.ID})

	expiresAt := time.Now().UTC().AddDate(0, 0, 7)

	instruction, err := client.Agentic.Agents.Instructions.Create(ctx, *agent.ID, &agents.CreateInstructionRequest{
		EnrollmentID: *enrollment.ID,
		Amount: &basistheory.Amount{
			Value:    "10.00",
			Currency: StringPtr("USD"),
		},
		Description: "(Deletable) go-SDK filter test",
		ExpiresAt:   expiresAt,
	})
	FailIfError(t, "Failed to create instruction", err)

	// Verify created instruction fields
	if instruction.ID == nil {
		t.Fatal("Expected instruction ID to be defined")
	}
	if *instruction.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected enrollment ID to match")
	}
	if string(*instruction.Status) != "pending_verification" {
		t.Fatalf("Expected status to be pending_verification, got %s", string(*instruction.Status))
	}
	if instruction.Amount.Value != "10.00" {
		t.Fatalf("Expected amount value to be 10.00, got %s", instruction.Amount.Value)
	}
	if *instruction.Amount.Currency != "USD" {
		t.Fatalf("Expected amount currency to be USD")
	}
	if instruction.ExpiresAt == nil {
		t.Fatal("Expected expires at to be defined")
	}
	if instruction.CreatedAt == nil {
		t.Fatal("Expected created at to be defined")
	}

	// List with enrollment filter (paginated)
	listed := findInstruction(t, client, *agent.ID, *instruction.ID, enrollment.ID)
	if listed == nil {
		t.Fatal("Expected to find instruction in filtered list")
	}
	if *listed.EnrollmentID != *enrollment.ID {
		t.Fatalf("Expected listed enrollment ID to match")
	}
	if listed.Amount.Value != "10.00" {
		t.Fatalf("Expected listed amount value to be 10.00, got %s", listed.Amount.Value)
	}

	// Cleanup
	err = client.Agentic.Agents.Instructions.Delete(ctx, *agent.ID, *instruction.ID)
	FailIfError(t, "Failed to delete instruction", err)

	err = client.Agentic.Enrollments.Delete(ctx, *enrollment.ID)
	FailIfError(t, "Failed to delete enrollment", err)

	err = client.Agentic.Agents.Delete(ctx, *agent.ID)
	FailIfError(t, "Failed to delete agent", err)
}
