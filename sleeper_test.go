package sleeper

import (
	"net/http"
	"time"
)

var (
	hClient http.Client = http.Client{Timeout: time.Duration(30) * time.Second}
	client  *Client
	err     error
)
