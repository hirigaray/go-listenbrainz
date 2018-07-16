package listenbrainz

import "net/http"

func GetListenHistory(user string, count int) (*http.Response, error) {
	url := "https://api.listenbrainz.org/1/user/" + user + "/listens"

	return http.Get(url)
}
