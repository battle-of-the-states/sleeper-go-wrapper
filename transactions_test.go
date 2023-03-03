package sleeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTransactions(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetTransactions("732292628472778752", 4)

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, 9, len(res))

	// Waiver Transaction

	waiver := res[0]

	require.Equal(t, TrnWaiver, waiver.Type)
	require.Equal(t, "751216071809306624", waiver.TransactionID)
	require.Equal(t, int64(1633504014704), waiver.StatusUpdated)
	require.Equal(t, "complete", waiver.Status)
	require.Equal(t, 2, (*waiver.Settings)["waiver_bid"])
	require.Equal(t, 2, (*waiver.Settings)["seq"])
	require.Equal(t, 6, waiver.RosterIDs[0])
	require.Equal(t, "Your waiver claim was processed successfully!", waiver.Metadata.Notes)
	require.Equal(t, 4, waiver.Leg)
	require.Equal(t, 6, (*waiver.Drops)["5046"])
	require.Empty(t, waiver.DraftPicks)
	require.Equal(t, "603031406913851392", waiver.Creator)
	require.Equal(t, int64(1633466377740), waiver.Created)
	require.Equal(t, 6, waiver.ConsenterIDs[0])
	require.Equal(t, 6, (*waiver.Adds)["2711"])

	// Free Agent Transaction

	freeAgent := res[2]

	require.Equal(t, TrnFreeAgent, freeAgent.Type)
	require.Equal(t, "750212452410658816", freeAgent.TransactionID)
	require.Equal(t, int64(1633227096229), freeAgent.StatusUpdated)
	require.Equal(t, "complete", freeAgent.Status)
	require.Nil(t, freeAgent.Settings)
	require.Equal(t, 7, freeAgent.RosterIDs[0])
	require.Nil(t, freeAgent.Metadata)
	require.Equal(t, 4, freeAgent.Leg)
	require.Equal(t, 7, (*freeAgent.Drops)["4994"])
	require.Empty(t, freeAgent.DraftPicks)
	require.Equal(t, "603077304578543616", freeAgent.Creator)
	require.Equal(t, int64(1633227096229), freeAgent.Created)
	require.Equal(t, 7, freeAgent.ConsenterIDs[0])
	require.Nil(t, freeAgent.Adds)

	// Trade Transaction

	trade := res[3]

	require.Equal(t, TrnTrade, trade.Type)
	require.Equal(t, "749867890106335232", trade.TransactionID)
	require.Equal(t, int64(1633146950678), trade.StatusUpdated)
	require.Equal(t, "complete", trade.Status)
	require.Nil(t, trade.Settings)
	require.Equal(t, 12, trade.RosterIDs[0])
	require.Equal(t, 10, trade.RosterIDs[1])
	require.Nil(t, trade.Metadata)
	require.Equal(t, 4, trade.Leg)
	require.Equal(t, 10, (*trade.Drops)["947"])
	require.Equal(t, 12, (*trade.Drops)["6801"])

	tradedDraftPick := trade.DraftPicks[0]
	require.Equal(t, "2023", tradedDraftPick.Season)
	require.Equal(t, 3, tradedDraftPick.Round)
	require.Equal(t, 10, tradedDraftPick.RosterID)
	require.Equal(t, 10, tradedDraftPick.PreviousOwnerID)
	require.Equal(t, 12, tradedDraftPick.OwnerID)
	require.Nil(t, tradedDraftPick.LeagueID)

	require.Equal(t, "607436038226391040", trade.Creator)
	require.Equal(t, int64(1633144946174), trade.Created)
	require.Equal(t, 12, trade.ConsenterIDs[0])
	require.Equal(t, 10, trade.ConsenterIDs[1])
	require.Equal(t, 12, (*trade.Adds)["947"])
	require.Equal(t, 10, (*trade.Adds)["6801"])
}
