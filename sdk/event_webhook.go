package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EventWebhook is a Sendgrid event webhook settings.
type EventWebhook struct { //nolint:maligned
	Enabled           bool   `json:"enabled"`
	ID                string `json:"id"`
	FriendlyName      string `json:"friendly_name,omitempty"`
	URL               string `json:"url,omitempty"`
	GroupResubscribe  bool   `json:"group_resubscribe"` //nolint:tagliatelle
	Delivered         bool   `json:"delivered"`
	GroupUnsubscribe  bool   `json:"group_unsubscribe"` //nolint:tagliatelle
	SpamReport        bool   `json:"spam_report"`       //nolint:tagliatelle
	Bounce            bool   `json:"bounce"`
	Deferred          bool   `json:"deferred"`
	Unsubscribe       bool   `json:"unsubscribe"`
	Processed         bool   `json:"processed"`
	Open              bool   `json:"open"`
	Click             bool   `json:"click"`
	Dropped           bool   `json:"dropped"`
	OAuthClientID     string `json:"oauth_client_id,omitempty"`     //nolint:tagliatelle
	OAuthClientSecret string `json:"oauth_client_secret,omitempty"` //nolint:tagliatelle
	OAuthTokenURL     string `json:"oauth_token_url,omitempty"`     //nolint:tagliatelle
}

type EventWebhookSigning struct {
	Enabled   bool   `json:"enabled"`
	PublicKey string `json:"public_key"` //nolint:tagliatelle
}

func parseEventWebhook(respBody string) (*EventWebhook, RequestError) {
	var body EventWebhook
	if err := json.Unmarshal([]byte(respBody), &body); err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed parsing event webhook: %w", err),
		}
	}

	return &body, RequestError{StatusCode: http.StatusOK, Err: nil}
}

func parseEventWebhookSigning(respBody string) (*EventWebhookSigning, RequestError) {
	var body EventWebhookSigning
	if err := json.Unmarshal([]byte(respBody), &body); err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed parsing event webhook: %w", err),
		}
	}

	return &body, RequestError{StatusCode: http.StatusOK, Err: nil}
}

func (c *Client) DeleteEventWebhook(ctx context.Context, id string) RequestError {
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)
	respBody, statusCode, err := c.Get(ctx, "DELETE", path)

	if err != nil {
		return RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed deleting event webhook %s: %w", id, err),
		}
	}

	if statusCode >= http.StatusMultipleChoices {
		return RequestError{
			StatusCode: statusCode,
			Err:        fmt.Errorf("%w, status: %d, response: %s", ErrFailedDeletingEventWebhook, statusCode, respBody),
		}
	}

	return RequestError{StatusCode: http.StatusOK, Err: nil}
}

func (c *Client) CreateEventWebhook(ctx context.Context, enabled bool, friendlyName string, url string, groupResubscribe bool, delivered bool, groupUnsubscribe bool, spamReport bool, bounce bool, deferred bool, unsubscribe bool, processed bool, open bool, click bool, dropped bool, oauthClientID string, oauthClientSecret string, oauthTokenURL string) (*EventWebhook, RequestError) {
	if url == "" {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        ErrURLRequired,
		}
	}

	respBody, statusCode, err := c.Post(ctx, "POST", "/user/webhooks/event/settings", EventWebhook{
		Enabled:           enabled,
		FriendlyName:      friendlyName,
		URL:               url,
		GroupResubscribe:  groupResubscribe,
		Delivered:         delivered,
		GroupUnsubscribe:  groupUnsubscribe,
		SpamReport:        spamReport,
		Bounce:            bounce,
		Deferred:          deferred,
		Unsubscribe:       unsubscribe,
		Processed:         processed,
		Open:              open,
		Click:             click,
		Dropped:           dropped,
		OAuthClientID:     oauthClientID,
		OAuthClientSecret: oauthClientSecret,
		OAuthTokenURL:     oauthTokenURL,
	})
	if err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed creating event webhook: %w", err),
		}
	}

	if statusCode >= http.StatusMultipleChoices {
		return nil, RequestError{
			StatusCode: statusCode,
			Err:        fmt.Errorf("%w, status: %d, response: %s", ErrFailedCreatingEventWebhook, statusCode, respBody),
		}
	}

	return parseEventWebhook(respBody)
}

func (c *Client) UpdateEventWebhook(ctx context.Context, id string, enabled bool, friendlyName string, url string, groupResubscribe bool, delivered bool, groupUnsubscribe bool, spamReport bool, bounce bool, deferred bool, unsubscribe bool, processed bool, open bool, click bool, dropped bool, oauthClientID string, oauthClientSecret string, oauthTokenURL string) (*EventWebhook, RequestError) {
	if url == "" {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        ErrURLRequired,
		}
	}
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)

	respBody, statusCode, err := c.Post(ctx, "PATCH", path, EventWebhook{
		Enabled:           enabled,
		FriendlyName:      friendlyName,
		URL:               url,
		GroupResubscribe:  groupResubscribe,
		Delivered:         delivered,
		GroupUnsubscribe:  groupUnsubscribe,
		SpamReport:        spamReport,
		Bounce:            bounce,
		Deferred:          deferred,
		Unsubscribe:       unsubscribe,
		Processed:         processed,
		Open:              open,
		Click:             click,
		Dropped:           dropped,
		OAuthClientID:     oauthClientID,
		OAuthClientSecret: oauthClientSecret,
		OAuthTokenURL:     oauthTokenURL,
	})
	if err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed patching event webhook: %w", err),
		}
	}

	if statusCode >= http.StatusMultipleChoices {
		return nil, RequestError{
			StatusCode: statusCode,
			Err:        fmt.Errorf("%w, status: %d, response: %s", ErrFailedUpdatingEventWebhook, statusCode, respBody),
		}
	}

	return parseEventWebhook(respBody)
}

// ReadEventWebhook retrieves an EventWebhook and returns it.
func (c *Client) ReadEventWebhook(ctx context.Context, id string) (*EventWebhook, RequestError) {
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)
	respBody, _, err := c.Get(ctx, "GET", path)
	if err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return parseEventWebhook(respBody)
}

func (c *Client) ConfigureEventWebhookSigning(ctx context.Context, enabled bool) (*EventWebhookSigning, RequestError) {
	respBody, statusCode, err := c.Post(ctx, "PATCH", "/user/webhooks/event/settings/signed", EventWebhookSigning{
		Enabled: enabled,
	})
	if err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("failed creating event webhook: %w", err),
		}
	}

	if statusCode >= http.StatusMultipleChoices {
		return nil, RequestError{
			StatusCode: statusCode,
			Err:        fmt.Errorf("%w, status: %d, response: %s", ErrFailedCreatingEventWebhook, statusCode, respBody),
		}
	}

	return parseEventWebhookSigning(respBody)
}

func (c *Client) ReadEventWebhookSigning(ctx context.Context) (*EventWebhookSigning, RequestError) {
	respBody, _, err := c.Get(ctx, "GET", "/user/webhooks/event/settings/signed")
	if err != nil {
		return nil, RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return parseEventWebhookSigning(respBody)
}
