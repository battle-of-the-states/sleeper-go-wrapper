package sleeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCurrentState(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetCurrentLeagueState("nfl")

	if err != nil {
		t.Error(err)
	}

	require.GreaterOrEqual(t, 0, res.Week)
	require.NotEmpty(t, res.SeasonType)
	require.NotEmpty(t, res.Season)
	require.NotEmpty(t, res.PreviousSeason)
	require.GreaterOrEqual(t, 0, res.Leg)
	require.NotEmpty(t, res.LeagueSeason)
	require.NotEmpty(t, res.LeagueCreateSeason)
	require.GreaterOrEqual(t, 0, res.DisplayWeek)
}
