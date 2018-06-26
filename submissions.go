package listenbrainz

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
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
	ListenedAt int `json:"listened_at,omitempty"`
	Track      `json:"track_metadata"`
}

// Track is a helper struct for marshalling the JSON payload
type Track struct {
	Title  string `json:"track_name"`
	Artist string `json:"artist_name"`
	Album  string `json:"release_name"`
}

// format listen info as JSON for ListenBrainz
func FormatAsJSON(t Track, listenType string) ([]byte, error) {
	// insert values into struct
	var p Payload
	if listenType == "playing_now" {
		p = Payload{Track: t}
	} else if listenType == "single" {
		p = Payload{
			ListenedAt: int(time.Now().Unix()),
			Track:      t,
		}
	} else {
		// there's nothing to return
		return nil, errors.New("Unrecognized listen type.")
	}

	sp := Submission{
		ListenType: listenType,
		Payloads: Payloads{
			p,
		},
	}

	// convert struct to JSON and return it
	pm, err := json.Marshal(sp)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

// GetSubmissionTime returns the number of seconds after which a track should be submitted.
func GetSubmissionTime(length int) (int, error) {
	hp := int(math.Floor(length / 2))

	// source: https://listenbrainz.readthedocs.io/en/latest/dev/api.html
	// Listens should be submitted for tracks when the user has listened to
	// half the track or 4 minutes of the track, whichever is lower. If the
	// user hasn’t listened to 4 minutes or half the track, it doesn’t fully
	// count as a listen and should not be submitted.
	var st int
	if hp > 240 {
		st = 240
	} else {
		st = hp
	}
	return st, nil
}

// SubmitRequest creates and executes a request based on the JSON passed, to the account delineated by the token.
func SubmitRequest(json []byte, token string) (*http.Response, error) {
	url := "https://api.listenbrainz.org/1/submit-listens"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SubmitPlayingNow posts the given Status to ListenBrainz as what's playing now.
func SubmitPlayingNow(t Track, token string) (*http.Response, error) {
	j, err := FormatAsJSON(t, "playing_now")
	if err != nil {
		return nil, err
	}

	response, err := SubmitRequest(j, token)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SubmitSingle posts the given Status to ListenBrainz as a single listen.
func SubmitSingle(t Track, token string) (*http.Response, error) {
	j, err := FormatAsJSON(t, "single")
	if err != nil {
		return nil, err
	}

	response, err := SubmitRequest(j, token)
	if err != nil {
		return nil, err
	}

	return response, nil
}
