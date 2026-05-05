// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thern1_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/thern.1-go"
	"github.com/stainless-sdks/thern.1-go/internal/testutil"
	"github.com/stainless-sdks/thern.1-go/option"
	"github.com/stainless-sdks/thern.1-go/shared"
)

func TestStoreOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Store.Orders.New(context.TODO(), thern1.StoreOrderNewParams{
		Order: shared.OrderParam{
			ID:       thern1.Int(10),
			Complete: thern1.Bool(true),
			PetID:    thern1.Int(198772),
			Quantity: thern1.Int(7),
			ShipDate: thern1.Time(time.Now()),
			Status:   shared.OrderStatusApproved,
		},
	})
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderGet(t *testing.T) {
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
	_, err := client.Store.Orders.Get(context.TODO(), 0)
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderDelete(t *testing.T) {
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
	err := client.Store.Orders.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
