package sleeper

import "fmt"

// Matchup is a single matchup from the league matchups API.
type Matchup struct {
	StartersPoints []float32          `json:"starters_points"`
	Starters       []string           `json:"starters"`
	RosterID       int                `json:"roster_id"`
	Points         float32            `json:"points"`
	PlayerPoints   map[string]float32 `json:"players_points"`
	Players        []string           `json:"players"`
	MatchupID      int                `json:"matchup_id"`
	CustomPoints   float32            `json:"custom_points"`
}

/*
This endpoint retrieves all users in a league.

leagueID (required) : The ID of the league to retrieve
week 	 (required) : The week number to get the matchups
*/
func (c *Client) GetLeagueMatchups(leagueID string, week int8) ([]Matchup, error) {
	// https://api.sleeper.app/v1/league/<league_id>/matchups/<week>
	reqURL := fmt.Sprintf("%s/league/%s/matchups/%d", c.sleeperURL, leagueID, week)

	matchups := new([]Matchup)

	err := c.get(reqURL, matchups)

	if err != nil {
		return *matchups, err
	}

	return *matchups, nil
}
