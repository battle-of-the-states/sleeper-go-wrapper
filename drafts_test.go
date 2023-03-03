package sleeper

import (
	"testing"
)

func TestGetUserDrafts(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetUserDrafts("446778421305929728", NFL, "2022")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].Type, "snake"},
		{res[0].Status, "complete"},
		{res[0].Sport, "nfl"},
		{res[0].Season, "2022"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetLeagueDrafts(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetLeagueDrafts("732292628472778752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].Type, "linear"},
		{res[0].Status, "complete"},
		{res[0].Sport, "nfl"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetDraft(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetDraft("589201391218634752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res.Type, "snake"},
		{res.Status, "complete"},
		{res.Sport, "nfl"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetDraftPicks(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetDraftPicks("589201391218634752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].PlayerID, "4034"},
		{res[0].DraftID, "589201391218634752"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}

func TestGetDraftTradedPicks(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetDraftTradedPicks("589201391218634752")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res[0].Season, "2020"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}
