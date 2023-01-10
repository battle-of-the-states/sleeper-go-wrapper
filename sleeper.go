package sleeper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is a client for working with the sleeper.app read only Web API
type Client struct {
	httpClient *http.Client

	sleeperURL string

	sleeperCDN string

	NFLPlayers AllPlayersJSON
}

// NewClient creates a new Sleeper Client.
func NewClient(httpClient *http.Client) (*Client, error) {
	c := &Client{
		httpClient: httpClient,
		sleeperURL: SLEEPER_BASE_URL,
		sleeperCDN: SLEEPER_BASE_CDN,
	}

	return c, nil
}

// Error represents an error returned by the sleeper.app Web API.
type Error struct {
	// A short description of the error.
	Message string `json:"message"`
	// The HTTP status code.
	Err int `json:"error"`
}

// Error ...
func (e Error) Error() string {
	return e.Message
}

// decodeError decodes an Error from response status code based off
// the docs in sleeper.app -> https://docs.sleeper.app/#errors
func (c *Client) decodeError(statusCode int) error {
	e := Error{Err: statusCode}

	switch e.Err {
	case 400:
		e.Message = "Bad Request -- Your request is invalid."
	case 404:
		e.Message = "Not Found -- The specified kitten could not be found."
	case 429:
		e.Message = "Too Many Requests -- You're requesting too many kittens! Slow down!"
	case 500:
		e.Message = "Internal Server Error -- We had a problem with our server. Try again later."
	case 503:
		e.Message = "Service Unavailable -- We're temporarily offline for maintenance. Please try again later."
	default:
		e.Message = fmt.Sprintf("lastfm: unexpected HTTP %d: %s (empty error)",
			statusCode, http.StatusText(statusCode))
	}

	return e
}

// get handles the get requests for the client
func (c *Client) get(url string, result interface{}) error {
	for {
		resp, err := c.httpClient.Get(url)

		if err != nil {
			return err
		}

		// body, err := ioutil.ReadAll(resp.Body)

		// if err != nil {
		// 	return err
		// }

		// fmt.Println(string(body))

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			return nil
		}
		if resp.StatusCode != http.StatusOK {
			return c.decodeError(resp.StatusCode)
		}

		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}

		break
	}

	return nil
}
