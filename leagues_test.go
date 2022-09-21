package sleeper

import (
	"testing"
)

func TestGetLeaguesForUser(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeaguesForUser(NFL, "446778421305929728", "2020")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     int
		expected int
	}{
		{len(res), 6},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %d expected, received: %d", test.expected, test.test)
		}
	}
}

func TestGetLeague(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeague("732292628472778752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res.Status, "complete"},
		{res.Sport, "nfl"},
		{res.SeasonType, "regular"},
		{res.Season, "2021"},
		{res.Name, "Sages Fantasy Football"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetLeagueRosters(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeagueRosters("732292628472778752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].OwnerID, "332664810955108352"},
		{res[0].LeagueID, "732292628472778752"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}

	var settingsTest = []struct {
		test     int
		expected int
	}{
		{res[0].Settings.Wins, 5},
		{res[0].Settings.WaiverPosition, 4},
		{res[0].Settings.WaiverBudgetUsed, 3},
		{res[0].Settings.TotalMoves, 0},
		{res[0].Settings.PptsDecimal, 18},
		{res[0].Settings.Ppts, 1864},
		{res[0].Settings.Losses, 9},
		{res[0].Settings.FptsDecimal, 10},
		{res[0].Settings.FptsAgainstDecimal, 64},
		{res[0].Settings.FptsAgainst, 1575},
		{res[0].Settings.Fpts, 1513},
	}

	for _, test := range settingsTest {
		if test.test != test.expected {
			t.Errorf("Test Failed: %d expected, received: %d", test.expected, test.test)
		}
	}

	var settingsMethodTest = []struct {
		test     float32
		expected float32
	}{
		{res[0].GetTotalFpts(), 1513.10},
		{res[0].GetTotalFptsAgainst(), 1575.64},
		{res[0].GetTotalPpts(), 1864.18},
		{res[0].GetManagerEfficiency(), 81.167046},
	}

	for _, test := range settingsMethodTest {
		if test.test != test.expected {
			t.Errorf("Test Failed: %f expected, received: %f", test.expected, test.test)
		}
	}
}

func TestGetLeagueUsers(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeagueUsers("732292628472778752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].UserID, "332664810955108352"},
		{res[0].LeagueID, "732292628472778752"},
		{res[0].Metadata.MentionPn, "on"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetLeagueMatchups(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeagueMatchups("732292628472778752", 4)

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     int
		expected int
	}{
		{res[0].RosterID, 1},
		{res[0].MatchupID, 4},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %d expected, received: %d", test.expected, test.test)
		}
	}

	var testPlayerPoints = []struct {
		test     float32
		expected float32
	}{
		{res[0].PlayerPoints["391"], 14},
		{res[0].PlayerPoints["3161"], 17.92},
	}

	for _, test := range testPlayerPoints {
		if test.test != test.expected {
			t.Errorf("Test Failed: %f expected, received: %f", test.expected, test.test)
		}
	}
}

func TestGetLeaguePlayoffBracket(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeaguePlayoffBracket("732292628472778752", WinnersBracket)

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     int
		expected int
	}{
		{res[0].Winner, 11},
		{res[0].Round, 1},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %d expected, received: %d", test.expected, test.test)
		}
	}
}
