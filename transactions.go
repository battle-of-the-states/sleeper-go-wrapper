package sleeper

import "fmt"

type Transaction struct {
	Type          TransactionType `json:"type"`
	TransactionID string          `json:"transaction_id"`
	StatusUpdated int64           `json:"status_updated"`
	Status        string          `json:"status"`
	// Settings -- Trades do not use this field; In free agent it could track waiver_bid
	Settings *map[string]int `json:"settings"`
	// RosterIDs are the roster IDs involved in this transaction
	RosterIDs []int `json:"roster_ids"`
	// Metadata can contain notes in waivers like why it didn't go through
	Metadata *struct {
		Notes string `json:"notes"`
	} `json:"metadata"`
	// Leg is the week in NFL
	Leg int `json:"leg"`
	// Drops is a map with player ID and the roster ID that dropped the respective player
	Drops *map[string]int `json:"drops"`
	// DraftPicks are the picks that were traded
	DraftPicks []DraftTradedPick `json:"draft_picks"`
	// Creator is the user ID who initiated the transaction
	Creator string `json:"creator"`
	Created int64  `json:"created"`
	// ConsenterIDs are the roster IDs of the people who agreed to this transaction
	ConsenterIDs []int `json:"consenter_ids"`
	// Adds is a map with player ID and the roster ID that added the respective player
	Adds         *map[string]int `json:"adds"`
	WaiverBudget []struct {
		Sender   int `json:"sender"`
		Receiver int `json:"receiver"`
		Amount   int `json:"amount"`
	} `json:"waiver_budget"`
}

/*
This endpoint retrieves all free agent transactions, waivers and trades.

https://docs.sleeper.com/#get-transactions

leagueID (required) : The ID of the league to retrieve bracket from
round  (required) : The week you want to pull from
*/
func (c *Client) GetTransactions(leagueID string, round int) ([]Transaction, error) {
	// https://api.sleeper.app/v1/league/<league_id>/transactions/<round>
	reqURL := fmt.Sprintf("%s/league/%s/transactions/%d", c.sleeperURL, leagueID, round)

	transactions := new([]Transaction)

	err := c.get(reqURL, transactions)

	if err != nil {
		return *transactions, err
	}

	return *transactions, nil
}
