// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thern1_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/PhoenixVazquez-glitch/Aether-OS"
	"github.com/PhoenixVazquez-glitch/Aether-OS/internal/testutil"
	"github.com/PhoenixVazquez-glitch/Aether-OS/option"
)

func TestStoreListInventory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := thern1.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.ListInventory(context.TODO())
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
