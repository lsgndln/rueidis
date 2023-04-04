package om

import (
	"testing"

	"github.com/lsgndln/rueidis"
)

func setup(t *testing.T) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6377"}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}
