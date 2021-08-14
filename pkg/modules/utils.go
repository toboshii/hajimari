package modules

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func updatedAt() string {
	return fmt.Sprintf("(Last update at %s)", time.Now().Format("15:04"))
}

func actualFetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
