// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contextdev_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/internal/testutil"
	"github.com/context-dot-dev/context-go-sdk/option"
)

func TestMonitorNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.New(context.TODO(), contextdev.MonitorNewParams{
		OfMonitorsCreatePageExactMonitorRequest: &contextdev.MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequest{
			ChangeDetection: contextdev.MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestChangeDetection{
				Type: "exact",
			},
			Name: "Acme pricing page",
			Schedule: contextdev.MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestSchedule{
				Frequency: 6,
				Type:      "interval",
				Unit:      "hours",
			},
			Target: contextdev.MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestTarget{
				Type:                "page",
				URL:                 "https://acme.com/pricing",
				NormalizeWhitespace: contextdev.Bool(true),
			},
			Tags: []string{"pricing", "competitor"},
			Webhook: contextdev.MonitorNewParamsBodyMonitorsCreatePageExactMonitorRequestWebhook{
				URL: "https://example.com/webhook",
			},
		},
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.Get(context.TODO(), "mon_123")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.Update(
		context.TODO(),
		"mon_123",
		contextdev.MonitorUpdateParams{
			ChangeDetection: contextdev.MonitorUpdateParamsChangeDetectionUnion{
				OfExact: &contextdev.MonitorUpdateParamsChangeDetectionExact{},
			},
			Name: contextdev.String("Acme pricing monitor"),
			Schedule: contextdev.MonitorUpdateParamsSchedule{
				Frequency: 1,
				Type:      "interval",
				Unit:      "hours",
			},
			Status: contextdev.MonitorUpdateParamsStatusActive,
			Tags:   []string{"pricing", "competitor"},
			Target: contextdev.MonitorUpdateParamsTargetUnion{
				OfPage: &contextdev.MonitorUpdateParamsTargetPage{
					URL:                 "https://acme.com/pricing",
					NormalizeWhitespace: contextdev.Bool(true),
				},
			},
			Webhook: contextdev.MonitorUpdateParamsWebhook{
				URL: "https://example.com/webhook",
			},
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.List(context.TODO(), contextdev.MonitorListParams{
		ChangeDetectionType: contextdev.MonitorListParamsChangeDetectionTypeExact,
		Cursor:              contextdev.String("cursor"),
		Limit:               contextdev.Int(1),
		Status:              contextdev.MonitorListParamsStatusActive,
		Tag:                 contextdev.String("tag"),
		TargetType:          contextdev.MonitorListParamsTargetTypePage,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.Delete(context.TODO(), "mon_123")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorListAccountChangesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.ListAccountChanges(context.TODO(), contextdev.MonitorListAccountChangesParams{
		ChangeDetectionType: contextdev.MonitorListAccountChangesParamsChangeDetectionTypeExact,
		Cursor:              contextdev.String("cursor"),
		Limit:               contextdev.Int(1),
		MonitorID:           contextdev.String("monitor_id"),
		Since:               contextdev.Time(time.Now()),
		Tag:                 contextdev.String("tag"),
		TargetType:          contextdev.MonitorListAccountChangesParamsTargetTypePage,
		Until:               contextdev.Time(time.Now()),
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorListAccountRunsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.ListAccountRuns(context.TODO(), contextdev.MonitorListAccountRunsParams{
		Cursor: contextdev.String("cursor"),
		Limit:  contextdev.Int(1),
		Status: contextdev.MonitorListAccountRunsParamsStatusQueued,
	})
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorListChangesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.ListChanges(
		context.TODO(),
		"mon_123",
		contextdev.MonitorListChangesParams{
			Cursor: contextdev.String("cursor"),
			Limit:  contextdev.Int(1),
			Since:  contextdev.Time(time.Now()),
			Tag:    contextdev.String("tag"),
			Until:  contextdev.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorListRunsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.ListRuns(
		context.TODO(),
		"mon_123",
		contextdev.MonitorListRunsParams{
			Cursor: contextdev.String("cursor"),
			Limit:  contextdev.Int(1),
			Status: contextdev.MonitorListRunsParamsStatusQueued,
		},
	)
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorGetChange(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.GetChange(context.TODO(), "chg_123")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := contextdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitors.Run(context.TODO(), "mon_123")
	if err != nil {
		var apierr *contextdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
