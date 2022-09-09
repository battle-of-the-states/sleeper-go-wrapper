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
