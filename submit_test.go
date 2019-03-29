// Copyright (C) 2019 Luiz de Milon (kori)

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
		st, err := GetSubmissionTime(test.Length)
		if err != nil {
			t.Log("Test failed successfully at:", test.Length, ":", err)
		}
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
