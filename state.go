package sleeper

import "fmt"

type State struct {
	Week               int    `json:"week"`
	SeasonType         string `json:"season_type"`
	Season             string `json:"season"`
	PreviousSeason     string `json:"previous_season"`
	Leg                int    `json:"leg"`
	LeagueSeason       string `json:"league_season"`
	LeagueCreateSeason string `json:"league_create_season"`
	DisplayWeek        int    `json:"display_week"`
}

/*
This endpoint returns information about the current state for any sport.

sport   (required) : nfl, nba, lcs, etc...
*/
func (c *Client) GetCurrentState(sport Sport) (State, error) {
	// https://api.sleeper.app/v1/state/nfl
	reqURL := fmt.Sprintf("%s/state/%s", c.sleeperURL, sport)

	state := new(State)

	err := c.get(reqURL, state)

	if err != nil {
		return *state, err
	}

	return *state, nil
}
