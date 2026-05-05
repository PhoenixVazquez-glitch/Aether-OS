// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thern1_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/thern.1-go"
	"github.com/stainless-sdks/thern.1-go/internal/testutil"
	"github.com/stainless-sdks/thern.1-go/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.New(context.TODO(), thern1.UserNewParams{
		User: thern1.UserParam{
			ID:         thern1.Int(10),
			Email:      thern1.String("john@email.com"),
			FirstName:  thern1.String("John"),
			LastName:   thern1.String("James"),
			Password:   thern1.String("12345"),
			Phone:      thern1.String("12345"),
			Username:   thern1.String("theUser"),
			UserStatus: thern1.Int(1),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "username")
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Users.Update(
		context.TODO(),
		"username",
		thern1.UserUpdateParams{
			User: thern1.UserParam{
				ID:         thern1.Int(10),
				Email:      thern1.String("john@email.com"),
				FirstName:  thern1.String("John"),
				LastName:   thern1.String("James"),
				Password:   thern1.String("12345"),
				Phone:      thern1.String("12345"),
				Username:   thern1.String("theUser"),
				UserStatus: thern1.Int(1),
			},
		},
	)
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
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
	err := client.Users.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.NewWithList(context.TODO(), thern1.UserNewWithListParams{
		Items: []thern1.UserParam{{
			ID:         thern1.Int(10),
			Email:      thern1.String("john@email.com"),
			FirstName:  thern1.String("John"),
			LastName:   thern1.String("James"),
			Password:   thern1.String("12345"),
			Phone:      thern1.String("12345"),
			Username:   thern1.String("theUser"),
			UserStatus: thern1.Int(1),
		}},
	})
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), thern1.UserLoginParams{
		Password: thern1.String("password"),
		Username: thern1.String("username"),
	})
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.Users.Logout(context.TODO())
	if err != nil {
		var apierr *thern1.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
