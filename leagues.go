package sleeper

import (
	"fmt"
)

// LeaguesInfoJSON is the return type for all leagues for user
type LeaguesInfoJSON []LeagueInfoJSON

// LeagueInfoJSON is the return type of the league info API.
type LeagueInfoJSON struct {
	TotalRosters int    `json:"total_rosters"`
	Status       string `json:"status"`
	Sport        string `json:"sport"`
	Shard        int    `json:"shard"`
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
	SeasonType       string             `json:"season_type"`
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

// RostersJSON is the return type for all rosters in a league
type RostersJSON []RosterJSON

// RosterJSON is a single roster from the league rosters API.
type RosterJSON struct {
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

// MatchupsJSON is the return type for all matchups for the week
type MatchupsJSON []MatchupJSON

// MatchupJSON is a single matchup from the league matchups API.
type MatchupJSON struct {
	StartersPoints []float32          `json:"starters_points"`
	Starters       []string           `json:"starters"`
	RosterID       int                `json:"roster_id"`
	Points         float32            `json:"points"`
	PlayerPoints   map[string]float32 `json:"players_points"`
	Players        []string           `json:"players"`
	MatchupID      int                `json:"matchup_id"`
	CustomPoints   float32            `json:"custom_points"`
}

// PlayoffBracketJSON is the playoff bracket for a league
type PlayoffBracketJSON []PlayoffRoundJSON

type PlayoffRoundJSON struct {
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
This endpoint retrieves all leagues for a user.

sport	(required) : We only support "nfl" right now.
userID (required) : The numerical ID of the user.
season	(required) : Season can be 2017, 2018, etc...
*/
func (c *Client) GetLeaguesForUser(sport Sport, userID, season string) (LeaguesInfoJSON, error) {
	// https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>
	lastfmURL := fmt.Sprintf("%s/user/%s/leagues/%s/%s", c.sleeperURL, userID, sport, season)

	leagues := new(LeaguesInfoJSON)

	err := c.get(lastfmURL, leagues)

	if err != nil {
		return *leagues, err
	}

	return *leagues, nil
}

/*
This endpoint retrieves all leagues for a user.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeague(leagueID string) (LeagueInfoJSON, error) {
	// https://api.sleeper.app/v1/league/<league_id>
	lastfmURL := fmt.Sprintf("%s/league/%s", c.sleeperURL, leagueID)

	league := new(LeagueInfoJSON)

	err := c.get(lastfmURL, league)

	if err != nil {
		return *league, err
	}

	return *league, nil
}

/*
This endpoint retrieves all rosters in a league.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeagueRosters(leagueID string) (RostersJSON, error) {
	// https://api.sleeper.app/v1/league/<league_id>/rosters
	lastfmURL := fmt.Sprintf("%s/league/%s/rosters", c.sleeperURL, leagueID)

	rosters := new(RostersJSON)

	err := c.get(lastfmURL, rosters)

	if err != nil {
		return *rosters, err
	}

	return *rosters, nil
}

/*
This endpoint retrieves all users in a league.

This also includes each user's display_name, avatar, and their
metadata which sometimes includes a nickname they gave their team.

leagueID (required) : The ID of the league to retrieve
*/
func (c *Client) GetLeagueUsers(leagueID string) (UsersJSON, error) {
	// https://api.sleeper.app/v1/league/<league_id>/users
	lastfmURL := fmt.Sprintf("%s/league/%s/users", c.sleeperURL, leagueID)

	users := new(UsersJSON)

	err := c.get(lastfmURL, users)

	if err != nil {
		return *users, err
	}

	return *users, nil
}

/*
This endpoint retrieves all users in a league.

leagueID (required) : The ID of the league to retrieve
week 	 (required) : The week number to get the matchups
*/
func (c *Client) GetLeagueMatchups(leagueID string, week int8) (MatchupsJSON, error) {
	// https://api.sleeper.app/v1/league/<league_id>/matchups/<week>
	lastfmURL := fmt.Sprintf("%s/league/%s/matchups/%d", c.sleeperURL, leagueID, week)

	matchups := new(MatchupsJSON)

	err := c.get(lastfmURL, matchups)

	if err != nil {
		return *matchups, err
	}

	return *matchups, nil
}

/*
This endpoint retrieves all users in a league.

leagueID (required) : The ID of the league to retrieve
week 	 (required) : The week number to get the matchups
*/
func (c *Client) GetLeaguePlayoffBracket(leagueID string, bracket Bracket) (PlayoffBracketJSON, error) {
	// https://api.sleeper.app/v1/league/<league_id>/<bracket_type>
	lastfmURL := fmt.Sprintf("%s/league/%s/%s", c.sleeperURL, leagueID, bracket)

	playoffBracket := new(PlayoffBracketJSON)

	err := c.get(lastfmURL, playoffBracket)

	if err != nil {
		return *playoffBracket, err
	}

	return *playoffBracket, nil
}

/*
	Methods for leagues structs
*/

// GetTotalFpts combines fpts and fpts_decimal for a total fpts
func (r RosterJSON) GetTotalFpts() float32 {
	return float32(r.Settings.Fpts) + float32(r.Settings.FptsDecimal)/100
}

// GetTotalPpts combines ppts and ppts_decimal for a total ppts
func (r RosterJSON) GetTotalPpts() float32 {
	return float32(r.Settings.Ppts) + float32(r.Settings.PptsDecimal)/100
}

// GetTotalFptsAgainst combines fpts_against and fpts_against_decimal for a total fpts_against
func (r RosterJSON) GetTotalFptsAgainst() float32 {
	return float32(r.Settings.FptsAgainst) + float32(r.Settings.FptsAgainstDecimal)/100
}

// GetManagerEfficiency will get the percentage of fpts the manager capitalized on vs what was possible
func (r RosterJSON) GetManagerEfficiency() float32 {
	totalFpts := r.GetTotalFpts()
	totalPpts := r.GetTotalPpts()

	return (totalFpts / totalPpts) * 100
}

// GetTotalFptsPerWeek will average the total points with the total weeks played for points / week
func (r RosterJSON) GetTotalFptsPerWeek() float32 {
	return r.GetTotalFpts() / float32(r.Settings.Wins+r.Settings.Losses)
}
