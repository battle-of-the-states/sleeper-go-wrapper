package sleeper

import "fmt"

// PlayoffBracket is the playoff bracket for a league
type PlayoffBracket []PlayoffRound

type PlayoffRound struct {
	Round       int `json:"r"`
	Matchup     int `json:"m"`
	TeamOne     int `json:"t1"`
	TeamTwo     int `json:"t2"`
	Winner      int `json:"w"`
	Loser       int `json:"l"`
	TeamOneFrom struct {
		Winner *int `json:"w,omitempty"`
		Loser  *int `json:"l,omitempty"`
	} `json:"t1_from,omitempty"`
	TeamTwoFrom struct {
		Winner *int `json:"w,omitempty"`
		Loser  *int `json:"l,omitempty"`
	} `json:"t2_from,omitempty"`
}

/*
This endpoint retrieves the playoff bracket for a league for 4, 6, and 8 team playoffs.

Each row represents a matchup between 2 teams.

leagueID (required) : The ID of the league to retrieve bracket from
bracket  (required) : The bracket type to retrieve
*/
func (c *Client) GetLeaguePlayoffBracket(leagueID string, bracket Bracket) (PlayoffBracket, error) {
	// https://api.sleeper.app/v1/league/<league_id>/<bracket_type>
	reqURL := fmt.Sprintf("%s/league/%s/%s", c.sleeperURL, leagueID, bracket)

	playoffBracket := new(PlayoffBracket)

	err := c.get(reqURL, playoffBracket)

	if err != nil {
		return *playoffBracket, err
	}

	return *playoffBracket, nil
}
