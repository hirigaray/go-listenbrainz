package listenbrainz

import (
	"reflect"
	"testing"
)

func TestGetSubmissionTime(t *testing.T) {
	var tests = []struct {
		Length, Result int
	}{
		{100, 50},
		{1000, 240},
	}

	for _, test := range tests {
		st := GetSubmissionTime(test.Length)
		if st != test.Result {
			t.Error("Expected", test.Result, "got", st)
		}
	}
}

func TestFormatPlayingNow(t *testing.T) {
	track := Track{
		"a",
		"b",
		"c",
	}

	ts := Submission{
		ListenType: "playing_now",
		Payloads: Payloads{
			Payload{
				Track: track,
			},
		},
	}

	s := FormatPlayingNow(track)

	if !reflect.DeepEqual(ts, s) {
		t.Error("Expected", ts, "got", s)
	}
}
