package listenbrainz

import "net/http"

// GetListenHistory fetches the listen history of the given user.
func GetListenHistory(user string) (*http.Response, error) {
	url := "https://api.listenbrainz.org/1/user/" + user + "/listens"

	return http.Get(url)
}
