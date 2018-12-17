package listenbrainz

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestGetSubmissionTime(t *testing.T) {
	var tests = []struct {
		Length, Result int
	}{
		{-1, 0},
		{128, 64},
		{math.MaxInt64, 240},
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
		Title:  "b",
		Artist: "a",
		Album:  "c",
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

func TestFormatSingle(t *testing.T) {
	track := Track{
		Title:  "b",
		Artist: "a",
		Album:  "c",
	}

	time := time.Now().Unix()

	ts := Submission{
		ListenType: "single",
		Payloads: Payloads{
			Payload{
				ListenedAt: time,
				Track:      track,
			},
		},
	}

	s := FormatSingle(track, time)

	if !reflect.DeepEqual(ts, s) {
		t.Error("Expected", ts, "got", s)
	}
}
