package sleeper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SLEEPER_BASE_URL = "https://api.sleeper.app/v1"
	SLEEPER_BASE_CDN = "https://sleepercdn.com/"
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

// Sport ensures users give valid options for sports
type Sport string

// TODO: Not sure what all sports are allowed
const (
	// NFL is the nfl sport tag
	NFL Sport = "nfl"
	// NBA is the nba sport tag
	NBA Sport = "nba"
	// LCS is the lcs sport tag
	LCS Sport = "lcs"
)

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
func (c *Client) decodeError(resp *http.Response) error {
	e := Error{Err: resp.StatusCode}

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
			resp.StatusCode, http.StatusText(resp.StatusCode))
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
			return c.decodeError(resp)
		}

		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}

		break
	}

	return nil
}