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

	res, err := client.GetCurrentState("nfl")

	if err != nil {
		t.Error(err)
	}

	require.NotEmpty(t, res.Week)
	require.NotEmpty(t, res.SeasonType)
	require.NotEmpty(t, res.Season)
	require.NotEmpty(t, res.PreviousSeason)
	require.NotEmpty(t, res.Leg)
	require.NotEmpty(t, res.LeagueSeason)
	require.NotEmpty(t, res.LeagueCreateSeason)
	require.NotEmpty(t, res.DisplayWeek)
}
