package sleeper

import (
	"fmt"
	"testing"
	"time"
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

	var tests = []struct {
		test     string
		expected string
	}{
		{res.Season, fmt.Sprint(time.Now().Year())},
		{res.LeagueSeason, fmt.Sprint(time.Now().Year())},
		{res.PreviousSeason, fmt.Sprint(time.Now().Year() - 1)},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}
