package sleeper

import "fmt"

// Roster is a single roster from the league rosters API.
type Roster struct {
	Taxi     []string `json:"taxi"`
	Starters []string `json:"starters"`
	Settings struct {
		Wins               int `json:"wins"`
		WaiverPosition     int `json:"waiver_position"`
		WaiverBudgetUsed   int `json:"waiver_budget_used"`
		TotalMoves         int `json:"total_moves"`
		Ties               int `json:"ties"`
		PptsDecimal        int `json:"ppts_decimal"`
		Ppts               int `json:"ppts"`
		Losses             int `json:"losses"`
		FptsDecimal        int `json:"fpts_decimal"`
		FptsAgainstDecimal int `json:"fpts_against_decimal"`
		FptsAgainst        int `json:"fpts_against"`
		Fpts               int `json:"fpts"`
		Division           int `json:"division,omitempty"`
	} `json:"settings"`
	RosterID  int         `json:"roster_id"`
	Reserve   []string    `json:"reserve"`
	Players   []string    `json:"players"`
	PlayerMap interface{} `json:"player_map"`
	OwnerID   string      `json:"owner_id"`
	LeagueID  string      `json:"league_id"`
	CoOwners  interface{} `json:"co_owners"`
	// known keys:
	// streak, record, p_nick_<player ID>, allow_pn_scoring, allow_pn_news
	Metadata interface{} `json:"metadata"`
}

/*
This endpoint retrieves all rosters in a league.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeagueRosters(leagueID string) ([]Roster, error) {
	// https://api.sleeper.app/v1/league/<league_id>/rosters
	reqURL := fmt.Sprintf("%s/league/%s/rosters", c.sleeperURL, leagueID)

	rosters := new([]Roster)

	err := c.get(reqURL, rosters)

	if err != nil {
		return *rosters, err
	}

	return *rosters, nil
}

/*
	Roster methods
*/

// GetTotalFpts combines fpts and fpts_decimal for a total fpts
func (r Roster) GetTotalFpts() float32 {
	return float32(r.Settings.Fpts) + float32(r.Settings.FptsDecimal)/100
}

// GetTotalPpts combines ppts and ppts_decimal for a total ppts
func (r Roster) GetTotalPpts() float32 {
	return float32(r.Settings.Ppts) + float32(r.Settings.PptsDecimal)/100
}

// GetTotalFptsAgainst combines fpts_against and fpts_against_decimal for a total fpts_against
func (r Roster) GetTotalFptsAgainst() float32 {
	return float32(r.Settings.FptsAgainst) + float32(r.Settings.FptsAgainstDecimal)/100
}

// GetManagerEfficiency will get the percentage of fpts the manager capitalized on vs what was possible
func (r Roster) GetManagerEfficiency() float32 {
	totalFpts := r.GetTotalFpts()
	totalPpts := r.GetTotalPpts()
	// Avoid division by 0
	if totalPpts == 0 {
		return 0
	}

	return (totalFpts / totalPpts) * 100
}

// GetTotalFptsPerWeek will average the total points with the total weeks played for points / week
func (r Roster) GetTotalFptsPerWeek() float32 {
	games := float32(r.Settings.Wins + r.Settings.Losses)
	// Avoid division by 0
	if games == 0 {
		return 0
	}
	return r.GetTotalFpts() / games
}
