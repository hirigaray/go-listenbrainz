package listenbrainz

import (
	"log"
	"testing"
)

func TestFormatAsJSON(t *testing.T) {
	track := Track{
		"a",
		"b",
		"c",
	}
	json, err := FormatAsJSON(track, "playing_now")
	if err != nil {
		log.Fatalln(err)
	}

	if string(json) != `{"listen_type":"playing_now","payload":[{"track_metadata":{"track_name":"a","artist_name":"b","release_name":"c"}}]}` {
		t.Error(`Expected "{"listen_type":"playing_now","payload":[{"track_metadata":{"track_name":"a","artist_name":"b","release_name":"c"}}]}", got `, json)
	}
}
