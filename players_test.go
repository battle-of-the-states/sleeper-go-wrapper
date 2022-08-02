package sleeper

import (
	"testing"
)

func TestGetTrendingPlayers(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetTrendingPlayers(NFL, Add, LimitOpt(10))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     int
		expected int
	}{
		{len(res), 10},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %d expected, received: %d", test.expected, test.test)
		}
	}
}
