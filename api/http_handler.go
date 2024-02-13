package api

import (
	"fmt"
	"io"
	"net/http"
	"todo_planning/util"
)

func GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		util.LogToFile(err.Error())
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		util.LogToFile("Get Data Error: " + fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		util.LogToFile(err.Error())
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}
