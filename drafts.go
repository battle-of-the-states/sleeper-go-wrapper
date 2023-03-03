package sleeper

import "fmt"

// Draft is the struct for a single draft event
type Draft struct {
	Type            string            `json:"type"`
	Status          string            `json:"status"`
	StartTime       int64             `json:"start_time"`
	Sport           string            `json:"sport"`
	Settings        map[string]int    `json:"settings"`
	SeasonType      string            `json:"season_type"`
	Season          string            `json:"season"`
	Metadata        map[string]string `json:"metadata"`
	LeagueID        string            `json:"league_id"`
	LastPicked      int64             `json:"last_picked"`
	LastMessageTime int64             `json:"last_message_time"`
	LastMessageID   string            `json:"last_message_id"`
	DraftOrder      interface{}       `json:"draft_order"`
	DraftID         string            `json:"draft_id"`
	Creators        interface{}       `json:"creators"`
	Created         int64             `json:"created"`
}

// DraftPick represents a single pick in the draft
type DraftPick struct {
	Round     int               `json:"round"`
	RosterID  int               `json:"roster_id"`
	PlayerID  string            `json:"player_id"`
	PickedBy  string            `json:"picked_by"`
	PickNo    int               `json:"pick_no"`
	Metadata  map[string]string `json:"metadata"`
	IsKeeper  interface{}       `json:"is_keeper"`
	DraftSlot int               `json:"draft_slot"`
	DraftID   string            `json:"draft_id"`
}

// DraftTradedPick is a pick that is traded in the draft
type DraftTradedPick struct {
	Season          string      `json:"season"`
	Round           int         `json:"round"`
	RosterID        int         `json:"roster_id"`
	PreviousOwnerID int         `json:"previous_owner_id"`
	OwnerID         int         `json:"owner_id"`
	DraftID         interface{} `json:"draft_id"`
	LeagueID        *string     `json:"league_id"`
}

/*
GetLeagueDrafts retrieves all drafts for a league. Keep in mind that a league can have
multiple drafts, especially dynasty leagues.

Drafts are sorted by most recent to earliest. Most leagues should only have one draft.

leagueID (required) : The ID of the league for which you are trying to retrieve drafts.
*/
func (c *Client) GetLeagueDrafts(leagueID string) ([]Draft, error) {
	// https://api.sleeper.app/v1/league/<league_id>/drafts
	reqURL := fmt.Sprintf("%s/league/%s/drafts", c.sleeperURL, leagueID)

	drafts := new([]Draft)

	err := c.get(reqURL, drafts)

	if err != nil {
		return *drafts, err
	}

	return *drafts, nil
}

/*
GetDraft retrieves a specific draft

draftID (required) : The ID of the draft to retrieve
*/ // Draft is the struct for a single draft event
func (c *Client) GetDraft(draftID string) (Draft, error) {
	// https://api.sleeper.app/v1/draft/<draft_id>
	reqURL := fmt.Sprintf("%s/draft/%s", c.sleeperURL, draftID)
	// Draft is the struct for a single draft event
	draft := new(Draft)

	err := c.get(reqURL, draft)

	if err != nil {
		return *draft, err
	}

	return *draft, nil
}

/*
GetDraft retrieves all picks in a draft.

draftID (required) : The ID of the draft to retrieve picks for
*/
func (c *Client) GetDraftPicks(draftID string) ([]DraftPick, error) {
	// https://api.sleeper.app/v1/draft/<draft_id>/picks
	reqURL := fmt.Sprintf("%s/draft/%s/picks", c.sleeperURL, draftID)

	picks := new([]DraftPick)

	err := c.get(reqURL, picks)

	if err != nil {
		return *picks, err
	}

	return *picks, nil
}

/*
GetDraftTradedPicks retrieves all traded picks in a draft

draftID (required) : The ID of the draft to retrieve traded picks for
*/
func (c *Client) GetDraftTradedPicks(draftID string) ([]DraftTradedPick, error) {
	// https://api.sleeper.app/v1/draft/<draft_id>/traded_picks
	reqURL := fmt.Sprintf("%s/draft/%s/traded_picks", c.sleeperURL, draftID)

	picks := new([]DraftTradedPick)

	err := c.get(reqURL, picks)

	if err != nil {
		return *picks, err
	}

	return *picks, nil
}
