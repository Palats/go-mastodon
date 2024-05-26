package mastodon

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// Marker describes a position in a timeline.
type Marker struct {
	LastReadID ID        `json:"last_read_id"`
	Version    int64     `json:"version"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// GetMarkers get saved timelines position.
func (c *Client) GetMarkers(ctx context.Context, timelines []string) (map[string]*Marker, error) {
	markers := map[string]*Marker{}
	values := url.Values{}
	for _, t := range timelines {
		values.Add("timeline[]", t)
	}
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/markers", values, &markers, nil)
	if err != nil {
		return nil, err
	}
	return markers, nil
}
