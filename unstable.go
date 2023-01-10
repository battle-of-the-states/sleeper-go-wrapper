package sleeper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type UnstableStatPlayer struct {
	YearsExp         int                     `json:"years_exp"`
	Team             string                  `json:"team"`
	Position         string                  `json:"position"`
	NewsUpdated      int64                   `json:"news_updated"`
	Metadata         *map[string]interface{} `json:"metadata"`
	LastName         string                  `json:"last_name"`
	InjuryStatus     *string                 `json:"injury_status"`
	InjuryStartDate  *interface{}            `json:"injury_start_date"`
	InjuryNotes      *interface{}            `json:"injury_notes"`
	InjuryBodyPart   *string                 `json:"injury_body_part"`
	FirstName        string                  `json:"first_name"`
	FantasyPositions []string                `json:"fantasy_positions"`
}

// UnstablePlayerStatsJSON is an unstable struct for player stats and projections
type UnstablePlayerStatsJSON struct {
	Week         *int                `json:"week"`
	Team         string              `json:"team"`
	Stats        map[string]float32  `json:"stats"`
	Sport        Sport               `json:"sport"`
	SeasonType   string              `json:"season_type"`
	Season       string              `json:"season"`
	PlayerID     string              `json:"player_id"`
	Player       *UnstableStatPlayer `json:"player,omitempty"`
	Opponent     *string             `json:"opponent"`
	LastModified *int64              `json:"last_modified"`
	GameID       string              `json:"game_id"`
	Date         *string             `json:"date"`
	Company      string              `json:"company"`
	Category     string              `json:"category"`
}

// RequestOption ...
type UnstableStatsRequestOption func(*unstableStatsRequestOptions)

type unstableStatsRequestOptions struct {
	urlParams url.Values
}

// PositionOpt – NFL player position you want
func PositionOpt(position FFPosition) UnstableStatsRequestOption {
	return func(o *unstableStatsRequestOptions) {
		o.urlParams.Add("position", string(position))
	}
}

// SeasonOpt – Fantasy Football season
func SeasonOpt(season string) UnstableStatsRequestOption {
	return func(o *unstableStatsRequestOptions) {
		o.urlParams.Set("season", season)
	}
}

// WeekOpt – Fantasy Football week
func WeekOpt(week int) UnstableStatsRequestOption {
	return func(o *unstableStatsRequestOptions) {
		o.urlParams.Set("week", strconv.Itoa(week))
	}
}

// SeasonTypeOpt – Fantasy Football season type {pre, regular, post}
func SeasonTypeOpt(st SeasonType) UnstableStatsRequestOption {
	return func(o *unstableStatsRequestOptions) {
		o.urlParams.Add("season_type", string(st))
	}
}

func processUnstableStatsOptions(options ...UnstableStatsRequestOption) unstableStatsRequestOptions {
	o := unstableStatsRequestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}

	return o
}

/*
**UNSTABLE**

This endpoint could change at any time without notice.

You can use this endpoint to get a player's projections or stats for the given week or season if week is nil.


sport			(required) : The sport to get player projections for
playerID		(required) : The Sleeper ID of the player
seasonType		(required) : The type of season {pre, post, regular}
season			(required) : The season the projections are for
week			(optional) : Week number to retrieve stats/projections for (entire season is default)
*/
func (c *Client) UnstableGetPlayerInfo(sport Sport, infoType PlayerInfoType, playerID string, opts ...UnstableStatsRequestOption) (UnstablePlayerStatsJSON, error) {
	// https://api.sleeper.app/{stats or projections}/{sport}/player/{player_id}?season_type=regular&season={year}&week={week_number}
	reqURL := fmt.Sprintf("%s/%s/%s/player/%s", SLEEPER_BASE_URL_UNSTABLE, infoType, sport, playerID)

	values := processUnstableStatsOptions(opts...).urlParams

	// Need to make a check that the required queries are actually set, else throw an error
	if !values.Has("season") || !values.Has("season_type") {
		return UnstablePlayerStatsJSON{}, errors.New("season and season_type are required options you must pass")
	}

	if query := values.Encode(); query != "" {
		reqURL += "?" + query
	}

	stat := new(UnstablePlayerStatsJSON)

	err := c.get(reqURL, stat)

	if err != nil {
		return *stat, err
	}

	return *stat, nil
}

/*
**UNSTABLE**

This endpoint could change at any time without notice.

You can use this endpoint to get a player's stats for the given week or season if week is nil.


sport			(required) : The sport to get player projections for
playerID		(required) : The Sleeper ID of the player
seasonType		(required) : The type of season {pre, post, regular}
season			(required) : The season the projections are for
week			(optional) : The week number the projections are for
*/
func (c *Client) UnstableGetPlayersInfoByPosition(sport Sport, infoType PlayerInfoType, season string, week int, opts ...UnstableStatsRequestOption) ([]UnstablePlayerStatsJSON, error) {
	// https://api.sleeper.app/{stats or projections}/{sport}/{season}/{week}?season_type=regular&position=DEF&position=K&position=QB&position=RB&position=TE&position=WR
	reqURL := fmt.Sprintf("%s/%s/%s/%s/%d", SLEEPER_BASE_URL_UNSTABLE, infoType, sport, season, week)

	values := processUnstableStatsOptions(opts...).urlParams

	// Need to make a check that the required queries are actually set, else throw an error
	if !values.Has("season_type") || !values.Has("position") {
		return []UnstablePlayerStatsJSON{}, errors.New("season and season_type are required options you must pass")
	}

	if query := values.Encode(); query != "" {
		reqURL += "?" + query
	}

	stat := new([]UnstablePlayerStatsJSON)

	err := c.get(reqURL, stat)

	if err != nil {
		return *stat, err
	}

	return *stat, nil
}
