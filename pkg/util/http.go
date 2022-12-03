package util

import (
	"io/ioutil"
	"net/http"
)

// ReadAndClose http response
func ReadAndClose(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	respBody := resp.Body
	defer respBody.Close()
	return ioutil.ReadAll(respBody)
}
