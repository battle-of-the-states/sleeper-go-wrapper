package sleeper

import (
	"fmt"
)

// LeagueInfo is the return type of the league info API.
type LeagueInfo struct {
	TotalRosters int          `json:"total_rosters"`
	Status       LeagueStatus `json:"status"`
	Sport        Sport        `json:"sport"`
	Shard        int          `json:"shard"`
	Settings     struct {
		MaxKeepers           int `json:"max_keepers"`
		DraftRounds          int `json:"draft_rounds"`
		TradeReviewDays      int `json:"trade_review_days"`
		Squads               int `json:"squads"`
		ReserveAllowDnr      int `json:"reserve_allow_dnr"`
		CapacityOverride     int `json:"capacity_override"`
		PickTrading          int `json:"pick_trading"`
		DisableTrades        int `json:"disable_trades"`
		TaxiYears            int `json:"taxi_years"`
		TaxiAllowVets        int `json:"taxi_allow_vets"`
		BestBall             int `json:"best_ball"`
		LastReport           int `json:"last_report"`
		DisableAdds          int `json:"disable_adds"`
		WaiverType           int `json:"waiver_type"`
		BenchLock            int `json:"bench_lock"`
		ReserveAllowSus      int `json:"reserve_allow_sus"`
		Type                 int `json:"type"`
		ReserveAllowCov      int `json:"reserve_allow_cov"`
		WaiverClearDays      int `json:"waiver_clear_days"`
		DailyWaiversLastRan  int `json:"daily_waivers_last_ran"`
		WaiverDayOfWeek      int `json:"waiver_day_of_week"`
		StartWeek            int `json:"start_week"`
		PlayoffTeams         int `json:"playoff_teams"`
		NumTeams             int `json:"num_teams"`
		ReserveSlots         int `json:"reserve_slots"`
		PlayoffRoundType     int `json:"playoff_round_type"`
		DailyWaiversHour     int `json:"daily_waivers_hour"`
		WaiverBudget         int `json:"waiver_budget"`
		ReserveAllowOut      int `json:"reserve_allow_out"`
		OffseasonAdds        int `json:"offseason_adds"`
		PlayoffSeedType      int `json:"playoff_seed_type"`
		DailyWaivers         int `json:"daily_waivers"`
		PlayoffWeekStart     int `json:"playoff_week_start"`
		DailyWaiversDays     int `json:"daily_waivers_days"`
		LeagueAverageMatch   int `json:"league_average_match"`
		Leg                  int `json:"leg"`
		TradeDeadline        int `json:"trade_deadline"`
		ReserveAllowDoubtful int `json:"reserve_allow_doubtful"`
		TaxiDeadline         int `json:"taxi_deadline"`
		ReserveAllowNa       int `json:"reserve_allow_na"`
		TaxiSlots            int `json:"taxi_slots"`
		PlayoffType          int `json:"playoff_type"`
	} `json:"settings"`
	SeasonType       SeasonType         `json:"season_type"`
	Season           string             `json:"season"`
	ScoringSettings  map[string]float64 `json:"scoring_settings"`
	RosterPositions  []string           `json:"roster_positions"`
	PreviousLeagueID string             `json:"previous_league_id"`
	Name             string             `json:"name"`
	Metadata         struct {
		LatestLeagueWinnerRosterID string `json:"latest_league_winner_roster_id"`
		Continued                  string `json:"continued"`
	} `json:"metadata"`
	LoserBracketID        interface{} `json:"loser_bracket_id"`
	LeagueID              string      `json:"league_id"`
	LastReadID            string      `json:"last_read_id"`
	LastPinnedMessageID   string      `json:"last_pinned_message_id"`
	LastMessageTime       int64       `json:"last_message_time"`
	LastMessageTextMap    interface{} `json:"last_message_text_map"`
	LastMessageID         string      `json:"last_message_id"`
	LastMessageAttachment interface{} `json:"last_message_attachment"`
	LastAuthorIsBot       bool        `json:"last_author_is_bot"`
	LastAuthorID          string      `json:"last_author_id"`
	LastAuthorDisplayName string      `json:"last_author_display_name"`
	LastAuthorAvatar      string      `json:"last_author_avatar"`
	GroupID               interface{} `json:"group_id"`
	DraftID               string      `json:"draft_id"`
	CompanyID             interface{} `json:"company_id"`
	BracketID             int64       `json:"bracket_id"`
	Avatar                string      `json:"avatar"`
}

/*
This endpoint retrieves all leagues for a user.

sport	(required) : We only support "nfl" right now.
userID (required) : The numerical ID of the user.
season	(required) : Season can be 2017, 2018, etc...
*/
func (c *Client) GetLeaguesForUser(sport Sport, userID, season string) ([]LeagueInfo, error) {
	// https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>
	reqURL := fmt.Sprintf("%s/user/%s/leagues/%s/%s", c.sleeperURL, userID, sport, season)

	leagues := new([]LeagueInfo)

	err := c.get(reqURL, leagues)

	if err != nil {
		return *leagues, err
	}

	return *leagues, nil
}

/*
This endpoint retrieves all leagues for a user.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeague(leagueID string) (LeagueInfo, error) {
	// https://api.sleeper.app/v1/league/<league_id>
	reqURL := fmt.Sprintf("%s/league/%s", c.sleeperURL, leagueID)

	league := new(LeagueInfo)

	err := c.get(reqURL, league)

	if err != nil {
		return *league, err
	}

	return *league, nil
}

/*
This endpoint retrieves all users in a league.

This also includes each user's display_name, avatar, and their
metadata which sometimes includes a nickname they gave their team.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeagueUsers(leagueID string) ([]User, error) {
	// https://api.sleeper.app/v1/league/<league_id>/users
	reqURL := fmt.Sprintf("%s/league/%s/users", c.sleeperURL, leagueID)

	users := new([]User)

	err := c.get(reqURL, users)

	if err != nil {
		return *users, err
	}

	return *users, nil
}
