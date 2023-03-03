package sleeper

import (
	"fmt"
	"net/url"
	"strconv"
)

// AllPlayers is the return type for the all players endpoint.
type AllPlayers map[string]PlayerInfo

// PlayerInfo is information about a single player in the all players endpoint.
type PlayerInfo struct {
	InjuryStatus          any      `json:"injury_status"`
	SportradarID          string   `json:"sportradar_id"`
	Number                int      `json:"number"`
	FirstName             string   `json:"first_name"`
	PracticeParticipation any      `json:"practice_participation"`
	SearchFullName        string   `json:"search_full_name"`
	Hashtag               string   `json:"hashtag"`
	DepthChartOrder       any      `json:"depth_chart_order"`
	FantasyDataID         int      `json:"fantasy_data_id"`
	Sport                 string   `json:"sport"`
	InjuryBodyPart        any      `json:"injury_body_part"`
	RotowireID            int      `json:"rotowire_id"`
	Metadata              any      `json:"metadata"`
	SearchFirstName       string   `json:"search_first_name"`
	LastName              string   `json:"last_name"`
	PracticeDescription   any      `json:"practice_description"`
	FantasyPositions      []string `json:"fantasy_positions"`
	Status                string   `json:"status"`
	NewsUpdated           any      `json:"news_updated"`
	HighSchool            any      `json:"high_school"`
	Position              string   `json:"position"`
	SwishID               any      `json:"swish_id"`
	SearchLastName        string   `json:"search_last_name"`
	InjuryStartDate       any      `json:"injury_start_date"`
	Weight                string   `json:"weight"`
	DepthChartPosition    any      `json:"depth_chart_position"`
	PlayerID              string   `json:"player_id"`
	InjuryNotes           any      `json:"injury_notes"`
	YahooID               int      `json:"yahoo_id"`
	PandascoreID          any      `json:"pandascore_id"`
	Height                string   `json:"height"`
	StatsID               any      `json:"stats_id"`
	BirthCity             any      `json:"birth_city"`
	RotoworldID           any      `json:"rotoworld_id"`
	BirthState            any      `json:"birth_state"`
	Team                  any      `json:"team"`
	College               string   `json:"college"`
	GsisID                any      `json:"gsis_id"`
	Age                   int      `json:"age"`
	BirthCountry          any      `json:"birth_country"`
	Active                bool     `json:"active"`
	SearchRank            int      `json:"search_rank"`
	BirthDate             string   `json:"birth_date"`
	YearsExp              int      `json:"years_exp"`
	EspnID                int      `json:"espn_id"`
	FullName              string   `json:"full_name"`
}

// TrendingPlayersJSON is the return type of the trending players API.
type TrendingPlayersJSON []struct {
	PlayerID string `json:"player_id"`
	Count    int    `json:"count"`
}

// TrendingType ensures users give valid options
type TrendingType string

const (
	// Add gives the trending players by being added
	Add TrendingType = "add"
	// Drop gives the trending players by being dropped
	Drop TrendingType = "drop"
)

// RequestOption ...
type TrendingRequestOption func(*trendingRequestOptions)

type trendingRequestOptions struct {
	urlParams url.Values
}

// LimitOpt – Number of results you want, (default is 25)
func LimitOpt(amount int) TrendingRequestOption {
	return func(o *trendingRequestOptions) {
		o.urlParams.Set("limit", strconv.Itoa(amount))
	}
}

// LoopbackHoursOpt – Number of hours to look back (default is 24)
func LoopbackHoursOpt(page int) TrendingRequestOption {
	return func(o *trendingRequestOptions) {
		o.urlParams.Set("loopback_hours", strconv.Itoa(page))
	}
}

func processTrendingOptions(options ...TrendingRequestOption) trendingRequestOptions {
	o := trendingRequestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}

	return o
}

/*
GetAllPlayers retrieves all players in the Sleeper DB

Please use this call sparingly, as it is intended only to be used once
per day at most to keep your player IDs updated. The average size of this
query is 5MB.

Since rosters and draft picks contain Player IDs which look like "1042",
"2403", "CAR", etc, you will need to know what those IDs map to. The
/players call provides you the map necessary to look up any player.

You should save this information on your own servers as this is not intended
to be called every time you need to look up players due to the file size being
close to 5MB in size. You do not need to call this endpoint more than once
per day.

https://docs.sleeper.com/#fetch-all-players

sport (required) : The sport to get all players for
*/
func (c *Client) GetAllPlayers(sport Sport) (AllPlayers, error) {
	// https://api.sleeper.app/v1/players/nfl
	reqURL := fmt.Sprintf("%s/players/%s", c.sleeperURL, sport)

	players := new(AllPlayers)

	err := c.get(reqURL, players)

	if err != nil {
		return *players, err
	}

	return *players, nil
}

/*
GetTrendingPlayers retrieves the trending players (adds / drops) over specified time period

Please give attribution to Sleeper you are using our trending data.
If you'd like to embed our trending list on your website or blog,
please use the embed code on the right.

You can use this endpoint to get a list of trending players based on
adds or drops in the past 24 hours.

https://docs.sleeper.com/#trending-players

sport			(required) : The sport to get trending players for
addOrDrop		(required) : Either add or drop
lookback_hours 	(optional) : Number of hours to look back (default is 24)
limit			(optional) : Number of results you want, (default is 25)
*/
func (c *Client) GetTrendingPlayers(sport Sport, addOrDrop TrendingType, opts ...TrendingRequestOption) (TrendingPlayersJSON, error) {
	// https://api.sleeper.app/v1/players/<sport>/trending/<type>?lookback_hours=<hours>&limit=<int>
	reqURL := fmt.Sprintf("%s/players/%s/trending/%s", c.sleeperURL, sport, addOrDrop)

	values := processTrendingOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		reqURL += "?" + query
	}

	players := new(TrendingPlayersJSON)

	err := c.get(reqURL, players)

	if err != nil {
		return *players, err
	}

	return *players, nil
}
