package sleeper

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	client, err = NewClient(&hClient)

	if err != nil {
		t.Error(err)
	}

	res, err := client.GetUser("446778421305929728")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{res.Username, "kylewithanr"},
		{res.UserID, "446778421305929728"},
		{res.DisplayName, "KyleWithAnR"},
	}

	for _, test := range tests {
		if test.test != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, test.test)
		}
	}
}
