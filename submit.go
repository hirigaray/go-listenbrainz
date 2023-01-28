// Copyright (C) 2023 Luiz de Milon (kori)

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
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Submission is a struct for marshalling the JSON payload
type Submission struct {
	ListenType string    `json:"listen_type"`
	Payloads   []Payload `json:"payload"`
}

// Payloads is a helper struct for marshalling the JSON payload
type Payloads []Payload

// Payload is a helper struct for marshalling the JSON payload
type Payload struct {
	ListenedAt int64 `json:"listened_at,omitempty"`
	Track      `json:"track_metadata"`
}

// Track is a helper struct for marshalling the JSON payload
type Track struct {
	Title  string `json:"track_name"`
	Artist string `json:"artist_name"`
	Album  string `json:"release_name"`
}

// FormatPlayingNow formats a Track as a playing_now Submission.
func FormatPlayingNow(track Track) Submission {
	return Submission{
		ListenType: "playing_now",
		Payloads: Payloads{
			Payload{
				Track: track,
			},
		},
	}
}

// FormatSingle formats a Track as a single Submission.
func FormatSingle(track Track, time int64) Submission {
	return Submission{
		ListenType: "single",
		Payloads: Payloads{
			Payload{
				ListenedAt: time,
				Track:      track,
			},
		},
	}
}

// GetSubmissionTime returns the number of seconds after which a track
// should be submitted.
func GetSubmissionTime(d time.Duration) (time.Duration, error) {
	if d < 0 {
		return time.Duration(0), errors.New("length can't be negative")
	}

	// source: https://listenbrainz.readthedocs.io/en/latest/dev/api.html
	// Listens should be submitted for tracks when the user has listened to
	// half the track or 4 minutes of the track, whichever is lower. If the
	// user hasn’t listened to 4 minutes or half the track, it doesn’t fully
	// count as a listen and should not be submitted.
	if hp := d / 2; hp > 4*time.Minute {
		return 4 * time.Minute, nil
	} else {
		return hp, nil
	}
}

// SubmitRequest creates and executes a request containing the JSON that's passed,
// to the account delineated by the token.
func SubmitRequest(json []byte, token string) (*http.Response, error) {
	url := "https://api.listenbrainz.org/1/submit-listens"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// SubmitPlayingNow posts the given track to ListenBrainz as what's playing now.
func SubmitPlayingNow(track Track, token string) (*http.Response, error) {
	json, err := json.Marshal(FormatPlayingNow(track))
	if err != nil {
		return nil, err
	}

	return SubmitRequest(json, token)
}

// SubmitSingle posts the given track to ListenBrainz as a single listen
// with the given time.
func SubmitSingle(track Track, token string, time int64) (*http.Response, error) {
	json, err := json.Marshal(FormatSingle(track, time))
	if err != nil {
		return nil, err
	}

	return SubmitRequest(json, token)
}
